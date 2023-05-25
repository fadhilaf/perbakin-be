package usecase

import (
	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/repository"
)

type AllUsecase interface {
	GetShooterRelationById(model.ByIdRequest) (model.ShooterRelation, error)

	GetShooterById(model.ByIdRequest) model.WebServiceResponse
	GetShootersByScorerId(model.ByScorerIdRequest) model.WebServiceResponse

	CreateResult(model.ByShooterIdRequest) model.WebServiceResponse

	GetResultRelationByShooterId(model.ByShooterIdRequest) (model.ResultRelation, error)
	GetResultById(model.ByIdRequest) model.WebServiceResponse

	CreateStage0(model.ByResultIdRequest) model.WebServiceResponse

	GetStage0RelationByResultId(model.ByResultIdRequest) (model.Stage0Relation, error)
	GetStage0ById(model.ByIdRequest) model.WebServiceResponse
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
