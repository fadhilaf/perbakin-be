package delivery

import (
	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/gin-gonic/gin"
)

func (handler *adminSuperHandler) DeleteResult(c *gin.Context) {
	result := c.MustGet("result").(model.ResultRelation)

	res := handler.Usecase.DeleteResult(model.ByIdRequest{
		ID: result.ID,
	})

	c.JSON(res.Status, res)
}
