package delivery

import (
	"net/http"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (handler *superHandler) CheckSuperLogin(c *gin.Context) {
	super, ok := c.MustGet("super").(model.Super)
	if !ok {
		res := util.ToWebServiceResponse("Error ketika parsing data super", http.StatusInternalServerError, nil)
		c.JSON(res.Status, res)
		return
	}

	res := util.ToWebServiceResponse("User adalah super admin", http.StatusOK, gin.H{"super": super})
	c.JSON(res.Status, res)
}
