package usecase

import (
	"context"

	"github.com/FadhilAF/perbakin-be/internal/model"
)

func (usecase *allUsecaseImpl) GetSuperRelationByUserId(req model.UserByUserIdRequest) (model.SuperRelation, error) {
	super, err := usecase.Store.GetSuperRelationByUserId(context.Background(), req.UserID)

	return model.SuperRelation{ID: super.ID, UserID: super.UserID}, err
}
