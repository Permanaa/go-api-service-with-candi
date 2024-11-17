package usecase

import (
	"context"
	"errors"
	"go-api-service-with-candi/internal/modules/auth/domain"
	sharedDomain "go-api-service-with-candi/pkg/shared/domain"

	"golang.org/x/crypto/bcrypt"
)

func (uc *authUsecaseImpl) Register(ctx context.Context, req *domain.RegisterRequest) (resp domain.RegisterResponse, err error) {
	if req.Password != req.PasswordConfirmation {
		return resp, errors.New("password is not the same as password confirmation")
	}

	hashPassword, errHashPassword := bcrypt.GenerateFromPassword([]byte(req.Password), 14)
	stringHashPassword := string(hashPassword)

	if errHashPassword != nil {
		return resp, errHashPassword
	}

	user, errGetUser := uc.repoSQL.AuthRepo().Find(ctx, &domain.FilterAuth{Email: req.Email})
	if errGetUser == nil && user.DeletedAt.Valid {
		return resp, errors.New("email has been used")
	}

	if err := uc.repoSQL.AuthRepo().Save(ctx, &sharedDomain.User{
		Email:    req.Email,
		Name:     req.Name,
		Password: stringHashPassword,
	}); err != nil {
		return resp, errors.New("Nomor HP atau email telah digunakan")
	}

	resp = domain.RegisterResponse{
		Name:  req.Name,
		Email: req.Email,
	}

	return resp, nil
}
