package usecase

import (
	"context"
	"net/http"

	respositoryModel "github.com/FadhilAF/perbakin-be/internal/repository/postgres/sqlc"
	"github.com/gin-gonic/gin"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
)

func (usecase *adminSuperUsecaseImpl) UpdateScorer(req model.UpdateOperatorRequest) model.WebServiceResponse {
	passwordHash, err := util.HashPassword(req.Body.Password)
	if err != nil {
		return util.ToWebServiceResponse("Gagal proses hash password: "+err.Error(), http.StatusInternalServerError, nil)
	}

	newScorer, err := usecase.Store.UpdateScorer(context.Background(), respositoryModel.UpdateScorerParams{
		ID:       req.ID,
		Username: req.Body.Username,
		Password: passwordHash,
		Name:     req.Body.Name,
	})
	if err != nil {
		return util.ToWebServiceResponse("Gagal mengubah scorer: "+err.Error(), http.StatusInternalServerError, nil)
	}

	return util.ToWebServiceResponse("Berhasil mengubah scorer", http.StatusOK, gin.H{
		"scorer": model.Operator{
			ID:     newScorer.ID,
			ExamID: newScorer.ExamID,
			User: model.User{
				ID:        newScorer.UserID,
				Username:  newScorer.Username,
				Name:      newScorer.Name,
				CreatedAt: newScorer.CreatedAt.Time,
				UpdatedAt: newScorer.UpdatedAt.Time,
			},
		}})
}
