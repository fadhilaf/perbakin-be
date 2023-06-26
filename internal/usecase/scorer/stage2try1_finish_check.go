package usecase

import (
	"context"

	"github.com/FadhilAF/perbakin-be/internal/model"
)

func (usecase *scorerUsecaseImpl) Stage2try1Finished(req model.ByIdRequest) (bool, error) {
	status, err := usecase.Store.GetStage2try1Status(context.Background(), req.ID)
	if status != model.Stage246EndStatus {
		return false, err
	}

	return true, nil
}
