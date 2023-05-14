package delivery

import (
	allUsecase "github.com/FadhilAF/perbakin-be/internal/usecase/all"
	scorerUsecase "github.com/FadhilAF/perbakin-be/internal/usecase/scorer"

	"github.com/gin-gonic/gin"
)

type ScorerDelivery interface {
	ScorerLogin(c *gin.Context)

	CheckScorerLogin(c *gin.Context)

	MustScorerMiddleware() gin.HandlerFunc

	GetShootersByScorerId(c *gin.Context)
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
