package usecase

import (
	"context"
	"net/http"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (usecase *scorerUsecaseImpl) ScorerLogin(req model.LoginRequest) model.WebServiceResponse {
	scorer, err := usecase.Store.GetScorerByUsername(context.Background(), req.Username)

	if err != nil {
		return util.ToWebServiceResponse("Username yang dimasukkan salah", http.StatusNotFound, nil)
	}

	if err := util.ComparePassword(req.Password, scorer.Password); err != nil {
		return util.ToWebServiceResponse("Password yang dimasukkan salah", http.StatusUnauthorized, nil)
	}

	return util.ToWebServiceResponse("Login berhasil", http.StatusOK, gin.H{
		"scorer": model.Scorer{
			ID: scorer.ID,
			User: model.User{
				ID:       scorer.UserID,
				Username: scorer.Username,
				Name:     scorer.Name,
			},
		}})
}
