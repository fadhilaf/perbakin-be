package usecase

import (
	"context"

	"github.com/FadhilAF/perbakin-be/internal/model"
)

func (usecase *superUsecaseImpl) GetAdminRelationById(req model.ByIdRequest) (model.OperatorRelation, error) {
	admin, err := usecase.Store.GetAdminRelationById(context.Background(), req.ID)

	return model.OperatorRelation{ID: admin.ID, UserID: admin.UserID, ExamID: admin.ExamID}, err
}
