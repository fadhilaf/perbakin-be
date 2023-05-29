package usecase

import (
	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/repository"
)

type AdminUsecase interface {
	AdminLogin(model.LoginRequest) model.WebServiceResponse
	GetAdminByUserId(model.UserByUserIdRequest) model.WebServiceResponse

	GetAdminExamRelationByUserId(model.UserByUserIdRequest) (model.OperatorAndExamRelation, error)
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
