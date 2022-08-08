package response

import "net/http"

type LoginResponse struct {
	Token   string
	Message string
	Code    int
}

func NewLoginSucess(token string) *LoginResponse {
	return &LoginResponse{
		Token:   token,
		Message: "success login",
		Code:    http.StatusOK,
	}
}
