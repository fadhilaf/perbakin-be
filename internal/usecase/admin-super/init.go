package usecase

import (
	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/repository"
)

type AdminSuperUsecase interface {
	GetAdminSuperByUserId(req model.GetByUserIdRequest) model.WebServiceResponse
	CreateScorer(req model.CreateUserRequest) model.WebServiceResponse
	GetScorer(req model.GetUserById) model.WebServiceResponse
	GetAllScorers() model.WebServiceResponse
	UpdateScorer(req model.UpdateUserRequest) model.WebServiceResponse
	DeleteScorer(req model.DeleteUserRequest) model.WebServiceResponse
}

var _ AdminSuperUsecase = &adminSuperUsecaseImpl{}

func NewAdminSuperUsecase(store repository.Store) AdminSuperUsecase {
	return &adminSuperUsecaseImpl{
		Store: store,
	}
}

type adminSuperUsecaseImpl struct {
	repository.Store
}
