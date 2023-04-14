package request

type CreateUserRequest struct {
	Name   string `json:"name" validate:"required"`
	Gender string `json:"gender" validate:"required"`
}

type UserRequest struct {
	UserId string `json:"userId" validate:"required"`
}
