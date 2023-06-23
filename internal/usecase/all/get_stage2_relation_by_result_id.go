package usecase

import (
	"context"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/jackc/pgx/v5"
)

func (usecase *allUsecaseImpl) GetStage2RelationByResultId(req model.ByResultIdRequest) (model.Stage123456Relation, error) {
	var err error

	stage2, err := usecase.Store.GetStage2RelationByResultId(context.Background(), req.ResultID)
	if err == pgx.ErrNoRows {
		stage2, err := usecase.Store.CreateStage2(context.Background(), req.ResultID)
		if err != nil {
			return model.Stage123456Relation{}, err
		}

		return model.Stage123456Relation{ID: stage2.ID, ResultID: stage2.ResultID, IsTry2: stage2.IsTry2}, nil
	}

	return model.Stage123456Relation{ID: stage2.ID, ResultID: stage2.ResultID, IsTry2: stage2.IsTry2}, err
}
