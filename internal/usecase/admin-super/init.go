package usecase

import (
	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/repository"
)

type AdminSuperUsecase interface {
	CreateScorer(req model.CreateOperatorRequest) model.WebServiceResponse
	GetAllScorers() model.WebServiceResponse
	GetScorersByExamId(req model.GetOperatorsByExamIdRequest) model.WebServiceResponse
	GetScorerById(req model.OperatorByIdRequest) model.WebServiceResponse
	UpdateScorer(req model.UpdateOperatorRequest) model.WebServiceResponse
	DeleteScorer(req model.OperatorByIdRequest) model.WebServiceResponse
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
