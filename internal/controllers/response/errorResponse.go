package response

type ErrorResponse struct {
	Error string
}

func NewErrorResponse(message string) ErrorResponse {
	return ErrorResponse{
		Error: message,
	}
}
