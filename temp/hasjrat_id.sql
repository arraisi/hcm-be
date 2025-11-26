CREATE TABLE tr_hasjratid_sequence
(
    c_source_code        char(1)  NOT NULL, -- H / C
    c_customer_type_code char(1)  NOT NULL, -- R / G / C
    c_outlet             char(5)  NOT NULL, -- 5 digit outlet
    c_year               char(2)  NOT NULL, -- 2 digit year (25)
    i_last_running       bigint   NOT NULL, -- last used number
    d_updated_at         datetime NOT NULL CONSTRAINT DF_tr_hasjratid_sequence_d_updated_at DEFAULT (GETDATE()),

    CONSTRAINT PK_tr_hasjratid_sequence
        PRIMARY KEY (c_source_code, c_customer_type_code, c_outlet, c_year)
);
GO

CREATE PROCEDURE sp_get_next_hasjrat_running @p_source_code        char(1),
    @p_customer_type_code char(1),
    @p_outlet_code        char(5),
    @p_year               char(2),
    @p_next_running       bigint OUTPUT
AS
BEGIN
    SET
NOCOUNT ON;

MERGE tr_hasjratid_sequence AS t
    USING (VALUES (@p_source_code, @p_customer_type_code, @p_outlet_code, @p_year)) AS v(c_source_code, c_customer_type_code, c_outlet, c_year)
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
    OUTPUT inserted.i_last_running;

SELECT @p_next_running = i_last_running
FROM tr_hasjratid_sequence
WHERE c_source_code = @p_source_code
  AND c_customer_type_code = @p_customer_type_code
  AND c_outlet = @p_outlet_code
  AND c_year = @p_year;
END
GO
