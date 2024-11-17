package domain

import (
	shareddomain "go-api-service-with-candi/pkg/shared/domain"
)

// RequestAuth model
type RequestAuth struct {
	ID    int    `json:"id"`
	Field string `json:"field"`
}

// Deserialize to db model
func (r *RequestAuth) Deserialize() (res shareddomain.Auth) {
	res.Field = r.Field
	return
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Name                 string `json:"name"`
	Email                string `json:"email"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"passwordConfirmation"`
}
