package delivery

import (
	"net/http"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (handler *superHandler) MustAdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		exam := c.MustGet("exam").(model.ExamRelation)

		adminId, ok := util.GetIdParam(c, "admin_id")
		if !ok {
			c.Abort()
			return
		}

		admin, err := handler.Usecase.GetAdminRelationById(model.ByIdRequest{ID: adminId})
		if err != nil {
			res := util.ToWebServiceResponse("Admin tidak ditemukan", http.StatusNotFound, nil)
			c.JSON(res.Status, res)
			c.Abort()
			return
		}

		if admin.ExamID != exam.ID {
			res := util.ToWebServiceResponse("Tidak dapat mengakses admin ujian lain", http.StatusUnauthorized, nil)
			c.JSON(res.Status, res)
			c.Abort()
			return
		}

		c.Set("admin", admin)
		c.Next()
	}
}
