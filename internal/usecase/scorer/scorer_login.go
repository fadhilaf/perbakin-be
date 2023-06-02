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
		return util.ToWebServiceResponse("Username yang dimasukkan tidak ditemukan", http.StatusNotFound, nil)
	}

	if !scorer.Active {
		return util.ToWebServiceResponse("Ujian tidak aktif", http.StatusForbidden, nil)
	}

	if err := util.ComparePassword(req.Password, scorer.Password); err != nil {
		return util.ToWebServiceResponse("Password yang dimasukkan salah", http.StatusUnauthorized, nil)
	}

	return util.ToWebServiceResponse("Login berhasil", http.StatusOK, gin.H{
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
