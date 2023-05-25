package delivery

import (
	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/gin-gonic/gin"
)

func (handler *scorerHandler) GetResultById(c *gin.Context) {
	result := c.MustGet("result").(model.ResultRelation)
	res := handler.AllUsecase.GetResultById(model.ByIdRequest{
		ID: result.ID,
	})

	c.JSON(res.Status, res)
}
