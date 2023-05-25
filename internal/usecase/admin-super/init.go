package usecase

import (
	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/repository"
)

type AdminSuperUsecase interface {
	CreateScorer(model.CreateOperatorRequest) model.WebServiceResponse
	GetAllScorers() model.WebServiceResponse
	GetScorersByExamId(model.ByExamIdRequest) model.WebServiceResponse
	GetScorerRelationById(model.ByIdRequest) (model.OperatorRelation, error)
	GetScorerById(model.ByIdRequest) model.WebServiceResponse
	UpdateScorer(model.UpdateOperatorRequest) model.WebServiceResponse
	DeleteScorer(model.UserByUserIdRequest) model.WebServiceResponse

	CreateShooter(model.CreateShooterRequest) model.WebServiceResponse
	GetAllShooters() model.WebServiceResponse
	GetShootersByExamId(model.ByExamIdRequest) model.WebServiceResponse
	UpdateShooter(model.UpdateShooterRequest) model.WebServiceResponse
	DeleteShooter(model.ByIdRequest) model.WebServiceResponse

	UpdateResultByShooterId(model.UpdateResultByShooterIdRequest) model.WebServiceResponse
	DeleteResultByShooterId(model.ByShooterIdRequest) model.WebServiceResponse
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
