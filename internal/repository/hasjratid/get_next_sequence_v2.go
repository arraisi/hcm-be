package hasjratid

import (
	"context"
	"database/sql"
	"fmt"
)

// GetNextSequenceV2 gets the next sequence number directly using SQL MERGE without stored procedure.
// This approach doesn't require EXECUTE permissions on stored procedures.
//
// Params should already be normalized:
// - sourceCode:        "H" / "C"
// - customerTypeCode:  "R" / "G" / "C"
// - outletCode:        5-char padded outlet (e.g. "10101" -> "10101")
// - year:              2-char year string (e.g. "25")
func (r *repository) GetNextSequenceV2(
	ctx context.Context,
	sourceCode string,
	customerTypeCode string,
	outletCode string,
	year string,
) (uint64, error) {
	if len(sourceCode) != 1 {
		return 0, fmt.Errorf("sourceCode must be 1 char, got %q", sourceCode)
	}
	if len(customerTypeCode) != 1 {
		return 0, fmt.Errorf("customerTypeCode must be 1 char, got %q", customerTypeCode)
	}
	if len(outletCode) != 5 {
		return 0, fmt.Errorf("outletCode must be 5 chars, got %q", outletCode)
	}
	if len(year) != 2 {
		return 0, fmt.Errorf("year must be 2 chars, got %q", year)
	}

	// Use MERGE directly to atomically get and increment the sequence
	const rawQuery = `
SET NOCOUNT ON;

DECLARE @OutputTable TABLE (next_running BIGINT);

MERGE tr_hasjratid_sequence WITH (HOLDLOCK) AS t
USING (
    SELECT 
        @p_source_code AS c_source_code,
        @p_customer_type_code AS c_customer_type_code,
        @p_outlet_code AS c_outlet,
        @p_year AS c_year
) AS v
ON t.c_source_code = v.c_source_code
    AND t.c_customer_type_code = v.c_customer_type_code
    AND t.c_outlet = v.c_outlet
    AND t.c_year = v.c_year
WHEN MATCHED THEN
    UPDATE SET
        t.i_last_running = t.i_last_running + 1,
        t.d_updated_at = GETDATE()
WHEN NOT MATCHED THEN
    INSERT (c_source_code, c_customer_type_code, c_outlet, c_year, i_last_running, d_updated_at)
    VALUES (v.c_source_code, v.c_customer_type_code, v.c_outlet, v.c_year, 1, GETDATE())
OUTPUT inserted.i_last_running INTO @OutputTable;

SELECT next_running FROM @OutputTable;
`

	var nextRunning sql.NullInt64
	if err := r.db.GetContext(
		ctx,
		&nextRunning,
		rawQuery,
		sql.Named("p_source_code", sourceCode),
		sql.Named("p_customer_type_code", customerTypeCode),
		sql.Named("p_outlet_code", outletCode),
		sql.Named("p_year", year),
	); err != nil {
		return 0, fmt.Errorf("failed to get next sequence: %w", err)
	}

	// Check if the query returned NULL (should not happen with MERGE)
	if !nextRunning.Valid {
		return 0, fmt.Errorf(
			"MERGE returned NULL for sourceCode=%q, customerType=%q, outlet=%q, year=%q",
			sourceCode, customerTypeCode, outletCode, year,
		)
	}

	return uint64(nextRunning.Int64), nil
}
