package engine

type CreateRoLeadsRequest struct {
	ForceUpdate bool `json:"force_update"`
}

var (
	CustomerResponseSent          = "SENT"
	CustomerResponseNotSent       = "NOT_SENT"
	CustomerResponseInterested    = "INTERESTED"
	CustomerResponseNotInterested = "NOT_INTERESTED"
)
