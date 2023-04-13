package request

type CreateUserRequest struct {
	Name   string `json:"name"`
	Gender string `json:"gender"`
}
