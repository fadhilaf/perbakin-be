package usecase

import (
	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/repository"
)

type ScorerUsecase interface {
	ScorerLogin(model.LoginRequest) model.WebServiceResponse

	GetScorerRelationByUserId(model.UserByUserIdRequest) (model.OperatorRelationAndStatus, error)

	GetScorerByUserId(model.UserByUserIdRequest) model.WebServiceResponse

	GetResultStatusById(model.ByIdRequest) (model.ResultStatus, error)

	UpdateStage0Series(model.UpdateStage0SeriesRequest) model.WebServiceResponse
	UpdateStage0NextSeries(model.ByIdRequest) model.WebServiceResponse
	UpdateStage0Checkmarks(model.UpdateStage0CheckmarksRequest) model.WebServiceResponse
	UpdateStage0Finish(model.UpdateStageFinishRequest) model.WebServiceResponse

	CreateStage1try2(model.ByIdRequest) model.WebServiceResponse
	UpdateStage1No(model.UpdateStage123456NoRequest) model.WebServiceResponse
	UpdateStage1NextNo(model.ByIdAndTryRequest) model.WebServiceResponse
	UpdateStage1Checkmarks(model.UpdateStage123456CheckmarksRequest) model.WebServiceResponse
	UpdateStage1Finish(model.UpdateStage123456FinishRequest) model.WebServiceResponse
}

var _ ScorerUsecase = &scorerUsecaseImpl{}

func NewScorerUsecase(store repository.Store) ScorerUsecase {
	return &scorerUsecaseImpl{
		Store: store,
	}
}

type scorerUsecaseImpl struct {
	repository.Store
}
