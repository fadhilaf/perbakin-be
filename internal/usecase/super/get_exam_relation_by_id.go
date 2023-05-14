package usecase

import (
	"context"

	"github.com/FadhilAF/perbakin-be/internal/model"
)

func (usecase *superUsecaseImpl) GetExamRelationById(req model.ByIdRequest) (model.ExamRelation, error) {
	exam, err := usecase.Store.GetExamRelationById(context.Background(), req.ID)

	return model.ExamRelation{ID: exam.ID, SuperID: exam.SuperID}, err
}
