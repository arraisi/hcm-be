package hasjratid

import (
	"context"
	"fmt"
)

// GetNextSequence calls sp_get_next_hasjrat_running and returns the next running number.
//
// Params should already be normalized:
// - sourceCode:        "H" / "C"
// - customerTypeCode:  "R" / "G" / "C"
// - outletCode:        5-char padded outlet (e.g. "10101" -> "10101")
// - year:              2-char year string (e.g. "25")
func (r *repository) GetNextSequence(
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

	const rawQuery = `
DECLARE @next_running BIGINT;

EXEC sp_get_next_hasjrat_running
    @p_source_code        = ?,
    @p_customer_type_code = ?,
    @p_outlet_code        = ?,
    @p_year               = ?,
    @p_next_running       = @next_running OUTPUT;

SELECT @next_running AS next_running;
`

	// Allow driver-specific rebinding if needed (keeps style consistent)
	sqlQuery := r.db.Rebind(rawQuery)

	var nextRunning uint64
	if err := r.db.GetContext(
		ctx,
		&nextRunning,
		sqlQuery,
		sourceCode,
		customerTypeCode,
		outletCode,
		year,
	); err != nil {
		return 0, err
	}

	return nextRunning, nil
}
