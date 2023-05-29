package usecase

import (
	"context"

	"github.com/FadhilAF/perbakin-be/internal/model"
)

func (usecase *adminUsecaseImpl) GetAdminExamRelationByUserId(req model.UserByUserIdRequest) (model.OperatorAndExamRelation, error) {
	adminExam, err := usecase.Store.GetAdminExamRelationByUserId(context.Background(), req.UserID)

	return model.OperatorAndExamRelation{ID: adminExam.ID, UserID: adminExam.UserID, ExamID: adminExam.ExamID, SuperID: adminExam.SuperID}, err
}
