package usecase

import (
	"context"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/jackc/pgx/v5"
)

func (usecase *allUsecaseImpl) GetResultRelationByShooterId(req model.ByShooterIdRequest) (model.ResultRelationAndStatus, error) {
	result, err := usecase.Store.GetResultRelationAndStatusByShooterId(context.Background(), req.ShooterID)
	if err == pgx.ErrNoRows {
		result, err := usecase.Store.CreateResult(context.Background(), req.ShooterID)
		if err != nil {
			return model.ResultRelationAndStatus{}, err
		}

		return model.ResultRelationAndStatus{ID: result.ID, ShooterID: result.ShooterID, Stage: string(result.Stage), Failed: result.Failed}, err
	}

	return model.ResultRelationAndStatus{ID: result.ID, ShooterID: result.ShooterID, Stage: string(result.Stage), Failed: result.Failed}, err
}
