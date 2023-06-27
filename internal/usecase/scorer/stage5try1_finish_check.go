package usecase

import (
	"context"

	"github.com/FadhilAF/perbakin-be/internal/model"
)

func (usecase *scorerUsecaseImpl) Stage5try1Finished(req model.ByIdRequest) (bool, error) {
	status, err := usecase.Store.GetStage5try1Status(context.Background(), req.ID)
	if status != model.Stage5EndStatus {
		return false, err
	}

	return true, nil
}
