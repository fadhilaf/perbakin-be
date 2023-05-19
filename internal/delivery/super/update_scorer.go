package delivery

import (
	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (handler *superHandler) UpdateScorer(c *gin.Context) {
	scorer := c.MustGet("scorer").(model.OperatorRelation)

	var req model.OperatorBodyRequest

	if ok := util.BindJSONAndValidate(c, &req); !ok {
		return
	}

	res := handler.AdminSuperUsecase.UpdateScorer(model.UpdateOperatorRequest{ID: scorer.ID, Body: req})

	c.JSON(res.Status, res)
}
