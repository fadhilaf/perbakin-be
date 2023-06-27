package usecase

import (
	"context"

	repositoryModel "github.com/FadhilAF/perbakin-be/internal/repository/postgres/sqlc"

	"github.com/FadhilAF/perbakin-be/internal/model"
)

func (usecase *adminSuperUsecaseImpl) UpdateResultStage(req model.UpdateResultStageRequest) bool {
	err := usecase.Store.UpdateResultStage(context.Background(), repositoryModel.UpdateResultStageParams{
		ID:    req.ID,
		Stage: repositoryModel.Stages(req.Stage),
	})
	if err != nil {
		return true
	}

	return false
}
