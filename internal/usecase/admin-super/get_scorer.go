package usecase

import (
	"context"
	"net/http"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (usecase *adminSuperUsecaseImpl) GetScorer(req model.GetUserById) model.WebServiceResponse {
	scorer, err := usecase.Store.GetScorer(context.Background(), req.ID)
	if err != nil {
		return util.ToWebServiceResponse("Id tidak terdaftar sebagai scorer", http.StatusUnauthorized, nil)
	}

	return util.ToWebServiceResponse("Data scorer ditemukan", http.StatusOK, gin.H{
		"scorer": model.Scorer{
			ID: scorer.ID,
			User: model.User{
				ID:       scorer.UserID,
				Username: scorer.Username,
				Name:     scorer.Name,
			},
		}})
}
