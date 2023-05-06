package usecase

import (
	"github.com/FadhilAF/perbakin-be/internal/repository"
)

type AuthUsecase interface {
}

var _ AuthUsecase = &authUsecaseImpl{}

func NewAuthUsecase(store repository.Store) AuthUsecase {
	return &authUsecaseImpl{
		Store: store,
	}
}

type authUsecaseImpl struct {
	repository.Store
}
