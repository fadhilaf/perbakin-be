package delivery

import (
	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (handler *adminSuperHandler) GetScorerById(c *gin.Context) {
	id, ok := util.GetIdParam(c)
	if !ok {
		return
	}

	res := handler.Usecase.GetScorerById(model.ByIdRequest{ID: id})

	c.JSON(res.Status, res)
}
