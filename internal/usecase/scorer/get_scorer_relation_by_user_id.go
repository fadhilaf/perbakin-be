package usecase

import (
	"context"

	"github.com/FadhilAF/perbakin-be/internal/model"
)

func (usecase *scorerUsecaseImpl) GetScorerRelationByUserId(req model.UserByUserIdRequest) (model.OperatorRelation, error) {
	scorer, err := usecase.Store.GetScorerRelationByUserId(context.Background(), req.UserID)

	return model.OperatorRelation{ID: scorer.ID, UserID: scorer.UserID, ExamID: scorer.ExamID}, err
}
