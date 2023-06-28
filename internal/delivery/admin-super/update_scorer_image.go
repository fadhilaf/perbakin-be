package delivery

import (
	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (handler *adminSuperHandler) UpdateScorerImage(c *gin.Context) {
	scorer := c.MustGet("scorer").(model.OperatorRelation)

	// Simpan upload file ke folder assets/images
	imagePath, ok := util.MustSaveFileFromForm(c, "image", "media/")
	if !ok {
		return
	}

	res := handler.Usecase.UpdateScorerImage(model.UpdateImageRequest{
		ID:        scorer.ID,
		ImagePath: imagePath,
	})

	c.JSON(res.Status, res)
}
