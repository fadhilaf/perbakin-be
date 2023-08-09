package usecase

import (
	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/repository"
)

type AllUsecase interface {
	GetScorerById(model.ByIdRequest) model.WebServiceResponse
	UpdateScorer(model.UpdateOperatorRequest) model.WebServiceResponse
	UpdateScorerImage(model.UpdateImageRequest) model.WebServiceResponse

	GetShooterRelationById(model.ByIdRequest) (model.ShooterRelation, error)

	GetShooterById(model.ByIdRequest) model.WebServiceResponse
	GetShootersByScorerId(model.ByScorerIdRequest) model.WebServiceResponse

	CreateResult(model.ByShooterIdRequest) model.WebServiceResponse

	GetResultRelationByShooterId(model.ByShooterIdRequest) (model.ResultRelationAndStatus, error)
	GetResultById(model.ByIdRequest) model.WebServiceResponse
	GetResultsByScorerId(model.ByScorerIdRequest) model.WebServiceResponse

	CreateStage0(model.ByResultIdRequest) model.WebServiceResponse

	GetStage0RelationByResultId(model.ByResultIdRequest) (model.Stage0Relation, error)
	GetStage0ById(model.ByIdRequest) model.WebServiceResponse

	CreateStage1(model.ByResultIdRequest) model.WebServiceResponse

	GetStage1RelationByResultId(model.ByResultIdRequest) (model.Stage123456RelationAndStatus, error)
	GetStage1ById(model.ByIdRequest) model.WebServiceResponse
	CreateStage1try2(model.ByIdRequest) model.WebServiceResponse

	CreateStage2(model.ByResultIdRequest) model.WebServiceResponse

	GetStage2RelationByResultId(model.ByResultIdRequest) (model.Stage123456RelationAndStatus, error)
	GetStage2ById(model.ByIdRequest) model.WebServiceResponse
	CreateStage2try2(req model.ByIdRequest) model.WebServiceResponse

	CreateStage3(model.ByResultIdRequest) model.WebServiceResponse

	GetStage3RelationByResultId(model.ByResultIdRequest) (model.Stage123456RelationAndStatus, error)
	GetStage3ById(model.ByIdRequest) model.WebServiceResponse
	CreateStage3try2(model.ByIdRequest) model.WebServiceResponse

	CreateStage4(model.ByResultIdRequest) model.WebServiceResponse

	GetStage4RelationByResultId(model.ByResultIdRequest) (model.Stage123456RelationAndStatus, error)
	GetStage4ById(model.ByIdRequest) model.WebServiceResponse
	CreateStage4try2(req model.ByIdRequest) model.WebServiceResponse

	CreateStage5(model.ByResultIdRequest) model.WebServiceResponse

	GetStage5RelationByResultId(model.ByResultIdRequest) (model.Stage123456RelationAndStatus, error)
	GetStage5ById(model.ByIdRequest) model.WebServiceResponse
	CreateStage5try2(req model.ByIdRequest) model.WebServiceResponse

	CreateStage6(model.ByResultIdRequest) model.WebServiceResponse

	GetStage6RelationByResultId(model.ByResultIdRequest) (model.Stage123456RelationAndStatus, error)
	GetStage6ById(model.ByIdRequest) model.WebServiceResponse
	CreateStage6try2(req model.ByIdRequest) model.WebServiceResponse
}

var _ AllUsecase = &allUsecaseImpl{}

func NewAllUsecase(store repository.Store) AllUsecase {
	return &allUsecaseImpl{
		Store: store,
	}
}

type allUsecaseImpl struct {
	repository.Store
}
