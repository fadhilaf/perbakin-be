package usecase

import (
	"context"
	"net/http"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (usecase *adminSuperUsecaseImpl) GetScorerById(req model.ByIdRequest) model.WebServiceResponse {
	scorer, _ := usecase.Store.GetScorerById(context.Background(), req.ID)

	return util.ToWebServiceResponse("Data scorer ditemukan", http.StatusOK, gin.H{
		"scorer": model.Operator{
			ID:     scorer.ID,
			ExamID: scorer.ExamID,
			User: model.User{
				ID:        scorer.UserID,
				Username:  scorer.Username,
				Name:      scorer.Name,
				CreatedAt: scorer.CreatedAt.Time,
				UpdatedAt: scorer.UpdatedAt.Time,
			},
		}})
}
