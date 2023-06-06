package usecase

import (
	"context"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/jackc/pgx/v5"
)

func (usecase *allUsecaseImpl) GetResultRelationByShooterId(req model.ByShooterIdRequest) (model.ResultRelation, error) {
	result, err := usecase.Store.GetResultRelationByShooterId(context.Background(), req.ShooterID)
	if err == pgx.ErrNoRows {
		result, err := usecase.Store.CreateResult(context.Background(), req.ShooterID)
		if err != nil {
			return model.ResultRelation{}, err
		}

		return model.ResultRelation{ID: result.ID, ShooterID: result.ShooterID}, err
	}

	return model.ResultRelation{ID: result.ID, ShooterID: result.ShooterID}, err
}
