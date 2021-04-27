package jsonapi

type Error struct {
	Status int    `json:"status"`
	Title  string `json:"title"`
}

type ErrorResponse struct {
	Errors []*Error `json:"errors"`
}

func NewErrorResponse(status int, err error) *ErrorResponse {
	return &ErrorResponse{
		Errors: []*Error{{
			Title:  err.Error(),
			Status: status,
		}},
	}
}
