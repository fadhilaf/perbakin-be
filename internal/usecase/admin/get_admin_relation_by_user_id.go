package usecase

import (
	"context"

	"github.com/FadhilAF/perbakin-be/internal/model"
)

func (usecase *adminUsecaseImpl) GetAdminRelationByUserId(req model.UserByUserIdRequest) (model.OperatorRelation, error) {
	admin, err := usecase.Store.GetAdminRelationByUserId(context.Background(), req.UserID)

	return model.OperatorRelation{ID: admin.ID, UserID: admin.UserID, ExamID: admin.ExamID}, err
}
