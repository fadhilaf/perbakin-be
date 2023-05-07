package usecase

import (
	// "github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/repository"
)

type ScorerUsecase interface {
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