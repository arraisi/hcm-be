-- Debug: Test the stored procedure directly
DECLARE @result BIGINT;

EXEC sp_get_next_hasjrat_running
    @p_source_code = 'H',
    @p_customer_type_code = 'R',
    @p_outlet_code = '10101',
    @p_year = '25',
    @p_next_running = @result OUTPUT;

SELECT @result AS result_from_output_param;

-- Also check what's in the table
SELECT * FROM tr_hasjratid_sequence;
