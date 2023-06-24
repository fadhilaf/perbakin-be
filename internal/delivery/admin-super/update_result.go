package delivery

import (
	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (handler *adminSuperHandler) UpdateResult(c *gin.Context) {
	result := c.MustGet("result").(model.ResultRelationAndStatus)

	var req model.UpdateResultBodyRequest

	if ok := util.BindJSONAndValidate(c, &req); !ok {
		return
	}

	res := handler.Usecase.UpdateResult(model.UpdateResultRequest{
		ID:   result.ID,
		Body: req,
	})

	c.JSON(res.Status, res)
}
