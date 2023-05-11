package usecase

import (
	"context"
	"net/http"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (usecase *scorerUsecaseImpl) GetScorerByUserId(req model.UserByUserIdRequest) model.WebServiceResponse {
	scorer, err := usecase.Store.GetScorerByUserId(context.Background(), req.UserID)
	if err != nil {
		return util.ToWebServiceResponse("User tidak terdaftar sebagai scorer", http.StatusUnauthorized, nil)
	}

	return util.ToWebServiceResponse("Data scorer ditemukan", http.StatusOK, gin.H{
		"scorer": model.Operator{
			ID:     scorer.ID,
			ExamID: scorer.ExamID,
			User: model.User{
				ID:       scorer.UserID,
				Username: scorer.Username,
				Name:     scorer.Name,
			},
		}})
}
