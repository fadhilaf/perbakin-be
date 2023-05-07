package delivery

import (
	"net/http"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"

	"github.com/gin-gonic/gin"
)

func (handler *adminHandler) AdminLogin(c *gin.Context) {
	var req model.LoginRequest

	if ok := util.BindJSONAndValidate(c, &req); !ok {
		return
	}

	res := handler.Usecase.AdminLogin(req)

	if res.Status == http.StatusOK {
		if admin, ok := res.Data["admin"].(model.Admin); ok {
			util.SaveUserToSession(c, admin.User.ID)
		}
	}

	c.JSON(res.Status, res)
}
