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

	var shootersDisplayData []model.ShooterDisplayExamData
	for _, shooter := range shooters {
		shootersDisplayData = append(shootersDisplayData, model.ShooterDisplayExamData{
			Exam:      shooter.Exam,
			Name:      shooter.Name,
			ImagePath: shooter.ImagePath,
			Province:  shooter.Province,
			Club:      shooter.Club,
		})
	}

	return util.ToWebServiceResponse("Berhasil mendapatkan data penembak", http.StatusOK, gin.H{
		"shooters": shootersDisplayData,
	})
}
