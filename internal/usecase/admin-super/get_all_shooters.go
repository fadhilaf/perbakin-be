package usecase

import (
	"context"
	"net/http"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (usecase *adminSuperUsecaseImpl) GetAllShooters() model.WebServiceResponse {
	shooters, err := usecase.Store.GetAllShooters(context.Background())
	if err != nil {
		return util.ToWebServiceResponse("Gagal mendapatkan data penembak"+err.Error(), http.StatusInternalServerError, nil)
	}

	var shootersDisplayData []model.ShooterDisplayData
	for _, shooter := range shooters {
		shootersDisplayData = append(shootersDisplayData, model.ShooterDisplayData{
			ID:        shooter.ID,
			Exam:      shooter.Exam,
			Name:      shooter.Name,
			Province:  shooter.Province,
			Club:      shooter.Club,
			CreatedAt: shooter.CreatedAt.Time,
			UpdatedAt: shooter.UpdatedAt.Time,
		})
	}

	return util.ToWebServiceResponse("Berhasil mendapatkan data penembak", http.StatusOK, gin.H{
		"shooters": shootersDisplayData,
	})
}
