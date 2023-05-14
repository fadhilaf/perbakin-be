package usecase

import (
	"context"
	"net/http"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (usecase *allUsecaseImpl) GetShootersByScorerId(req model.ByScorerIdRequest) model.WebServiceResponse {
	shooters, err := usecase.Store.GetShootersByScorerId(context.Background(), req.ScorerID)
	if err != nil {
		return util.ToWebServiceResponse("Gagal mendapatkan data penembak", http.StatusInternalServerError, nil)
	}

	var shootersData []model.ShooterDisplayData
	for _, shooter := range shooters {
		shootersData = append(shootersData, model.ShooterDisplayData{
			ID:       shooter.ID,
			Name:     shooter.Name,
			Province: shooter.Province,
			Club:     shooter.Club,
		})
	}

	return util.ToWebServiceResponse("Berhasil mendapatkan data penembak", http.StatusOK, gin.H{
		"shooters": shootersData,
	})
}
