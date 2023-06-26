package usecase

import (
	"context"

	"github.com/FadhilAF/perbakin-be/internal/model"
)

func (usecase *adminSuperUsecaseImpl) FinishStage3(req model.ByIdRequest) error {
	return usecase.Store.FinishStage3(context.Background(), req.ID)
}
