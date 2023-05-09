package delivery

import (
	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (handler *superHandler) UpdateAdmin(c *gin.Context) {
	id, ok := util.GetIdParam(c)
	if !ok {
		return
	}

	var req model.UpdateUserBodyRequest

	if ok := util.BindJSONAndValidate(c, &req); !ok {
		return
	}

	res := handler.Usecase.UpdateAdmin(model.UpdateUserRequest{ID: id, Body: req})

	c.JSON(res.Status, res)
}
