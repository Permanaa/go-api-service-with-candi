// Code generated by candi v1.17.16.

package usecase

import (
	"context"
	"go-api-service-with-candi/internal/modules/auth/domain"
	"go-api-service-with-candi/pkg/shared/repository"
	"go-api-service-with-candi/pkg/shared/usecase/common"
	"net/http"

	"github.com/golangid/candi/codebase/factory/dependency"
)

// AuthUsecase abstraction
type AuthUsecase interface {
	GenerateAnonymousToken() (token string, err error)
	ValidateAnonymousToken(next http.Handler) http.Handler
	Login(ctx context.Context, req *domain.LoginRequest) (resp domain.LoginResponse, err error)
}

type authUsecaseImpl struct {
	deps          dependency.Dependency
	sharedUsecase common.Usecase
	repoSQL       repository.RepoSQL
	// repoMongo     repository.RepoMongo
}

// NewAuthUsecase usecase impl constructor
func NewAuthUsecase(deps dependency.Dependency) (AuthUsecase, func(sharedUsecase common.Usecase)) {
	uc := &authUsecaseImpl{
		deps:    deps,
		repoSQL: repository.GetSharedRepoSQL(),
		// repoMongo: repository.GetSharedRepoMongo(),

	}
	return uc, func(sharedUsecase common.Usecase) {
		uc.sharedUsecase = sharedUsecase
	}
}
