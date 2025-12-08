-- ============================================================================
-- FIX for existing stored procedure sp_get_next_hasjrat_running
-- This ALTER statement fixes the NULL return issue
-- ============================================================================

ALTER PROCEDURE sp_get_next_hasjrat_running 
    @p_source_code        char(1),
    @p_customer_type_code char(1),
    @p_outlet_code        char(5),
    @p_year               char(2),
    @p_next_running       bigint OUTPUT
AS
BEGIN
    SET NOCOUNT ON;
    
    -- Declare a table variable to capture the OUTPUT from MERGE
    DECLARE @OutputTable TABLE (next_running bigint);
    
    -- Use MERGE to atomically insert or update, capturing the result
    MERGE tr_hasjratid_sequence WITH (HOLDLOCK) AS t
    USING (
        VALUES (@p_source_code, @p_customer_type_code, @p_outlet_code, @p_year)
    ) AS v(c_source_code, c_customer_type_code, c_outlet, c_year)
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
    
    -- Assign the captured value to the OUTPUT parameter
    SELECT @p_next_running = next_running FROM @OutputTable;
END
GO

PRINT 'Stored procedure sp_get_next_hasjrat_running created successfully';
GO
