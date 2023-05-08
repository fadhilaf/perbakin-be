package delivery

import (
	usecase "github.com/FadhilAF/perbakin-be/internal/usecase/scorer"
	"github.com/gin-gonic/gin"
)

type ScorerDelivery interface {
	ScorerLogin(c *gin.Context)
	CheckScorerLogin(c *gin.Context)
	MustScorerMiddleware() gin.HandlerFunc
}

var _ ScorerDelivery = &scorerHandler{}

func NewScorerDelivery(usecase usecase.ScorerUsecase) ScorerDelivery {
	return &scorerHandler{
		Usecase: usecase,
	}
}

type scorerHandler struct {
	Usecase usecase.ScorerUsecase
}
