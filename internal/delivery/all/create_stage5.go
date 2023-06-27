package delivery

import (
	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/gin-gonic/gin"
)

func (handler *allHandler) CreateStage5(c *gin.Context) {
	result := c.MustGet("result").(model.ResultRelationAndStatus)

	res := handler.Usecase.CreateStage5(model.ByResultIdRequest{
		ResultID: result.ID,
	})

	c.JSON(res.Status, res)
}
