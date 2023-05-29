package delivery

import (
	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (handler *adminSuperHandler) CreateShooter(c *gin.Context) {
	scorer := c.MustGet("scorer").(model.OperatorRelation)

	var req model.CreateShooterBodyRequest

	if ok := util.BindFormAndValidate(c, &req); !ok {
		return
	}

	// Simpan upload file ke folder assets/images
	imagePath, ok := util.SaveFileFromForm(c, "image", "media/")
	if !ok {
		return
	}

	res := handler.Usecase.CreateShooter(model.CreateShooterRequest{
		ScorerID:  scorer.ID,
		ImagePath: imagePath,
		Body:      req,
	})

	c.JSON(res.Status, res)
}
