package delivery

import (
	"net/http"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (handler *adminSuperHandler) MustResultMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		shooter := c.MustGet("shooter").(model.ShooterRelation)

		result, err := handler.AllUsecase.GetResultRelationByShooterId(model.ByShooterIdRequest{
			ShooterID: shooter.ID,
		})
		if err != nil {
			res := util.ToWebServiceResponse("Hasil ujian belum ada", http.StatusNotFound, nil)
			c.JSON(res.Status, res)
			c.Abort()
			return
		}

		if shooter.ID != result.ShooterID {
			res := util.ToWebServiceResponse("Tidak dapat mengakses hasil ujian penembak lain", http.StatusUnauthorized, nil)
			c.JSON(res.Status, res)
			c.Abort()
			return
		}

		c.Set("result", result)
		c.Next()
	}
}
