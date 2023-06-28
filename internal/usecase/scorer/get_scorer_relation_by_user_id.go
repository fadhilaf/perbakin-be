package usecase

import (
	"context"

	"github.com/FadhilAF/perbakin-be/internal/model"
)

func (usecase *scorerUsecaseImpl) GetScorerRelationByUserId(req model.UserByUserIdRequest) (model.ScorerRelationAndStatus, error) {
	scorer, err := usecase.Store.GetScorerRelationByUserId(context.Background(), req.UserID)

	return model.ScorerRelationAndStatus{ID: scorer.ID, UserID: scorer.UserID, ExamID: scorer.ExamID, Active: scorer.Active}, err
}
