package delivery

import (
	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
)

func (handler *scorerHandler) UpdateStage2Finish(c *gin.Context) {
	stage2 := c.MustGet("stage2").(model.Stage123456RelationAndStatus)
	try := c.MustGet("try").(string)

	var req model.UpdateStageFinishBodyRequest

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

	res := handler.Usecase.UpdateStage2Finish(model.UpdateStage123456FinishRequest{
		ID:          stage2.ID,
		Try:         try,
		Success:     *req.Success,
		ScorerSign:  scorerSignText,
		ShooterSign: shooterSignText,
	})

	c.JSON(res.Status, res)
}
