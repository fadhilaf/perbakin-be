package usecase

import (
	"context"

	"github.com/FadhilAF/perbakin-be/internal/model"
)

func (usecase *adminSuperUsecaseImpl) FinishStage2(req model.ByIdRequest) error {
	return usecase.Store.FinishStage2(context.Background(), req.ID)
}
