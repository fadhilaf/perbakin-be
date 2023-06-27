package usecase

import (
	"context"

	"github.com/FadhilAF/perbakin-be/internal/model"
)

func (usecase *adminSuperUsecaseImpl) FinishStage6(req model.ByIdRequest) error {
	return usecase.Store.FinishStage6(context.Background(), req.ID)
}
