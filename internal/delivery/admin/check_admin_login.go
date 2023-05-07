package delivery

import (
	"net/http"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (handler *adminHandler) CheckAdminLogin(c *gin.Context) {
	admin, ok := c.MustGet("admin").(model.Admin)
	if !ok {
		res := util.ToWebServiceResponse("Error ketika parsing data admin", http.StatusInternalServerError, nil)
		c.JSON(res.Status, res)
		return
	}

	res := util.ToWebServiceResponse("User adalah admin", http.StatusOK, gin.H{"admin": admin})
	c.JSON(res.Status, res)
}
