package usecase

import (
	"context"

	"github.com/FadhilAF/perbakin-be/internal/model"
)

func (usecase *allUsecaseImpl) GetStage0RelationByResultId(req model.ByResultIdRequest) (model.Stage0Relation, error) {
	stage0, err := usecase.Store.GetStage0RelationByResultId(context.Background(), req.ResultID)

	return model.Stage0Relation{ID: stage0.ID, ResultID: stage0.ResultID}, err
}
