// Code generated by candi v1.17.16.

package usecase

import (
	"sync"

	// @candi:usecaseImport
	authusecase "go-api-service-with-candi/internal/modules/auth/usecase"
	"go-api-service-with-candi/pkg/shared/usecase/common"

	"github.com/golangid/candi/codebase/factory/dependency"
)

type (
	// Usecase unit of work for all usecase in modules
	Usecase interface {
		// @candi:usecaseMethod
		Auth() authusecase.AuthUsecase
	}

	usecaseUow struct {
		// @candi:usecaseField
		authusecase.AuthUsecase
	}
)

var usecaseInst *usecaseUow
var once sync.Once

// SetSharedUsecase set singleton usecase unit of work instance
func SetSharedUsecase(deps dependency.Dependency) {
	once.Do(func() {
		usecaseInst = new(usecaseUow)
		var setSharedUsecaseFuncs []func(common.Usecase)
		var setSharedUsecaseFunc func(common.Usecase)

		// @candi:usecaseCommon
		usecaseInst.AuthUsecase, setSharedUsecaseFunc = authusecase.NewAuthUsecase(deps)
		setSharedUsecaseFuncs = append(setSharedUsecaseFuncs, setSharedUsecaseFunc)

		sharedUsecase := common.SetCommonUsecase(usecaseInst)
		for _, setFunc := range setSharedUsecaseFuncs {
			setFunc(sharedUsecase)
		}
	})
}

// GetSharedUsecase get usecase unit of work instance
func GetSharedUsecase() Usecase {
	return usecaseInst
}

// @candi:usecaseImplementation
func (uc *usecaseUow) Auth() authusecase.AuthUsecase {
	return uc.AuthUsecase
}

