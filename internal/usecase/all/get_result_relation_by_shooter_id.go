package usecase

import (
	"context"

	"github.com/FadhilAF/perbakin-be/internal/model"
)

func (usecase *allUsecaseImpl) GetResultRelationByShooterId(req model.ByShooterIdRequest) (model.ResultRelation, error) {
	result, err := usecase.Store.GetResultRelationByShooterId(context.Background(), req.ShooterID)

	return model.ResultRelation{ID: result.ID, ShooterID: result.ShooterID}, err
}
