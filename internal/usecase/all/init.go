package usecase

import (
	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/repository"
)

type AllUsecase interface {
	GetSuperRelationByUserId(model.UserByUserIdRequest) (*model.SuperRelation, error)

	GetExamRelationById(model.ByIdRequest) (*model.ExamRelation, error)

	GetAdminRelationByUserId(model.UserByUserIdRequest) (*model.OperatorRelation, error)
	GetScorerRelationByUserId(model.UserByUserIdRequest) (*model.OperatorRelation, error)
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
