package request

type RegisterRequest struct {
	BaseRequest
	Username string
	Email string
	Phone string
	Password string
}

func NewRegisterRequest() *RegisterRequest {
	return &RegisterRequest{}
}

