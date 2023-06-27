package usecase

import (
	"context"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/jackc/pgx/v5"
)

func (usecase *allUsecaseImpl) GetStage5RelationByResultId(req model.ByResultIdRequest) (model.Stage123456RelationAndStatus, error) {
	var err error

	stage5, err := usecase.Store.GetStage5RelationByResultId(context.Background(), req.ResultID)
	if err == pgx.ErrNoRows {
		stage5, err := usecase.Store.CreateStage5(context.Background(), req.ResultID)
		if err != nil {
			return model.Stage123456RelationAndStatus{}, err
		}

		return model.Stage123456RelationAndStatus{ID: stage5.ID, ResultID: stage5.ResultID, IsTry2: stage5.IsTry2}, nil
	}

	return model.Stage123456RelationAndStatus{ID: stage5.ID, ResultID: stage5.ResultID, IsTry2: stage5.IsTry2}, err
}
