package usecase

import (
	"context"

	"github.com/FadhilAF/perbakin-be/internal/model"
)

func (usecase *adminSuperUsecaseImpl) FinishStage5(req model.ByIdRequest) error {
	return usecase.Store.FinishStage5(context.Background(), req.ID)
}
