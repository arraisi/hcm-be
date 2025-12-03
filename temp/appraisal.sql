CREATE TABLE tr_appraisal
(
    -- Internal PK
    i_id UNIQUEIDENTIFIER
        CONSTRAINT DF_tr_appraisal_i_id DEFAULT NEWID()
        CONSTRAINT PK_tr_appraisal PRIMARY KEY,

    -- Mapping
    i_leads_id        VARCHAR(50) NULL;
    i_one_account_id  VARCHAR(50) NULL;
    c_vin             VARCHAR(50) NULL;

    -- From appraisal request / confirm / update
    i_appraisal_booking_id     UNIQUEIDENTIFIER        NOT NULL, -- "appraisal_booking_ID"
    c_appraisal_booking_number VARCHAR(50)             NOT NULL, -- "appraisal_booking_number"
    c_appraisal_type           VARCHAR(50)             NULL,     -- "APPRAISAL_BOOKING"

    i_outlet_id                VARCHAR(50)             NOT NULL, -- "outlet_ID"
    n_outlet_name              VARCHAR(150)            NOT NULL, -- "outlet_name"

    c_appraisal_location       VARCHAR(50)             NULL,     -- "HOME_OR_OTHER_ADDRESS" / "DEALER"
    e_home_address             NVARCHAR(255)           NULL,
    e_province                 VARCHAR(100)            NULL,
    e_city                     VARCHAR(100)            NULL,
    e_district                 VARCHAR(100)            NULL,
    e_subdistrict              VARCHAR(100)            NULL,
    e_postal_code              VARCHAR(10)             NULL,

    -- Timestamps (converted from UNIX epoch)
    d_created_datetime                      DATETIME2(0) NOT NULL,
    d_appraisal_start_datetime              DATETIME2(0) NULL,
    d_appraisal_end_datetime                DATETIME2(0) NULL,
    d_appraisal_confirmation_start_datetime DATETIME2(0) NULL,
    d_appraisal_confirmation_end_datetime   DATETIME2(0) NULL,

    -- Booking status & cancellation (request / confirm)
    c_appraisal_booking_status  VARCHAR(30)   NOT NULL, -- SUBMITTED / CONFIRMED / CANCELLED / NO_SHOW
    c_cancelled_by              VARCHAR(30)   NULL,     -- DEALER / CUSTOMER / SYSTEM
    c_cancellation_reason       VARCHAR(50)   NULL,     -- NO_SHOW / OTHERS / etc.
    e_other_cancellation_reason NVARCHAR(255) NULL,

    -- Vehicle info from update
    c_katashiki_suffix          VARCHAR(50)   NULL,
    c_color_code                VARCHAR(20)   NULL,
    n_model                     VARCHAR(100)  NULL,
    n_variant                   VARCHAR(100)  NULL,
    n_color                     VARCHAR(50)   NULL,

    -- Trade-in summary (latest status)
    c_final_trade_in_status         VARCHAR(30)   NULL, -- NEGOTIATION / HANDOVER / NO_DEAL / CANCELLED
    d_last_trade_in_status_datetime DATETIME2(0)  NULL,

    -- Sales data
    c_spk_number                VARCHAR(50)   NULL, -- "spk_number"
    c_so_number                 VARCHAR(50)   NULL, -- "so_number"

    -- Negotiation detail (from appraisal update > negotiation)
    v_customer_negotiation_price                DECIMAL(18,2) NULL,
    v_dealer_negotiation_price                  DECIMAL(18,2) NULL,
    v_deal_price                                DECIMAL(18,2) NULL,
    v_down_payment_estimation                   DECIMAL(18,2) NULL,
    v_estimated_remaining_payment               DECIMAL(18,2) NULL,
    c_no_deal_reason                            VARCHAR(50)    NULL,
    e_no_deal_reason_old_vehicle_others         NVARCHAR(255)  NULL,
    v_no_deal_reason_old_vehicle_expected_sell_price DECIMAL(18,2) NULL,
    v_no_deal_reason_old_vehicle_price_sold     DECIMAL(18,2) NULL,
    e_no_deal_reason_new_vehicle_others         NVARCHAR(255)  NULL,

    -- Handover detail (from appraisal update > handover)
    d_trade_in_payment_datetime     DATETIME2(0) NULL,
    c_trade_in_handover_status      VARCHAR(30)  NULL, -- COMPLETED / etc.
    d_trade_in_handover_datetime    DATETIME2(0) NULL,
    c_trade_in_handover_location    VARCHAR(30)  NULL, -- DEALER / HOME
    e_trade_in_handover_address     NVARCHAR(255) NULL,
    e_handover_province             VARCHAR(100) NULL,
    e_handover_city                 VARCHAR(100) NULL,
    e_handover_district             VARCHAR(100) NULL,
    e_handover_subdistrict          VARCHAR(100) NULL,
    e_handover_postal_code          VARCHAR(10)  NULL,

    -- Audit columns
    d_createdate  DATETIME2(0) NOT NULL CONSTRAINT DF_tr_appraisal_d_createdate DEFAULT SYSDATETIME(),
    d_updatedate  DATETIME2(0) NULL
);

-- 1 booking = 1 appraisal row
CREATE UNIQUE INDEX UX_tr_appraisal_booking
    ON tr_appraisal (i_appraisal_booking_id, c_appraisal_booking_number);
