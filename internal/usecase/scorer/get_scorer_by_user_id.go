package usecase

import (
	"context"
	"net/http"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (usecase *scorerUsecaseImpl) GetScorerByUserId(req model.GetByUserIdRequest) model.WebServiceResponse {
	scorer, err := usecase.Store.GetScorerByUserId(context.Background(), req.UserID)
	if err != nil {
		return util.ToWebServiceResponse("User tidak terdaftar sebagai penguji", http.StatusUnauthorized, nil)
	}

	return util.ToWebServiceResponse("User adalah penguji", http.StatusOK, gin.H{
		"scorer": model.Scorer{
			ID: scorer.ID,
			User: model.User{
				ID:       scorer.UserID,
				Username: scorer.Username,
				Name:     scorer.Name,
			},
		}})
}
