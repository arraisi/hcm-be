package dmsaftersales

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/arraisi/hcm-be/internal/config"
	"github.com/arraisi/hcm-be/internal/domain/dto/servicebooking"
	"github.com/jmoiron/sqlx"
)

type iDB interface {
	BeginTxx(ctx context.Context, opts *sql.TxOptions) (*sqlx.Tx, error)
	PreparexContext(ctx context.Context, query string) (*sqlx.Stmt, error)
	Rebind(query string) string
	GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
}

type client struct {
	cfg *config.Config
	db  iDB
}

// New creates a new customer repository instance
func New(cfg *config.Config, db iDB) *client {
	return &client{
		db:  db,
		cfg: cfg,
	}
}

func (c *client) InsertServiceBookingRequest(ctx context.Context, event servicebooking.ServiceBookingEvent) error {
	// Call Oracle stored procedure to insert service booking request into DMS After Sales database
	//
	// IMPORTANT NOTES:
	// 1. This implementation calls the procedure once per job. If you have multiple jobs,
	//    it will make multiple procedure calls. Consider if this is the desired behavior.
	// 2. Similarly for warranties, recalls, and parts - only the first item of each is sent.
	//    You may need to modify this to handle multiple items properly.
	// 3. The stored procedure name 'insert_booking_data' should match your actual Oracle procedure.
	// 4. Package ID mapping and service category ID mapping may need adjustment based on your DB schema.
	// 5. Date formatting for recall_date and timestamps should be verified with your Oracle setup.
	// 6. Consider using Oracle VARRAY or TABLE types if you need to pass multiple items in one call.

	// Prepare arrays for storing procedure parameters
	// Note: Oracle stored procedures handle arrays/collections differently depending on the implementation
	// This example assumes the procedure handles individual parameters or we call it multiple times

	// Extract data for easier access
	oneAccount := event.Data.OneAccount
	vehicle := event.Data.CustomerVehicle
	booking := event.Data.ServiceBookingRequest

	// For each warranty, recall, job, and part, we need to call the procedure
	// Since Oracle doesn't easily support passing arrays in sqlx, we'll need to call it per item
	// or use Oracle's VARRAY/TABLE types

	// Build the PL/SQL block for calling the stored procedure
	// This assumes insert_booking_data can handle individual records

	for i := range event.Data.Job {
		job := event.Data.Job[i]

		// Get first warranty if exists
		var warrantyType, warrantyStatus string
		if len(booking.Warranty) > 0 {
			warrantyType = booking.Warranty[0].WarrantyName
			warrantyStatus = booking.Warranty[0].WarrantyStatus
		}

		// Get first recall if exists
		var recallID, recallDate, recallDesc string
		if len(booking.Recalls) > 0 {
			recallID = booking.Recalls[0].RecallID
			recallDate = booking.Recalls[0].RecallDate
			recallDesc = booking.Recalls[0].RecallDescription
		}

		// Get first recall part if exists
		var recallPartID string
		if len(booking.Recalls) > 0 && len(booking.Recalls[0].AffectedParts) > 0 {
			recallPartID = booking.Recalls[0].AffectedParts[0]
		}

		// Get first part if exists
		var partCode, partName string
		var partQty int32
		var partPrice, partDiscount float32
		if len(event.Data.Part) > 0 {
			part := event.Data.Part[0]
			partCode = part.PartNumber
			partName = part.PartName
			partQty = part.PartQuantity
			partPrice = part.PartEstPrice
			// Discount calculation might need to be added to PartRequest
			partDiscount = 0
		}

		// Get package parts if exists
		var pkgID int32
		var pkgQty int32 = 1
		var pkgPartCode, pkgPartName string
		if len(event.Data.Part) > 0 && len(event.Data.Part[0].PackageParts) > 0 {
			// Package ID might need to be converted or mapped
			pkgPartCode = event.Data.Part[0].PackageParts[0].PartNumber
			pkgPartName = event.Data.Part[0].PackageParts[0].PartName
		}

		// Build PL/SQL block
		query := `
			BEGIN
				insert_booking_data(
					-- TBL_APP_ONE_ACCOUNT
					p_acc_id              => :1,
					p_first_name          => :2,
					p_last_name           => :3,
					p_gender              => :4,
					p_phone               => :5,
					p_email               => :6,
					p_account_type        => :7,

					-- TBL_APP_CUSTOMER_VEHICLE
					p_vin                 => :8,
					p_model_code          => :9,
					p_model_name          => :10,
					p_variant             => :11,
					p_license_plate       => :12,
					p_mileage             => :13,
					p_color_code          => :14,
					p_color_name          => :15,

					-- TBL_APP_SERVICE_BOOKING
					p_booking_id          => :16,
					p_channel             => :17,
					p_status              => :18,
					p_created_at          => :19,
					p_service_type        => :20,
					p_service_category_id => :21,
					p_start_time_epoch    => :22,
					p_end_time_epoch      => :23,
					p_booking_start_epoch => :24,
					p_booking_end_epoch   => :25,
					p_contact_name        => :26,
					p_contact_phone       => :27,
					p_customer_phone      => :28,
					p_start_time_text     => :29,
					p_end_time_text       => :30,
					p_booking_type        => :31,
					p_workshop_id         => :32,
					p_workshop_name       => :33,
					p_workshop_address    => :34,
					p_province            => :35,
					p_city                => :36,
					p_district            => :37,
					p_subdistrict         => :38,
					p_postal_code         => :39,
					p_notes               => :40,

					-- TBL_APP_SB_WARRANTY
					p_warranty_type       => :41,
					p_warranty_status     => :42,

					-- TBL_APP_SB_RECALL
					p_recall_id           => :43,
					p_recall_date         => TO_DATE(:44, 'DD-MM-YYYY'),
					p_recall_desc         => :45,

					-- TBL_APP_SB_RECALL_PART
					p_recall_part_id      => :46,

					-- TBL_APP_SERVICE_BOOKING_JOB
					p_job_id              => :47,
					p_job_desc            => :48,
					p_job_price           => :49,
					p_job_category        => :50,

					-- TBL_APP_SERVICE_BOOKING_PART
					p_part_code           => :51,
					p_part_name           => :52,
					p_part_qty            => :53,
					p_part_price          => :54,
					p_part_discount       => :55,

					-- TBL_APP_SB_PART_PACKAGE_PARTS
					p_pkg_id              => :56,
					p_pkg_qty             => :57,
					p_pkg_part_code       => :58,
					p_pkg_part_name       => :59
				);
			END;`

		// Execute the stored procedure
		_, err := c.db.ExecContext(ctx, query,
			// TBL_APP_ONE_ACCOUNT (1-7)
			oneAccount.OneAccountID,
			oneAccount.FirstName,
			oneAccount.LastName,
			oneAccount.Gender,
			oneAccount.PhoneNumber,
			oneAccount.Email,
			"MASTER", // Default primary user type

			// TBL_APP_CUSTOMER_VEHICLE (8-15)
			vehicle.Vin,
			vehicle.KatashikiSuffix,
			vehicle.Model,
			vehicle.Variant,
			vehicle.PoliceNumber,
			vehicle.ActualMileage,
			vehicle.ColorCode,
			vehicle.Color,

			// TBL_APP_SERVICE_BOOKING (16-40)
			booking.BookingId,
			booking.BookingSource,
			booking.BookingStatus,
			booking.CreatedDatetime,
			event.Process, // service_type from process field
			getServiceCategoryID(booking.ServiceCategory),
			booking.SlotDatetimeStart,
			booking.SlotDatetimeEnd,
			booking.SlotRequestedDatetimeStart,
			booking.SlotRequestedDatetimeEnd,
			fmt.Sprintf("%s %s", oneAccount.FirstName, oneAccount.LastName),
			booking.PreferenceContactPhoneNumber,
			oneAccount.PhoneNumber,
			formatEpochToDateTime(booking.SlotDatetimeStart),
			formatEpochToDateTime(booking.SlotDatetimeEnd),
			event.Process,
			booking.OutletID,
			booking.OutletName,
			booking.MobileServiceAddress,
			booking.Province,
			booking.City,
			booking.District,
			booking.SubDistrict,
			booking.PostalCode,
			getBookingNotes(booking),

			// TBL_APP_SB_WARRANTY (41-42)
			warrantyType,
			warrantyStatus,

			// TBL_APP_SB_RECALL (43-45)
			recallID,
			recallDate,
			recallDesc,

			// TBL_APP_SB_RECALL_PART (46)
			recallPartID,

			// TBL_APP_SERVICE_BOOKING_JOB (47-50)
			job.JobID,
			job.JobName,
			job.LaborEstPrice,
			"MAINTENANCE", // Default category or derive from job type

			// TBL_APP_SERVICE_BOOKING_PART (51-55)
			partCode,
			partName,
			partQty,
			partPrice,
			partDiscount,

			// TBL_APP_SB_PART_PACKAGE_PARTS (56-59)
			pkgID,
			pkgQty,
			pkgPartCode,
			pkgPartName,
		)

		if err != nil {
			return fmt.Errorf("failed to insert service booking data for job %s: %w", job.JobID, err)
		}
	}

	return nil
}

// Helper functions
func getServiceCategoryID(category string) int {
	// Map service category string to ID
	categoryMap := map[string]int{
		"PERIODIC_MAINTENANCE": 1,
		"BODY_AND_PAINT":       2,
		// Add more mappings as needed
	}
	if id, ok := categoryMap[category]; ok {
		return id
	}
	return 0
}

func formatEpochToDateTime(epoch int64) string {
	if epoch == 0 {
		return ""
	}
	// Convert epoch to "YYYY-MM-DD HH:MM" format
	// This is a placeholder - adjust format as needed
	return fmt.Sprintf("%d", epoch) // TODO: proper time formatting
}

func getBookingNotes(booking servicebooking.ServiceBookingRequest) string {
	if booking.VehicleProblem != "" {
		return booking.VehicleProblem
	}
	if booking.AdditionalVehicleProblem != "" {
		return booking.AdditionalVehicleProblem
	}
	return ""
}
