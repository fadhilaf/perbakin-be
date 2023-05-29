package delivery

import (
	allUsecase "github.com/FadhilAF/perbakin-be/internal/usecase/all"
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
}

var _ ScorerDelivery = &scorerHandler{}

func NewScorerDelivery(usecase scorerUsecase.ScorerUsecase, allUsecase allUsecase.AllUsecase) ScorerDelivery {
	return &scorerHandler{
		Usecase:    usecase,
		AllUsecase: allUsecase,
	}
}

type scorerHandler struct {
	Usecase    scorerUsecase.ScorerUsecase
	AllUsecase allUsecase.AllUsecase
}
