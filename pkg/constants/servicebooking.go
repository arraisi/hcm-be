package constants

const (
	// Service Categories
	ServiceCategoryPeriodicMaintenance = "PERIODIC_MAINTENANCE"
	ServiceCategoryBodyAndPaint        = "BODY_AND_PAINT"

	// Service Booking Statuses
	ServiceBookingStatusSubmitted         = "SUBMITTED"
	ServiceBookingStatusManuallyConfirmed = "MANUALLY_CONFIRMED"
	ServiceBookingStatusSystemConfirmed   = "SYSTEM_CONFIRMED"
	ServiceBookingStatusChangeRequest     = "CHANGE_REQUEST"
)

var (
	ServiceCategoryMap = map[string]string{
		ServiceCategoryPeriodicMaintenance: "Periodic Maintenance",
		ServiceCategoryBodyAndPaint:        "Body and Paint",
	}

	ServiceBookingStatusMap = map[string]string{
		ServiceBookingStatusSubmitted:         "Submitted",
		ServiceBookingStatusManuallyConfirmed: "Manually Confirmed",
		ServiceBookingStatusSystemConfirmed:   "System Confirmed",
		ServiceBookingStatusChangeRequest:     "Change Request",
	}

	// Active service booking statuses that prevent new periodic maintenance bookings
	ActiveServiceBookingStatuses = []string{
		ServiceBookingStatusSubmitted,
		ServiceBookingStatusManuallyConfirmed,
		ServiceBookingStatusSystemConfirmed,
		ServiceBookingStatusChangeRequest,
	}
)
