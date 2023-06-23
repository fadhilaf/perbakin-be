package usecase

import (
	"context"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/jackc/pgx/v5"
)

func (usecase *allUsecaseImpl) GetStage1RelationByResultId(req model.ByResultIdRequest) (model.Stage123456Relation, error) {
	var err error

	stage1, err := usecase.Store.GetStage1RelationByResultId(context.Background(), req.ResultID)
	if err == pgx.ErrNoRows {
		stage1, err := usecase.Store.CreateStage1(context.Background(), req.ResultID)
		if err != nil {
			return model.Stage123456Relation{}, err
		}

		return model.Stage123456Relation{ID: stage1.ID, ResultID: stage1.ResultID, IsTry2: stage1.IsTry2}, nil
	}

	return model.Stage123456Relation{ID: stage1.ID, ResultID: stage1.ResultID, IsTry2: stage1.IsTry2}, err
}
