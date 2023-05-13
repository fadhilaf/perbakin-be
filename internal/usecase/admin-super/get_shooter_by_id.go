package usecase

import (
	"context"
	"net/http"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (usecase *adminSuperUsecaseImpl) GetShooterById(req model.ShooterByIdRequest) model.WebServiceResponse {
	shooter, err := usecase.Store.GetShooterById(context.Background(), req.ID)
	if err != nil {
		return util.ToWebServiceResponse("Id tidak terdaftar sebagai penembak", http.StatusUnauthorized, nil)
	}

	if shooter.ScorerID != req.ScorerID {
		return util.ToWebServiceResponse("Tidak diperbolehkan menampilkan data akun penembak ujian lain", http.StatusUnauthorized, nil)
	}

	return util.ToWebServiceResponse("Data penembak ditemukan", http.StatusOK, gin.H{
		"shooter": model.Shooter{
			ID:        shooter.ID,
			ScorerID:  shooter.ScorerID,
			Name:      shooter.Name,
			Province:  shooter.Province,
			Club:      shooter.Club,
			CreatedAt: shooter.CreatedAt.Time,
			UpdatedAt: shooter.UpdatedAt.Time,
		}})
}
