package delivery

import (
	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
)

func (handler *adminSuperHandler) CreateScorer(c *gin.Context) {
	exam := c.MustGet("exam").(model.ExamRelation)

	var req model.CreateScorerBodyRequest

	if ok := util.BindFormAndValidate(c, &req); !ok {
		return
	}

	var imagePathText pgtype.Text

	// Simpan upload file ke folder assets/images
	if imagePath, ok := util.SaveFileFromForm(c, "image", "media/"); ok {
		imagePathText.Scan(imagePath)
	}

	res := handler.Usecase.CreateScorer(model.CreateScorerRequest{
		ExamID:    exam.ID,
		ImagePath: imagePathText,
		Body:      req,
	})

	c.JSON(res.Status, res)
}
