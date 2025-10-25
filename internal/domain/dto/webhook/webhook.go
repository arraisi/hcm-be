package webhook

// Headers represents the structured webhook headers
type Headers struct {
	ContentType string
	APIKey      string
	Signature   string
	EventID     string
	Timestamp   string
}

// Response represents the response structure for webhook
type Response struct {
	Data    ResponseData `json:"data"`
	Message string       `json:"message"`
}

// ResponseData represents the data part of webhook response
type ResponseData struct {
	EventID string `json:"eventId"`
	Status  string `json:"status"`
}

// ErrorResponse represents the error response structure
type ErrorResponse struct {
	Data    map[string]interface{} `json:"data"`
	Message string                 `json:"message"`
}
