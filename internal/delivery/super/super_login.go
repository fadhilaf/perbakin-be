package delivery

import (
	"net/http"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"

	"github.com/gin-gonic/gin"
)

func (handler *superHandler) SuperLogin(c *gin.Context) {
	var req model.LoginRequest

	if ok := util.BindJSONAndValidate(c, &req); !ok {
		return
	}

	res := handler.Usecase.SuperLogin(req)

	if res.Status == http.StatusOK {
		if super, ok := res.Data["super"].(model.Super); ok {
			util.SaveUserToSession(c, super.User.ID)
		}
	}

	c.JSON(res.Status, res)
}
