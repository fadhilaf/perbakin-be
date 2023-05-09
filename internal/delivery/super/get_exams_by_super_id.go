package delivery

import (
	"github.com/FadhilAF/perbakin-be/internal/model"

	"github.com/gin-gonic/gin"
)

func (handler *superHandler) GetExamsBySuperId(c *gin.Context) {
	super := c.MustGet("super").(model.SuperRelation)

	res := handler.Usecase.GetExamsByUserId(model.GetExamsBySuperIdRequest{
		SuperID: super.ID,
	})

	c.JSON(res.Status, res.Data)
}
