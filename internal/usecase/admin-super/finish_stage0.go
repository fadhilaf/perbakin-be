package usecase

import (
	"context"

	"github.com/FadhilAF/perbakin-be/internal/model"
)

func (usecase *adminSuperUsecaseImpl) FinishStage0(req model.ByIdRequest) error {
	return usecase.Store.FinishStage0(context.Background(), req.ID)
}
