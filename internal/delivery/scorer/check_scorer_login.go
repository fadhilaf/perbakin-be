package delivery

import (
	"net/http"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (handler *scorerHandler) CheckScorerLogin(c *gin.Context) {
	scorer, ok := c.MustGet("scorer").(model.Scorer)
	if !ok {
		res := util.ToWebServiceResponse("Error ketika parsing data scorer", http.StatusInternalServerError, nil)
		c.JSON(res.Status, res)
		return
	}

	res := util.ToWebServiceResponse("User adalah scorer", http.StatusOK, gin.H{"scorer": scorer})
	c.JSON(res.Status, res)
}
