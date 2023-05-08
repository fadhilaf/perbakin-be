package delivery

import (
	"net/http"

	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (handler *superHandler) CheckSuperLogin(c *gin.Context) {
	super, ok := util.MustGetSuper(c)
	if !ok {
		return
	}

	res := util.ToWebServiceResponse("User adalah super admin", http.StatusOK, gin.H{"super": super})
	c.JSON(res.Status, res)
}
