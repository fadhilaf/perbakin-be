package usecase

import (
	"context"

	"github.com/FadhilAF/perbakin-be/internal/model"
)

func (usecase *allUsecaseImpl) GetShooterRelationById(req model.ByIdRequest) (model.ShooterRelation, error) {
	shooter, err := usecase.Store.GetShooterRelationById(context.Background(), req.ID)

	return model.ShooterRelation{ID: shooter.ID, ScorerID: shooter.ScorerID}, err
}
