package delivery

import (
	"net/http"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (handler *superHandler) CreateAdmin(c *gin.Context) {
	_, ok := c.MustGet("super").(model.Super)
	if !ok {
		res := util.ToWebServiceResponse("Error ketika parsing data super", http.StatusInternalServerError, nil)
		c.JSON(res.Status, res)
		return
	}

	var req model.CreateUserRequest

	if ok := util.BindJSONAndValidate(c, &req); !ok {
		return
	}

	res := handler.Usecase.CreateAdmin(req)

	c.JSON(res.Status, res)
}
