package delivery

import (
	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/gin-gonic/gin"
)

func (handler *allHandler) GetResultById(c *gin.Context) {
	result := c.MustGet("result").(model.ResultRelationAndStatus)
	res := handler.Usecase.GetResultById(model.ByIdRequest{
		ID: result.ID,
	})

	c.JSON(res.Status, res)
}
