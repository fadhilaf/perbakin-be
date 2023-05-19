package usecase

import (
	"context"

	"github.com/FadhilAF/perbakin-be/internal/model"
)

func (usecase *adminSuperUsecaseImpl) GetScorerRelationById(req model.ByIdRequest) (model.OperatorRelation, error) {
	scorer, err := usecase.Store.GetScorerRelationById(context.Background(), req.ID)

	return model.OperatorRelation{ID: scorer.ID, UserID: scorer.UserID, ExamID: scorer.ExamID}, err
}
