package usecase

import (
	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/repository"
)

type ScorerUsecase interface {
	ScorerLogin(model.LoginRequest) model.WebServiceResponse

	GetScorerRelationByUserId(model.UserByUserIdRequest) (model.OperatorRelation, error)

	GetScorerByUserId(model.UserByUserIdRequest) model.WebServiceResponse

	GetResultStageById(model.ByIdRequest) (string, error)
	UpdateStage0Series(model.UpdateStage0SeriesRequest) model.WebServiceResponse
	UpdateStage0NextSeries(model.ByIdRequest) model.WebServiceResponse
	UpdateStage0Checkmarks(model.UpdateStage0CheckmarksRequest) model.WebServiceResponse
	UpdateStage0Finish(model.UpdateStage0FinishRequest) model.WebServiceResponse
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
