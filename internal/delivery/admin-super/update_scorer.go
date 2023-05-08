package delivery

import (
	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (handler *adminSuperHandler) UpdateScorer(c *gin.Context) {
	id, ok := util.GetIdParam(c)
	if !ok {
		return
	}

	var req model.UpdateUserDataRequest

	if ok := util.BindJSONAndValidate(c, &req); !ok {
		return
	}

	res := handler.Usecase.UpdateScorer(model.UpdateUserRequest{ID: id, Data: req})

	c.JSON(res.Status, res)
}