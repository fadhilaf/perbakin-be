package delivery

import (
	"net/http"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"

	"github.com/gin-gonic/gin"
)

func (handler *scorerHandler) ScorerLogin(c *gin.Context) {
	if _, exist := c.Get("user_id"); exist {
		res := util.ToWebServiceResponse("Sudah log in, log out terlebih dahulu sebelum log in kembali", http.StatusConflict, nil)
		c.JSON(res.Status, res)
		return
	}

	var req model.LoginRequest

	if ok := util.BindJSONAndValidate(c, &req); !ok {
		return
	}

	res := handler.Usecase.ScorerLogin(req)

	if res.Status == http.StatusOK {
		if scorer, ok := res.Data["scorer"].(model.Operator); ok {
			util.SaveUserToSession(c, scorer.User.ID)
			util.SetAuthStatusCookie(c, "scorer")
		}
	} else {
		util.SetAuthStatusCookie(c, "")
	}

	c.JSON(res.Status, res)
}
