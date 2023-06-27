package delivery

import (
	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
)

func (handler *adminSuperHandler) UpdateStage1Signs(c *gin.Context) {
	stage1 := c.MustGet("stage1").(model.Stage123456RelationAndStatus)

	// Simpan upload file ke folder assets/images
	shooterSign, ok := util.MustSaveFileFromForm(c, "shooter_sign", "media/")
	if !ok {
		return
	}

	scorerSign, ok := util.MustSaveFileFromForm(c, "scorer_sign", "media/")
	if !ok {
		return
	}

	var shooterSignText pgtype.Text
	var scorerSignText pgtype.Text

	shooterSignText.Scan(shooterSign)
	scorerSignText.Scan(scorerSign)

	res := handler.Usecase.UpdateStage1Signs(model.UpdateStageSignsRequest{
		ID:          stage1.ID,
		ShooterSign: shooterSignText,
		ScorerSign:  scorerSignText,
	})

	c.JSON(res.Status, res)
}
