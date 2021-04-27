package jsonapi

type Error struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type ErrorResponse struct {
	Errors []*Error `json:"errors"`
}

func NewErrorResponse(status int, err error) *ErrorResponse {
	return &ErrorResponse{
		Errors: []*Error{{
			Message: err.Error(),
			Status:  status,
		}},
	}
}
