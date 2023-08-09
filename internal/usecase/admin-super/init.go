package usecase

import (
	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/repository"
)

type AdminSuperUsecase interface {
	GetAdminById(model.ByIdRequest) model.WebServiceResponse
	UpdateAdmin(model.UpdateOperatorRequest) model.WebServiceResponse

	CreateScorer(model.CreateScorerRequest) model.WebServiceResponse
	GetAllScorers() model.WebServiceResponse
	GetScorersByExamId(model.ByExamIdRequest) model.WebServiceResponse
	GetScorerRelationById(model.ByIdRequest) (model.OperatorRelation, error)
	DeleteScorer(model.UserByUserIdRequest) model.WebServiceResponse

	CreateShooter(model.CreateShooterRequest) model.WebServiceResponse
	GetAllShooters() model.WebServiceResponse
	GetShootersByExamId(model.ByExamIdRequest) model.WebServiceResponse
	UpdateShooter(model.UpdateShooterRequest) model.WebServiceResponse
	UpdateShooterImage(model.UpdateImageRequest) model.WebServiceResponse
	DeleteShooter(model.ByIdRequest) model.WebServiceResponse

	GetResultsByExamId(model.ByExamIdRequest) model.WebServiceResponse
	UpdateResult(model.UpdateResultRequest) model.WebServiceResponse
	UpdateResultStage(model.UpdateResultStageRequest) bool
	DeleteResult(model.ByIdRequest) model.WebServiceResponse

	UpdateStage0(model.UpdateStage0Request) model.WebServiceResponse
	UpdateStage0Signs(model.UpdateStageSignsRequest) model.WebServiceResponse
	DeleteStage0(model.ByIdRequest) model.WebServiceResponse
	FinishStage0(model.ByIdRequest) error

	UpdateStage1try1(model.UpdateStage13try1Request) model.WebServiceResponse
	UpdateStage1try2(model.UpdateStage13try2Request) model.WebServiceResponse
	UpdateStage1Signs(model.UpdateStageSignsRequest) model.WebServiceResponse
	DeleteStage1(model.ByIdRequest) model.WebServiceResponse
	DeleteStage1try2(model.ByIdRequest) model.WebServiceResponse
	FinishStage1(model.ByIdRequest) error

	UpdateStage2try1(model.UpdateStage246try1Request) model.WebServiceResponse
	UpdateStage2try2(model.UpdateStage246try2Request) model.WebServiceResponse
	UpdateStage2Signs(model.UpdateStageSignsRequest) model.WebServiceResponse
	DeleteStage2(model.ByIdRequest) model.WebServiceResponse
	DeleteStage2try2(model.ByIdRequest) model.WebServiceResponse
	FinishStage2(model.ByIdRequest) error

	UpdateStage3try1(model.UpdateStage13try1Request) model.WebServiceResponse
	UpdateStage3try2(model.UpdateStage13try2Request) model.WebServiceResponse
	UpdateStage3Signs(model.UpdateStageSignsRequest) model.WebServiceResponse
	DeleteStage3(model.ByIdRequest) model.WebServiceResponse
	DeleteStage3try2(model.ByIdRequest) model.WebServiceResponse
	FinishStage3(model.ByIdRequest) error

	UpdateStage4try1(model.UpdateStage246try1Request) model.WebServiceResponse
	UpdateStage4try2(model.UpdateStage246try2Request) model.WebServiceResponse
	UpdateStage4Signs(model.UpdateStageSignsRequest) model.WebServiceResponse
	DeleteStage4(model.ByIdRequest) model.WebServiceResponse
	DeleteStage4try2(model.ByIdRequest) model.WebServiceResponse
	FinishStage4(model.ByIdRequest) error

	UpdateStage5try1(model.UpdateStage5try1Request) model.WebServiceResponse
	UpdateStage5try2(model.UpdateStage5try2Request) model.WebServiceResponse
	UpdateStage5Signs(model.UpdateStageSignsRequest) model.WebServiceResponse
	DeleteStage5(model.ByIdRequest) model.WebServiceResponse
	DeleteStage5try2(model.ByIdRequest) model.WebServiceResponse
	FinishStage5(model.ByIdRequest) error

	UpdateStage6try1(model.UpdateStage246try1Request) model.WebServiceResponse
	UpdateStage6try2(model.UpdateStage246try2Request) model.WebServiceResponse
	UpdateStage6Signs(model.UpdateStageSignsRequest) model.WebServiceResponse
	DeleteStage6(model.ByIdRequest) model.WebServiceResponse
	DeleteStage6try2(model.ByIdRequest) model.WebServiceResponse
	FinishStage6(model.ByIdRequest) error
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
