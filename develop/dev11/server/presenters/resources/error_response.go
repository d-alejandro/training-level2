package resources

/*
ErrorResponse structure
*/
type ErrorResponse struct {
	Error struct {
		Message string `json:"message"`
		Status  string `json:"status"`
	} `json:"error"`
}

/*
NewErrorResponse constructor
*/
func NewErrorResponse(message, status string) *ErrorResponse {
	return &ErrorResponse{
		Error: struct {
			Message string `json:"message"`
			Status  string `json:"status"`
		}{
			Message: message,
			Status:  status,
		},
	}
}
