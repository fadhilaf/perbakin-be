package usecase

import (
	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/repository"
)

type AdminUsecase interface {
	AdminLogin(model.LoginRequest) model.WebServiceResponse
}

var _ AdminUsecase = &adminUsecaseImpl{}

func NewDivisionUsecase(store repository.Store) AdminUsecase {
	return &adminUsecaseImpl{
		Store: store,
	}
}

type adminUsecaseImpl struct {
	repository.Store
}
