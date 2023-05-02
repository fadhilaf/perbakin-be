package usecase

import (
	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/repository"
)

type SuperUsecase interface {
	SuperLogin(model.LoginRequest) model.WebServiceResponse
}

var _ SuperUsecase = &superUsecaseImpl{}

func NewSuperUsecase(store repository.Store) SuperUsecase {
	return &superUsecaseImpl{
		Store: store,
	}
}

type superUsecaseImpl struct {
	repository.Store
}
