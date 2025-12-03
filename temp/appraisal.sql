CREATE TABLE tr_appraisal
(
    -- Internal PK
    i_id UNIQUEIDENTIFIER
        CONSTRAINT DF_tr_appraisal_i_id DEFAULT NEWID()
        CONSTRAINT PK_tr_appraisal PRIMARY KEY,

    -- From appraisal request / confirm / update
    appraisal_booking_id      UNIQUEIDENTIFIER        NOT NULL, -- "appraisal_booking_ID"
    appraisal_booking_number  VARCHAR(50)             NOT NULL, -- "appraisal_booking_number"
    appraisal_type            VARCHAR(50)             NULL,     -- "APPRAISAL_BOOKING"

    outlet_id                 VARCHAR(50)             NOT NULL, -- "outlet_ID"
    outlet_name               VARCHAR(150)            NOT NULL, -- "outlet_name"

    appraisal_location        VARCHAR(50)             NULL,     -- "HOME_OR_OTHER_ADDRESS" / "DEALER"
    home_address              NVARCHAR(255)           NULL,
    province                  VARCHAR(100)            NULL,
    city                      VARCHAR(100)            NULL,
    district                  VARCHAR(100)            NULL,
    subdistrict               VARCHAR(100)            NULL,
    postal_code               VARCHAR(10)             NULL,

    -- Timestamps (converted from UNIX epoch)
    created_datetime                      DATETIME2(0) NOT NULL,
    appraisal_start_datetime              DATETIME2(0) NULL,
    appraisal_end_datetime                DATETIME2(0) NULL,
    appraisal_confirmation_start_datetime DATETIME2(0) NULL,
    appraisal_confirmation_end_datetime   DATETIME2(0) NULL,

    -- Booking status & cancellation (request / confirm)
    appraisal_booking_status   VARCHAR(30)   NOT NULL, -- SUBMITTED / CONFIRMED / CANCELLED / NO_SHOW
    cancelled_by               VARCHAR(30)   NULL,     -- DEALER / CUSTOMER / SYSTEM
    cancellation_reason        VARCHAR(50)   NULL,     -- NO_SHOW / OTHERS / etc.
    other_cancellation_reason  NVARCHAR(255) NULL,

    -- Vehicle info from update
    katashiki_suffix           VARCHAR(50)   NULL,
    color_code                 VARCHAR(20)   NULL,
    model                      VARCHAR(100)  NULL,
    variant                    VARCHAR(100)  NULL,
    color                      VARCHAR(50)   NULL,

    -- Trade-in summary (latest status)
    final_trade_in_status          VARCHAR(30)   NULL, -- NEGOTIATION / HANDOVER / NO_DEAL / CANCELLED
    last_trade_in_status_datetime  DATETIME2(0)  NULL,

    -- Sales data
    spk_number                 VARCHAR(50)   NULL, -- "spk_number"
    so_number                  VARCHAR(50)   NULL, -- "so_number"

    -- Negotiation detail (from appraisal update > negotiation)
    customer_negotiation_price                DECIMAL(18,2) NULL,
    dealer_negotiation_price                  DECIMAL(18,2) NULL,
    deal_price                                DECIMAL(18,2) NULL,
    down_payment_estimation                   DECIMAL(18,2) NULL,
    estimated_remaining_payment               DECIMAL(18,2) NULL,
    no_deal_reason                            VARCHAR(50)    NULL,
    no_deal_reason_old_vehicle_others         NVARCHAR(255)  NULL,
    no_deal_reason_old_vehicle_expected_sell_price DECIMAL(18,2) NULL,
    no_deal_reason_old_vehicle_price_sold     DECIMAL(18,2) NULL,
    no_deal_reason_new_vehicle_others         NVARCHAR(255)  NULL,

    -- Handover detail (from appraisal update > handover)
    trade_in_payment_datetime     DATETIME2(0) NULL,
    trade_in_handover_status      VARCHAR(30)  NULL, -- COMPLETED / etc.
    trade_in_handover_datetime    DATETIME2(0) NULL,
    trade_in_handover_location    VARCHAR(30)  NULL, -- DEALER / HOME
    trade_in_handover_address     NVARCHAR(255) NULL,
    handover_province             VARCHAR(100) NULL,
    handover_city                 VARCHAR(100) NULL,
    handover_district             VARCHAR(100) NULL,
    handover_subdistrict          VARCHAR(100) NULL,
    handover_postal_code          VARCHAR(10)  NULL,

    -- Audit columns
    d_createdate  DATETIME2(0) NOT NULL CONSTRAINT DF_tr_appraisal_d_createdate DEFAULT SYSDATETIME(),
    d_updatedate  DATETIME2(0) NULL
);

-- Optional: unique index so 1 booking = 1 appraisal row
CREATE UNIQUE INDEX UX_tr_appraisal_booking
    ON tr_appraisal (appraisal_booking_id, appraisal_booking_number);

CREATE TABLE tr_appraisal_status_update
(
    i_id UNIQUEIDENTIFIER
        CONSTRAINT DF_tr_appraisal_status_update_i_id DEFAULT NEWID()
        CONSTRAINT PK_tr_appraisal_status_update PRIMARY KEY,

    appraisal_id UNIQUEIDENTIFIER NOT NULL
        CONSTRAINT FK_tr_appraisal_status_update_tr_appraisal
            REFERENCES tr_appraisal (i_id),

    trade_in_status           VARCHAR(30)  NOT NULL, -- NEGOTIATION / HANDOVER / NO_DEAL / CANCELLED
    trade_in_status_datetime  DATETIME2(0) NOT NULL,

    d_createdate DATETIME2(0) NOT NULL
        CONSTRAINT DF_tr_appraisal_status_update_d_createdate DEFAULT SYSDATETIME()
);

CREATE INDEX IX_tr_appraisal_status_update_appraisal
    ON tr_appraisal_status_update (appraisal_id, trade_in_status_datetime);
