package usecase

import (
	"context"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/jackc/pgx/v5"
)

func (usecase *allUsecaseImpl) GetStage4RelationByResultId(req model.ByResultIdRequest) (model.Stage123456RelationAndStatus, error) {
	var err error

	stage4, err := usecase.Store.GetStage4RelationByResultId(context.Background(), req.ResultID)
	if err == pgx.ErrNoRows {
		stage4, err := usecase.Store.CreateStage4(context.Background(), req.ResultID)
		if err != nil {
			return model.Stage123456RelationAndStatus{}, err
		}

		return model.Stage123456RelationAndStatus{ID: stage4.ID, ResultID: stage4.ResultID, IsTry2: stage4.IsTry2}, nil
	}

	return model.Stage123456RelationAndStatus{ID: stage4.ID, ResultID: stage4.ResultID, IsTry2: stage4.IsTry2}, err
}
