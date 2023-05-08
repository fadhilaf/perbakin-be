package usecase

import (
	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/repository"
)

type SuperUsecase interface {
	SuperLogin(model.LoginRequest) model.WebServiceResponse
	GetSuperByUserId(model.GetByUserIdRequest) model.WebServiceResponse
	CreateAdmin(model.CreateUserRequest) model.WebServiceResponse
	GetAdmin(model.GetUserById) model.WebServiceResponse
	UpdateAdmin(model.UpdateUserRequest) model.WebServiceResponse
	DeleteAdmin(model.DeleteUserRequest) model.WebServiceResponse
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
