package usecase

import (
	"context"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

func (usecase *allUsecaseImpl) GetStage1RelationByResultId(req model.ByResultIdRequest) (model.Stage1Relation, error) {
	var err error

	stage1, err := usecase.Store.GetStage1RelationByResultId(context.Background(), req.ResultID)
	if err == pgx.ErrNoRows {
		stage1, err := usecase.Store.CreateStage1(context.Background(), req.ResultID)
		if err != nil {
			return model.Stage1Relation{}, err
		}

		return model.Stage1Relation{ID: stage1.ID, ResultID: stage1.ResultID, Try1ID: stage1.Try1ID, Try2ID: pgtype.UUID{}, IsTry2: stage1.IsTry2}, nil
	}

	return model.Stage1Relation{ID: stage1.ID, ResultID: stage1.ResultID, Try1ID: stage1.Try1ID, Try2ID: stage1.Try2ID, IsTry2: stage1.IsTry2}, err
}
