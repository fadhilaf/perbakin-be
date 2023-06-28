package usecase

import (
	"context"
	"net/http"

	repositoryModel "github.com/FadhilAF/perbakin-be/internal/repository/postgres/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
)

func (usecase *adminSuperUsecaseImpl) UpdateScorer(req model.UpdateOperatorRequest) model.WebServiceResponse {
	var query repositoryModel.UpdateScorerParams

	if req.Password.Valid {
		passwordHash, err := util.HashPassword(req.Password.String)
		if err != nil {
			return util.ToWebServiceResponse("Gagal proses hash password: "+err.Error(), http.StatusInternalServerError, nil)
		}

		var passwordHashText pgtype.Text
		passwordHashText.Scan(passwordHash)

		query = repositoryModel.UpdateScorerParams{
			ID:       req.ID,
			Username: req.Username,
			Password: passwordHashText,
			Name:     req.Name,
		}
	} else {
		query = repositoryModel.UpdateScorerParams{
			ID:       req.ID,
			Username: req.Username,
			Name:     req.Name,
		}
	}

	newScorer, err := usecase.Store.UpdateScorer(context.Background(), query)
	if err != nil {
		return util.ToWebServiceResponse("Gagal mengubah scorer: "+err.Error(), http.StatusInternalServerError, nil)
	}

	return util.ToWebServiceResponse("Berhasil mengubah scorer", http.StatusOK, gin.H{
		"scorer": model.Scorer{
			ID:        newScorer.ID,
			ExamID:    newScorer.ExamID,
			ImagePath: newScorer.ImagePath,
			User: model.User{
				ID:        newScorer.UserID,
				Username:  newScorer.Username,
				Name:      newScorer.Name,
				CreatedAt: newScorer.CreatedAt.Time,
				UpdatedAt: newScorer.UpdatedAt.Time,
			},
		}})
}
