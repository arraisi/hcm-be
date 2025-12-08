-- Test the new V2 approach (direct MERGE without stored procedure)
-- This doesn't require EXECUTE permissions

SET NOCOUNT ON;

-- Test 1: First sequence for H/R/10101/25
DECLARE @OutputTable1 TABLE (next_running BIGINT);

MERGE tr_hasjratid_sequence WITH (HOLDLOCK) AS t
USING (
    SELECT 
        'H' AS c_source_code,
        'R' AS c_customer_type_code,
        '10101' AS c_outlet,
        '25' AS c_year
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
OUTPUT inserted.i_last_running INTO @OutputTable1;

SELECT next_running AS Test1_Result FROM @OutputTable1;

-- Test 2: Second sequence for same combination (should increment)
DECLARE @OutputTable2 TABLE (next_running BIGINT);

MERGE tr_hasjratid_sequence WITH (HOLDLOCK) AS t
USING (
    SELECT 
        'H' AS c_source_code,
        'R' AS c_customer_type_code,
        '10101' AS c_outlet,
        '25' AS c_year
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
OUTPUT inserted.i_last_running INTO @OutputTable2;

SELECT next_running AS Test2_Result FROM @OutputTable2;

-- View all sequences
SELECT * FROM tr_hasjratid_sequence ORDER BY d_updated_at DESC;
