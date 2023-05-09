package usecase

import (
	"context"
	"net/http"

	"github.com/FadhilAF/perbakin-be/internal/model"
	repositoryModel "github.com/FadhilAF/perbakin-be/internal/repository/postgres/sqlc"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (usecase *adminSuperUsecaseImpl) CreateScorer(req model.CreateOperatorRequest) model.WebServiceResponse {
	if _, err := usecase.Store.GetUserByUsername(context.Background(), req.Body.Username); err == nil {
		return util.ToWebServiceResponse("Username sudah digunakan", http.StatusConflict, nil)
	}

	passwordHash, err := util.HashPassword(req.Body.Password)
	if err != nil {
		return util.ToWebServiceResponse("Gagal proses hash password: "+err.Error(), http.StatusInternalServerError, nil)
	}

	scorer, err := usecase.Store.CreateScorer(context.Background(), repositoryModel.CreateScorerParams{
		ExamID:   req.ExamID,
		Username: req.Body.Username,
		Password: passwordHash,
		Name:     req.Body.Name,
	})
	if err != nil {
		return util.ToWebServiceResponse("Gagal membuat penguji: "+err.Error(), http.StatusInternalServerError, nil)
	}

	return util.ToWebServiceResponse("Berhasil membuat penguji", http.StatusCreated, gin.H{
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
