package usecase

import (
	"context"
	"net/http"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (usecase *adminSuperUsecaseImpl) GetShootersByExamId(req model.ByExamIdRequest) model.WebServiceResponse {
	shooters, err := usecase.Store.GetShooterByExamId(context.Background(), req.ExamID)
	if err != nil {
		return util.ToWebServiceResponse("Gagal mendapatkan data penembak", http.StatusInternalServerError, nil)
	}

	var shootersData []model.ShooterDisplayScorerData
	for _, shooter := range shooters {
		shootersData = append(shootersData, model.ShooterDisplayScorerData{
			ID:       shooter.ID,
			ScorerID: shooter.ScorerID,
			Scorer:   shooter.Scorer,
			Name:     shooter.Name,
			Province: shooter.Province,
			Club:     shooter.Club,
		})
	}

	return util.ToWebServiceResponse("Berhasil mendapatkan data penembak", http.StatusOK, gin.H{
		"shooters": shootersData,
	})
}
