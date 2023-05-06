package delivery

import (
	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (handler *superHandler) CheckSuperLogin(c *gin.Context) {
	userId := util.GetUserIdFromContext(c)

	if !userId.Valid {
		c.JSON(401, util.ToWebServiceResponse("User belum login", 401, nil))
		return
	}

	res := handler.Usecase.GetSuperByUserId(model.GetByUserIdRequest{UserID: userId})

	c.JSON(res.Status, res)
}
