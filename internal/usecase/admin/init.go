package usecase

import (
	// "github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/repository"
)

type AdminUsecase interface {
}

var _ AdminUsecase = &adminUsecaseImpl{}

func NewAdminUsecase(store repository.Store) AdminUsecase {
	return &adminUsecaseImpl{
		Store: store,
	}
}

type adminUsecaseImpl struct {
	repository.Store
}
