package usecase

import (
	"context"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/jackc/pgx/v5"
)

func (usecase *allUsecaseImpl) GetStage3RelationByResultId(req model.ByResultIdRequest) (model.Stage123456RelationAndStatus, error) {
	var err error

	stage3, err := usecase.Store.GetStage3RelationByResultId(context.Background(), req.ResultID)
	if err == pgx.ErrNoRows {
		stage3, err := usecase.Store.CreateStage3(context.Background(), req.ResultID)
		if err != nil {
			return model.Stage123456RelationAndStatus{}, err
		}

		return model.Stage123456RelationAndStatus{ID: stage3.ID, ResultID: stage3.ResultID, IsTry2: stage3.IsTry2}, nil
	}

	return model.Stage123456RelationAndStatus{ID: stage3.ID, ResultID: stage3.ResultID, IsTry2: stage3.IsTry2}, err
}
