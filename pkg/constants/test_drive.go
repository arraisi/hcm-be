package constants

const (
	TestDriveBookingStatusSubmitted       = "SUBMITTED"
	TestDriveBookingStatusChangeRequest   = "CHANGE_REQUEST"
	TestDriveBookingStatusCancelSubmitted = "CANCEL_SUBMITTED"
	TestDriveBookingStatusConfirmed       = "CONFIRMED"
	TestDriveBookingStatusCancelled       = "CANCELLED"
	TestDriveBookingStatusCompleted       = "COMPLETED"
	TestDriveBookingStatusNoShow          = "NO_SHOW"

	// Test Drive Locations
	TestDriveLocationHome   = "HOME"
	TestDriveLocationDealer = "DEALER"

	// Cancellation Reasons
	CancellationReasonScheduleConflict            = "SCHEDULE_CONFLICT"
	CancellationReasonPricingConsideration        = "PRICING_CONSIDERATION"
	CancellationReasonFoundBetterDealElsewhere    = "FOUND_A_BETTER_DEAL_ELSEWHERE"
	CancellationReasonPoorCustomerService         = "POOR_CUSTOMER_SERVICE"
	CancellationReasonChangeInVehicleNeeds        = "CHANGE_IN_VEHICLE_NEEDS"
	CancellationReasonDoNotWantToDiscloseMyReason = "I_DO_NOT_WANT_TO_DISCLOSE_MY_REASONS"
	CancellationReasonOthers                      = "OTHERS"
)

var (
	TestDriveStatusMap = map[string]string{
		TestDriveBookingStatusSubmitted:       "Submitted",
		TestDriveBookingStatusChangeRequest:   "Change Request",
		TestDriveBookingStatusCancelSubmitted: "Cancel Submitted",
		TestDriveBookingStatusConfirmed:       "Confirmed",
		TestDriveBookingStatusCancelled:       "Cancelled",
		TestDriveBookingStatusCompleted:       "Completed",
		TestDriveBookingStatusNoShow:          "No Show",
	}

	TestDriveOnGoingStatus = []string{
		TestDriveBookingStatusSubmitted,
		TestDriveBookingStatusChangeRequest,
		TestDriveBookingStatusConfirmed,
	}

	TestDriveLocationMap = map[string]string{
		TestDriveLocationHome:   "Home",
		TestDriveLocationDealer: "Dealer",
	}

	CancellationReasonMap = map[string]string{
		CancellationReasonScheduleConflict:            "Schedule Conflict",
		CancellationReasonPricingConsideration:        "Pricing Consideration",
		CancellationReasonFoundBetterDealElsewhere:    "Found a Better Deal Elsewhere",
		CancellationReasonPoorCustomerService:         "Poor Customer Service",
		CancellationReasonChangeInVehicleNeeds:        "Change in Vehicle Needs",
		CancellationReasonDoNotWantToDiscloseMyReason: "I Do Not Want to Disclose My Reasons",
		CancellationReasonOthers:                      "Others",
	}
)
