package delivery

import (
	"net/http"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"

	"github.com/gin-gonic/gin"
)

func (handler *superHandler) SuperLogin(c *gin.Context) {
	if _, exist := c.Get("user_id"); exist {
		res := util.ToWebServiceResponse("Sudah log in, log out terlebih dahulu sebelum log in kembali", http.StatusConflict, nil)
		c.JSON(res.Status, res)
		return
	}

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
