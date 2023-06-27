package usecase

import (
	"context"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/jackc/pgx/v5"
)

func (usecase *allUsecaseImpl) GetStage6RelationByResultId(req model.ByResultIdRequest) (model.Stage123456RelationAndStatus, error) {
	var err error

	stage6, err := usecase.Store.GetStage6RelationByResultId(context.Background(), req.ResultID)
	if err == pgx.ErrNoRows {
		stage6, err := usecase.Store.CreateStage6(context.Background(), req.ResultID)
		if err != nil {
			return model.Stage123456RelationAndStatus{}, err
		}

		return model.Stage123456RelationAndStatus{ID: stage6.ID, ResultID: stage6.ResultID, IsTry2: stage6.IsTry2}, nil
	}

	return model.Stage123456RelationAndStatus{ID: stage6.ID, ResultID: stage6.ResultID, IsTry2: stage6.IsTry2}, err
}
