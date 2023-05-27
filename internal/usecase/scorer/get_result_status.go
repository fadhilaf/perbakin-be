package usecase

import (
	"context"

	"github.com/FadhilAF/perbakin-be/internal/model"
)

func (usecase *scorerUsecaseImpl) GetResultStatusById(req model.ByIdRequest) (model.ResultStatus, error) {
	status, err := usecase.Store.GetResultStatusById(context.Background(), req.ID)

	return model.ResultStatus{Failed: status.Failed, Stage: string(status.Stage)}, err
}
