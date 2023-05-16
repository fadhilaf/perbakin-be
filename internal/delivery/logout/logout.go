package delivery

import (
	"net/http"

	"github.com/FadhilAF/perbakin-be/internal/util"

	"github.com/gin-gonic/gin"
)

func Logout(c *gin.Context) {
	uuid := util.GetUserIdFromContext(c)

	if !uuid.Valid {
		res := util.ToWebServiceResponse("Sudah Logout", http.StatusOK, nil)
		c.JSON(res.Status, res)
		return
	}

	util.RemoveAuthSession(c)
	res := util.ToWebServiceResponse("Logout berhasil", http.StatusOK, nil)
	util.SetAuthStatusCookie(c, "")
	c.JSON(res.Status, res)
}
