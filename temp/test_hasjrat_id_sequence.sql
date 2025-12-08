-- ============================================================================
-- Test script for sp_get_next_hasjrat_running
-- This will verify the sequence increments correctly: 1, 2, 3, ..., n
-- ============================================================================

-- Test Case 1: Generate first ID for H/R/10101/25
DECLARE @next1 BIGINT;
EXEC sp_get_next_hasjrat_running 
    @p_source_code = 'H',
    @p_customer_type_code = 'R',
    @p_outlet_code = '10101',
    @p_year = '25',
    @p_next_running = @next1 OUTPUT;
SELECT @next1 AS Test1_Should_Be_1;

-- Test Case 2: Generate second ID for same combination
DECLARE @next2 BIGINT;
EXEC sp_get_next_hasjrat_running 
    @p_source_code = 'H',
    @p_customer_type_code = 'R',
    @p_outlet_code = '10101',
    @p_year = '25',
    @p_next_running = @next2 OUTPUT;
SELECT @next2 AS Test2_Should_Be_2;

-- Test Case 3: Generate third ID for same combination
DECLARE @next3 BIGINT;
EXEC sp_get_next_hasjrat_running 
    @p_source_code = 'H',
    @p_customer_type_code = 'R',
    @p_outlet_code = '10101',
    @p_year = '25',
    @p_next_running = @next3 OUTPUT;
SELECT @next3 AS Test3_Should_Be_3;

-- Test Case 4: Different combination (different customer type)
DECLARE @next4 BIGINT;
EXEC sp_get_next_hasjrat_running 
    @p_source_code = 'H',
    @p_customer_type_code = 'C',
    @p_outlet_code = '10101',
    @p_year = '25',
    @p_next_running = @next4 OUTPUT;
SELECT @next4 AS Test4_Should_Be_1;

-- Test Case 5: Different outlet
DECLARE @next5 BIGINT;
EXEC sp_get_next_hasjrat_running 
    @p_source_code = 'H',
    @p_customer_type_code = 'R',
    @p_outlet_code = '10102',
    @p_year = '25',
    @p_next_running = @next5 OUTPUT;
SELECT @next5 AS Test5_Should_Be_1;

-- View all sequence records
SELECT 
    c_source_code,
    c_customer_type_code,
    c_outlet,
    c_year,
    i_last_running,
    d_updated_at
FROM tr_hasjratid_sequence
ORDER BY c_source_code, c_customer_type_code, c_outlet, c_year;
