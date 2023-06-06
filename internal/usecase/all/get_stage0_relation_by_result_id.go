package usecase

import (
	"context"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/jackc/pgx/v5"
)

func (usecase *allUsecaseImpl) GetStage0RelationByResultId(req model.ByResultIdRequest) (model.Stage0Relation, error) {
	stage0, err := usecase.Store.GetStage0RelationByResultId(context.Background(), req.ResultID)
	if err == pgx.ErrNoRows {
		stage0, err := usecase.Store.CreateStage0(context.Background(), req.ResultID)
		if err != nil {
			return model.Stage0Relation{}, err
		}

		return model.Stage0Relation{ID: stage0.ID, ResultID: stage0.ResultID}, err
	}

	return model.Stage0Relation{ID: stage0.ID, ResultID: stage0.ResultID}, err
}
