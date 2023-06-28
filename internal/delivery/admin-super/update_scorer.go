package delivery

import (
	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
)

func (handler *adminSuperHandler) UpdateScorer(c *gin.Context) {
	scorer := c.MustGet("scorer").(model.OperatorRelation)

	var req model.OperatorBodyRequest

	if ok := util.BindJSONAndValidate(c, &req); !ok {
		return
	}

	var passwordText pgtype.Text
	if req.Password != "" {
		passwordText.Scan(req.Password)
	}

	res := handler.Usecase.UpdateScorer(model.UpdateOperatorRequest{ID: scorer.ID, UserID: scorer.UserID, Username: req.Username, Password: passwordText, Name: req.Name})

	c.JSON(res.Status, res)
}
