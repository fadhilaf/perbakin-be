package delivery

import (
	"net/http"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (handler *adminHandler) MustShooterMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		scorer := c.MustGet("scorer").(model.OperatorRelation)

		shooterId, ok := util.GetIdParam(c, "shooter_id")
		if !ok {
			c.Abort()
			return
		}

		shooter, err := handler.AllUsecase.GetShooterRelationById(model.ByIdRequest{ID: shooterId})
		if err != nil {
			res := util.ToWebServiceResponse("Penembak tidak ditemukan", http.StatusNotFound, nil)
			c.JSON(res.Status, res)
			c.Abort()
			return
		}

		if shooter.ScorerID != scorer.ID {
			res := util.ToWebServiceResponse("Tidak dapat mengakses penembak penguji lain", http.StatusUnauthorized, nil)
			c.JSON(res.Status, res)
			c.Abort()
			return
		}

		c.Set("shooter", shooter)
		c.Next()
	}
}
