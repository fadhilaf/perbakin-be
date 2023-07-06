package usecase

import (
	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/repository"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
)

type UtilUsecase interface {
	CheckCheckmarkAmountAndSeriesMinimumScore(c *gin.Context, ID pgtype.UUID, checkmarks []bool, stageType model.StageList) bool
}

var _ UtilUsecase = &utilUsecaseImpl{}

func NewUtilUsecase(store repository.Store) UtilUsecase {
	return &utilUsecaseImpl{
		Store: store,
	}
}

type utilUsecaseImpl struct {
	repository.Store
}
