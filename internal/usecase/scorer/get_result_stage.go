package usecase

import (
	"context"

	"github.com/FadhilAF/perbakin-be/internal/model"
)

func (usecase *scorerUsecaseImpl) GetResultStageById(req model.ByIdRequest) (string, error) {
	stage, err := usecase.Store.GetResultStageById(context.Background(), req.ID)

	return string(stage.Stages), err
}
