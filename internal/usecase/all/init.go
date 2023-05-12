package usecase

import (
	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/repository"
)

type AllUsecase interface {
	GetShootersByScorerId(model.ByScorerIdRequest) model.WebServiceResponse
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
