package request

type LoginRequest struct {
	BaseRequest
	Email string
	Password string
}

func NewLoginRequest() *LoginRequest {
	return &LoginRequest{}
}