package usecase

import (
	"context"

	"github.com/FadhilAF/perbakin-be/internal/model"
)

func (usecase *scorerUsecaseImpl) Stage1try1Finished(req model.ByIdRequest) (bool, error) {
	status, err := usecase.Store.GetStage1try1Status(context.Background(), req.ID)
	if status != "6" {
		return false, err
	}

	return true, nil
}
