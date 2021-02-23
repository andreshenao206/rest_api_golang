package httpmodels

type UserPublicRequest struct {
	Email string `json:"email" binding:"required"`
}
type ErrorRespose struct {
	Status bool
	Msg    string
}
