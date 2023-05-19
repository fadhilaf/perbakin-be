package delivery

import (
	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/gin-gonic/gin"
)

func (handler *scorerHandler) CreateStage0(c *gin.Context) {
	result := c.MustGet("result").(model.ResultRelation)

	res := handler.AllUsecase.CreateStage0(model.ByResultIdRequest{
		ResultID: result.ID,
	})

	c.JSON(res.Status, res)
}
