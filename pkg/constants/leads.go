package constants

const (
	LeadsFollowUpStatusOnConsideration  = "ON_CONSIDERATION"
	LeadsFollowUpStatusNotYetFollowedUp = "NOT_YET_FOLLOWED_UP"
	LeadsFollowUpStatusNoResponse       = "NO_RESPONSE"
)

var (
	LeadsFollowUpStatusMap = map[string]string{
		LeadsFollowUpStatusOnConsideration:  "On Consideration",
		LeadsFollowUpStatusNotYetFollowedUp: "Not Yet Followed Up",
		LeadsFollowUpStatusNoResponse:       "No Response",
	}
)
