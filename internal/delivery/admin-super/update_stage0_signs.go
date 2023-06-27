package delivery

import (
	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
)

func (handler *adminSuperHandler) UpdateStage0Signs(c *gin.Context) {
	stage0 := c.MustGet("stage0").(model.Stage0Relation)

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

	res := handler.Usecase.UpdateStage0Signs(model.UpdateStageSignsRequest{
		ID:          stage0.ID,
		ShooterSign: shooterSignText,
		ScorerSign:  scorerSignText,
	})

	c.JSON(res.Status, res)
}
