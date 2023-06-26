package usecase

import (
	"context"

	"github.com/FadhilAF/perbakin-be/internal/model"
)

func (usecase *adminSuperUsecaseImpl) FinishStage4(req model.ByIdRequest) error {
	return usecase.Store.FinishStage4(context.Background(), req.ID)
}
