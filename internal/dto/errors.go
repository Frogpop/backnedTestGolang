package dto

type ErrorResponse struct {
	Error   string `json:"error" example:"some error"`
	Details string `json:"details,omitempty" example:"some details"`
}
