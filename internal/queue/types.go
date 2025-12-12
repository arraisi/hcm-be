package queue

const (
	// TaskTypeDIDXServiceBookingConfirm represents the task type for DIDX confirm operations
	TaskTypeDIDXServiceBookingConfirm  = "didx:service_booking_confirm"
	TaskTypeDIDXTestDriveConfirm       = "didx:test_drive_confirm"
	TaskTypeDIDXAppraisalConfirm       = "didx:appraisal_confirm"
	TaskTypeDMSTestDriveRequest        = "dms:test_drive_request"
	TaskTypeDMSCreateOneAccess         = "dms:create_one_access"
	TaskTypeDMSCreateToyotaID          = "dms:create_toyota_id"
	TaskTypeDMSCreateGetOffer          = "dms:create_get_offer"
	TaskTypeDMSAppraisalBookingRequest = "dms:appraisal_booking_request"
	TaskTypeDMSServiceBookingRequest   = "dms:service_booking_request"
)
