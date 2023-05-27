package delivery

import (
	// "strconv"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
)

func (handler *scorerHandler) UpdateStage0Finish(c *gin.Context) {
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

	// success, _ := strconv.ParseBool(req.Success)
	success := req.Success

	res := handler.Usecase.UpdateStage0Finish(model.UpdateStage0FinishRequest{
		ID:          stage0.ID,
		Success:     success,
		ScorerSign:  scorerSignText,
		ShooterSign: shooterSignText,
	})

	c.JSON(res.Status, res)
}
