package delivery

import (
	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
)

func (handler *scorerHandler) UpdateStage0Finish(c *gin.Context) {
	result := c.MustGet("result").(model.ResultRelation)
	stage0 := c.MustGet("stage0").(model.Stage0Relation)

	var req model.UpdateStage0FinishBodyRequest

	if ok := util.BindFormAndValidate(c, &req); !ok {
		return
	}

	// Simpan upload file ke folder assets/images
	shooterSign, ok := util.SaveFileFromForm(c, "shooter_sign", "media/")
	if !ok {
		return
	}

	// Simpan upload file ke folder assets/images
	scorerSign, ok := util.SaveFileFromForm(c, "scorer_sign", "media/")
	if !ok {
		return
	}

	var shooterSignText pgtype.Text
	var scorerSignText pgtype.Text

	shooterSignText.Scan(shooterSign)
	scorerSignText.Scan(scorerSign)

	res := handler.Usecase.UpdateStage0Finish(model.UpdateStage0FinishRequest{
		ID:          stage0.ID,
		ResultID:    result.ID,
		ScorerSign:  scorerSignText,
		ShooterSign: shooterSignText,
		Body:        req,
	})

	c.JSON(res.Status, res)
}
