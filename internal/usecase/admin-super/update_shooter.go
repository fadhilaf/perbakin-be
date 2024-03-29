package usecase

import (
	"context"
	"net/http"

	repositoryModel "github.com/FadhilAF/perbakin-be/internal/repository/postgres/sqlc"
	"github.com/gin-gonic/gin"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
)

func (usecase *adminSuperUsecaseImpl) UpdateShooter(req model.UpdateShooterRequest) model.WebServiceResponse {
	scorer, err := usecase.Store.GetScorerById(context.Background(), req.Body.ScorerID)
	if err != nil {
		return util.ToWebServiceResponse("Tidak ditemukan penguji pengganti dengan ID yang diberikan", http.StatusNotFound, nil)
	}

	if req.ExamID != scorer.ExamID {
		return util.ToWebServiceResponse("Tidak dapat mengubah penguji penembak dengan penguji ujian lain", http.StatusUnauthorized, nil)
	}

	newShooter, err := usecase.Store.UpdateShooter(context.Background(), repositoryModel.UpdateShooterParams{
		ID:       req.ID,
		ScorerID: req.Body.ScorerID,
		Name:     req.Body.Name,
		Province: req.Body.Province,
		Club:     req.Body.Club,
	})
	if err != nil {
		return util.ToWebServiceResponse("Gagal mengubah penembak: "+err.Error(), http.StatusInternalServerError, nil)
	}

	return util.ToWebServiceResponse("Berhasil mengubah penembak", http.StatusOK, gin.H{
		"shooter": model.Shooter{
			ID:        newShooter.ID,
			ScorerID:  newShooter.ScorerID,
			Name:      newShooter.Name,
			Province:  newShooter.Province,
			Club:      newShooter.Club,
			CreatedAt: newShooter.CreatedAt.Time,
			UpdatedAt: newShooter.UpdatedAt.Time,
		}})
}
