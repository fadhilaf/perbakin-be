package delivery

import (
	"net/http"

	"github.com/FadhilAF/perbakin-be/internal/util"

	"github.com/gin-gonic/gin"
)

func (handler *authHandler) Logout(ctx *gin.Context) {
	util.RemoveAuthSession(ctx)

	res := util.ToWebServiceResponse("Logout berhasil", http.StatusOK, nil)

	ctx.JSON(res.Status, res)
}
