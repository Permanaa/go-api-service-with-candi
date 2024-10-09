package usecase

import (
	"context"
	"errors"
	"go-api-service-with-candi/internal/modules/auth/domain"

	"golang.org/x/crypto/bcrypt"
)

func (uc *authUsecaseImpl) Login(ctx context.Context, req *domain.LoginRequest) (resp domain.LoginResponse, err error) {
	emailPayload := req.Email
	passwordPayload := req.Password

	user, errGetUser := uc.repoSQL.AuthRepo().Find(ctx, &domain.FilterAuth{Email: emailPayload})
	if errGetUser != nil {
		return resp, errGetUser
	}

	errComparePassword := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(passwordPayload))
	if errComparePassword != nil {
		return resp, errors.New("wrong user email or password")
	}

	token, errToken := uc.GenerateGlobalToken(req)
	if errToken != nil {
		return resp, errToken
	}

	return token, nil
}
