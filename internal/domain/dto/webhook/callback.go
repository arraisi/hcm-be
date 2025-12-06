package webhook

// CallbackTamResponse represents the callback payload to DMS for TAM response
type CallbackTamResponse struct {
	EventID    string `json:"event_ID"`
	DocumentID string `json:"document_id"`
	Status     string `json:"status"`  // "success" or "failed"
	Code       string `json:"code"`    // HTTP status code as string
	Message    string `json:"massage"` // Note: typo in DMS API (massage instead of message)
	Data       string `json:"data"`    // Additional data, usually empty
}

// NewSuccessCallback creates a success callback response
func NewSuccessCallback(eventID, documentID, message string) CallbackTamResponse {
	return CallbackTamResponse{
		EventID:    eventID,
		DocumentID: documentID,
		Status:     "success",
		Code:       "200",
		Message:    message,
		Data:       "",
	}
}

// NewFailureCallback creates a failure callback response
func NewFailureCallback(eventID, documentID string, err error) CallbackTamResponse {
	message := "Unknown error"
	if err != nil {
		message = err.Error()
	}

	return CallbackTamResponse{
		EventID:    eventID,
		DocumentID: documentID,
		Status:     "failed",
		Code:       "500",
		Message:    message,
		Data:       "",
	}
}
