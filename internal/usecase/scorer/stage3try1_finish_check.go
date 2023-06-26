package usecase

import (
	"context"

	"github.com/FadhilAF/perbakin-be/internal/model"
)

func (usecase *scorerUsecaseImpl) Stage3try1Finished(req model.ByIdRequest) (bool, error) {
	status, err := usecase.Store.GetStage3try1Status(context.Background(), req.ID)
	if status != model.Stage13EndStatus {
		return false, err
	}

	return true, nil
}
