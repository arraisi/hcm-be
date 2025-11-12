-- auto-generated definition for Customer Reminder
CREATE TABLE dbo.tr_customer_reminder
(
    i_id                        VARCHAR(36)                NOT NULL,  -- PK UUID
    i_customer_id               VARCHAR(36)                NOT NULL,  -- FK to tr_customer
    i_customer_vehicle_id       VARCHAR(36)                NULL,      -- FK to tm_customer_vehicle
    i_outlet_id                 VARCHAR(36)                NOT NULL,  -- outlet ID
    i_reminder_id               VARCHAR(36)                NOT NULL,  -- external reminder ID
    c_activity                  VARCHAR(64)                NOT NULL,  -- e.g., SERVICE_BOOKING, NCS, etc.
    d_activity_plan_scheduled_date DATETIME                NULL,
    c_auto_reminder_status      VARCHAR(32)                NULL,      -- e.g., DELIVERED, READ
    e_reminder_message          VARCHAR(512)               NULL,
    v_priority_call             INT                        NULL,
    c_extended_warranty_status  VARCHAR(32)                NULL,      -- e.g., ELIGIBLE, NOT_ELIGIBLE
    c_customer_habit            VARCHAR(32)                NULL,      -- TIME_BASED, MILEAGE
    c_last_habit                VARCHAR(32)                NULL,      -- PUNCTUAL, LATE, etc.
    c_next_service_status       VARCHAR(32)                NULL,
    d_last_service_date         DATETIME                   NULL,
    d_next_service_date         DATETIME                   NULL,
    c_ncs_status                VARCHAR(32)                NULL,      -- SAME_OUTLET, DIFFERENT
    c_program_tab               VARCHAR(32)                NULL,      -- T_CARE, GBSB, REGULAR
    v_next_service_stage        INT                        NULL,

    -- Audit fields
    d_created_at                DATETIME DEFAULT GETDATE() NOT NULL,
    c_created_by                VARCHAR(32)                NULL,
    d_updated_at                DATETIME DEFAULT GETDATE() NULL,
    c_updated_by                VARCHAR(32)                NULL,

    CONSTRAINT PK_tr_customer_reminder PRIMARY KEY (i_id),
    CONSTRAINT FK_tr_customer_reminder_customer FOREIGN KEY (i_customer_id) REFERENCES dbo.tr_customer(i_id),
    CONSTRAINT FK_tr_customer_reminder_vehicle FOREIGN KEY (i_customer_vehicle_id) REFERENCES dbo.tm_customer_vehicle(i_id)
);
GO
