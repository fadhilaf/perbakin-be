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
	CreateStage1try2(c *gin.Context)
	UpdateStage1No(c *gin.Context)
	UpdateStage1NextNo(c *gin.Context)
	UpdateStage1Checkmarks(c *gin.Context)
	UpdateStage1Finish(c *gin.Context)
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
