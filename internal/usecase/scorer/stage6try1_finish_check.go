package usecase

import (
	"context"

	"github.com/FadhilAF/perbakin-be/internal/model"
)

func (usecase *scorerUsecaseImpl) Stage6try1Finished(req model.ByIdRequest) (bool, error) {
	status, err := usecase.Store.GetStage6try1Status(context.Background(), req.ID)
	if status != model.Stage246EndStatus {
		return false, err
	}

	return true, nil
}
