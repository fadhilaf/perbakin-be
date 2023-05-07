package delivery

import (
	"net/http"

	"github.com/FadhilAF/perbakin-be/internal/util"

	"github.com/gin-gonic/gin"
)

func Logout(c *gin.Context) {
	uuid := util.GetUserIdFromContext(c)

	res := util.ToWebServiceResponse("Sudah Logout", http.StatusOK, nil)

	if uuid.Valid {
		util.RemoveAuthSession(c)
		res = util.ToWebServiceResponse("Logout berhasil", http.StatusOK, nil)
	}

	c.JSON(res.Status, res)
}
