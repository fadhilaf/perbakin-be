package delivery

import (
	scorerUsecase "github.com/FadhilAF/perbakin-be/internal/usecase/scorer"

	"github.com/gin-gonic/gin"
)

type ScorerDelivery interface {
	ScorerLogin(c *gin.Context)

	MustScorerMiddleware() gin.HandlerFunc
	CheckScorerLogin(c *gin.Context)

	MustStage0ModifyMiddleware() gin.HandlerFunc
	UpdateStage0Series(c *gin.Context)
	UpdateStage0NextSeries(c *gin.Context)
	UpdateStage0Checkmarks(c *gin.Context)
	UpdateStage0Finish(c *gin.Context)

	MustStage1ModifyMiddleware() gin.HandlerFunc
	MustStage1try1FinishedMiddleware() gin.HandlerFunc
	UpdateStage1No(c *gin.Context)
	UpdateStage1NextNo(c *gin.Context)
	UpdateStage1Checkmarks(c *gin.Context)
	UpdateStage1Finish(c *gin.Context)

	MustStage2ModifyMiddleware() gin.HandlerFunc
	MustStage2try1FinishedMiddleware() gin.HandlerFunc
	UpdateStage2No(c *gin.Context)
	UpdateStage2NextNo(c *gin.Context)
	UpdateStage2Checkmarks(c *gin.Context)
	UpdateStage2Finish(c *gin.Context)

	MustStage3ModifyMiddleware() gin.HandlerFunc
	MustStage3try1FinishedMiddleware() gin.HandlerFunc
	UpdateStage3No(c *gin.Context)
	UpdateStage3NextNo(c *gin.Context)
	UpdateStage3Checkmarks(c *gin.Context)
	UpdateStage3Finish(c *gin.Context)

	MustStage4ModifyMiddleware() gin.HandlerFunc
	MustStage4try1FinishedMiddleware() gin.HandlerFunc
	UpdateStage4No(c *gin.Context)
	UpdateStage4NextNo(c *gin.Context)
	UpdateStage4Checkmarks(c *gin.Context)
	UpdateStage4Finish(c *gin.Context)
}

var _ ScorerDelivery = &scorerHandler{}

func NewScorerDelivery(usecase scorerUsecase.ScorerUsecase) ScorerDelivery {
	return &scorerHandler{
		Usecase: usecase,
	}
}

type scorerHandler struct {
	Usecase scorerUsecase.ScorerUsecase
}
