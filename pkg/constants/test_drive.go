package constants

const (
	TestDriveBookingStatusSubmitted       = "SUBMITTED"
	TestDriveBookingStatusChangeRequest   = "CHANGE_REQUEST"
	TestDriveBookingStatusCancelSubmitted = "CANCEL_SUBMITTED"
	TestDriveBookingStatusConfirmed       = "CONFIRMED"
)

var (
	TestDriveStatusMap = map[string]string{
		TestDriveBookingStatusSubmitted:       "Submitted",
		TestDriveBookingStatusChangeRequest:   "Change Request",
		TestDriveBookingStatusCancelSubmitted: "Cancel Submitted",
	}
)
