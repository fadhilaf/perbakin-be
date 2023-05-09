package delivery

import (
	"net/http"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (handler *adminHandler) CheckAdminLogin(c *gin.Context) {
	admin := c.MustGet("admin").(model.OperatorRelation)

	res := util.ToWebServiceResponse("User adalah admin", http.StatusOK, gin.H{"admin": admin})
	c.JSON(res.Status, res)
}
