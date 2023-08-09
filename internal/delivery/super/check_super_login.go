package delivery

import (
	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/gin-gonic/gin"
)

func (handler *superHandler) CheckSuperLogin(c *gin.Context) {
	superRelation := c.MustGet("super").(model.SuperRelation)

	res := handler.Usecase.GetSuperById(model.ByIdRequest{ID: superRelation.ID})

	c.JSON(res.Status, res)
}
