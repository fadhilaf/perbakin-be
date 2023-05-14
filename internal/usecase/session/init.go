package usecase

import (
	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/repository"
)

type SessionUsecase interface {
	GetExamRelationById(model.ByIdRequest) (model.ExamRelation, error)

	GetScorerRelationByUserId(model.UserByUserIdRequest) (model.OperatorRelation, error)
}

var _ SessionUsecase = &sessionUsecaseImpl{}

func NewSessionUsecase(store repository.Store) SessionUsecase {
	return &sessionUsecaseImpl{
		Store: store,
	}
}

type sessionUsecaseImpl struct {
	repository.Store
}
