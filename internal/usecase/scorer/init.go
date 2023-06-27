package usecase

import (
	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/repository"
)

type ScorerUsecase interface {
	ScorerLogin(model.LoginRequest) model.WebServiceResponse

	GetScorerRelationByUserId(model.UserByUserIdRequest) (model.OperatorRelationAndStatus, error)

	GetScorerByUserId(model.UserByUserIdRequest) model.WebServiceResponse

	UpdateStage0Series(model.UpdateStage0SeriesRequest) model.WebServiceResponse
	UpdateStage0NextSeries(model.ByIdRequest) model.WebServiceResponse
	UpdateStage0Checkmarks(model.UpdateStage0CheckmarksRequest) model.WebServiceResponse
	UpdateStage0Finish(model.UpdateStageFinishRequest) model.WebServiceResponse

	Stage1try1Finished(model.ByIdRequest) (bool, error)
	UpdateStage1No(model.UpdateStage123456NoRequest) model.WebServiceResponse
	UpdateStage1NextNo(model.ByIdAndTryRequest) model.WebServiceResponse
	UpdateStage1Checkmarks(model.UpdateStage123456CheckmarksRequest) model.WebServiceResponse
	UpdateStage1Finish(model.UpdateStage123456FinishRequest) model.WebServiceResponse

	Stage2try1Finished(model.ByIdRequest) (bool, error)
	UpdateStage2No(model.UpdateStage123456NoRequest) model.WebServiceResponse
	UpdateStage2NextNo(model.ByIdAndTryRequest) model.WebServiceResponse
	UpdateStage2Checkmarks(model.UpdateStage123456CheckmarksRequest) model.WebServiceResponse
	UpdateStage2Finish(model.UpdateStage123456FinishRequest) model.WebServiceResponse

	Stage3try1Finished(model.ByIdRequest) (bool, error)
	UpdateStage3No(model.UpdateStage123456NoRequest) model.WebServiceResponse
	UpdateStage3NextNo(model.ByIdAndTryRequest) model.WebServiceResponse
	UpdateStage3Checkmarks(model.UpdateStage123456CheckmarksRequest) model.WebServiceResponse
	UpdateStage3Finish(model.UpdateStage123456FinishRequest) model.WebServiceResponse

	Stage4try1Finished(model.ByIdRequest) (bool, error)
	UpdateStage4No(model.UpdateStage123456NoRequest) model.WebServiceResponse
	UpdateStage4NextNo(model.ByIdAndTryRequest) model.WebServiceResponse
	UpdateStage4Checkmarks(model.UpdateStage123456CheckmarksRequest) model.WebServiceResponse
	UpdateStage4Finish(model.UpdateStage123456FinishRequest) model.WebServiceResponse

	Stage5try1Finished(model.ByIdRequest) (bool, error)
	UpdateStage5No(model.UpdateStage123456NoRequest) model.WebServiceResponse
	UpdateStage5NextNo(model.ByIdAndTryRequest) model.WebServiceResponse
	UpdateStage5Checkmarks(model.UpdateStage123456CheckmarksRequest) model.WebServiceResponse
	UpdateStage5Finish(model.UpdateStage123456FinishRequest) model.WebServiceResponse
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
