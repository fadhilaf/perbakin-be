package delivery

import (
	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (handler *adminSuperHandler) UpdateShooterImage(c *gin.Context) {
	shooter := c.MustGet("shooter").(model.ShooterRelation)

	// Simpan upload file ke folder assets/images
	imagePath, ok := util.SaveFileFromForm(c, "image", "media/")
	if !ok {
		return
	}

	res := handler.Usecase.UpdateShooterImage(model.UpdateShooterImageRequest{
		ID:        shooter.ID,
		ImagePath: imagePath,
	})

	c.JSON(res.Status, res)
}
