package constants

const (
	LeadsFollowUpStatusOnConsideration  = "ON_CONSIDERATION"
	LeadsFollowUpStatusNotYetFollowedUp = "NOT_YET_FOLLOWED_UP"
	LeadsFollowUpStatusNoResponse       = "NO_RESPONSE"
	LeadsTypeTestDriveRequest           = "TEST_DRIVE_REQUEST"

	// Lead Scores
	LeadScoreLow    = "LOW"
	LeadScoreMedium = "MEDIUM"
	LeadScoreHot    = "HOT"

	// Leads Sources
	LeadsSourceCustomerDatabase        = "CUSTOMER_DATABASE"
	LeadsSourceDigitalToyotaWeb        = "DIGITAL_TOYOTA_WEB"
	LeadsSourceDigitalMToyota          = "DIGITAL_MTOYOTA"
	LeadsSourceDigitalOthers           = "DIGITAL_OTHERS"
	LeadsSourceDigitalDealer           = "DIGITAL_DEALER"
	LeadsSourceOfflineEvent            = "OFFLINE_EVENT"
	LeadsSourceOfflineGathering        = "OFFLINE_GATHERING"
	LeadsSourceOfflineCustomerReferral = "OFFLINE_CUSTOMER_REFERRAL"
	LeadsSourceOfflineWalkInOrCallIn   = "OFFLINE_WALK_IN_OR_CALL_IN"
	LeadsSourceOfflineServiceReferral  = "OFFLINE_SERVICE_REFERRAL"
)

var (
	LeadsFollowUpStatusMap = map[string]string{
		LeadsFollowUpStatusOnConsideration:  "On Consideration",
		LeadsFollowUpStatusNotYetFollowedUp: "Not Yet Followed Up",
		LeadsFollowUpStatusNoResponse:       "No Response",
	}

	LeadScoreMap = map[string]string{
		LeadScoreLow:    "Low",
		LeadScoreMedium: "Medium",
		LeadScoreHot:    "Hot",
	}

	LeadsSourceMap = map[string]string{
		LeadsSourceCustomerDatabase:        "Customer Database",
		LeadsSourceDigitalToyotaWeb:        "Digital Toyota Web",
		LeadsSourceDigitalMToyota:          "Digital M-Toyota",
		LeadsSourceDigitalOthers:           "Digital Others",
		LeadsSourceDigitalDealer:           "Digital Dealer",
		LeadsSourceOfflineEvent:            "Offline Event",
		LeadsSourceOfflineGathering:        "Offline Gathering",
		LeadsSourceOfflineCustomerReferral: "Offline Customer Referral",
		LeadsSourceOfflineWalkInOrCallIn:   "Offline Walk-in or Call-in",
		LeadsSourceOfflineServiceReferral:  "Offline Service Referral",
	}
)
