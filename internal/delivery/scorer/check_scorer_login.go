package delivery

import (
	"net/http"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (handler *scorerHandler) CheckScorerLogin(c *gin.Context) {
	scorer := c.MustGet("scorer").(model.OperatorRelation)

	res := util.ToWebServiceResponse("User adalah scorer", http.StatusOK, gin.H{"scorer": scorer})
	c.JSON(res.Status, res)
}
