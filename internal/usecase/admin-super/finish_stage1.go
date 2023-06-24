package usecase

import (
	"context"

	"github.com/FadhilAF/perbakin-be/internal/model"
)

func (usecase *adminSuperUsecaseImpl) FinishStage1(req model.ByIdRequest) error {
	return usecase.Store.FinishStage1(context.Background(), req.ID)
}
