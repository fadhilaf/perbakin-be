package delivery

import (
	scorerUsecase "github.com/FadhilAF/perbakin-be/internal/usecase/scorer"
	sessionUsecase "github.com/FadhilAF/perbakin-be/internal/usecase/session"

	"github.com/gin-gonic/gin"
)

type ScorerDelivery interface {
	ScorerLogin(c *gin.Context)
	CheckScorerLogin(c *gin.Context)
	MustScorerMiddleware() gin.HandlerFunc
}

var _ ScorerDelivery = &scorerHandler{}

func NewScorerDelivery(usecase scorerUsecase.ScorerUsecase, sessionUsecase sessionUsecase.SessionUsecase) ScorerDelivery {
	return &scorerHandler{
		Usecase:        usecase,
		SessionUsecase: sessionUsecase,
	}
}

type scorerHandler struct {
	Usecase        scorerUsecase.ScorerUsecase
	SessionUsecase sessionUsecase.SessionUsecase
}
