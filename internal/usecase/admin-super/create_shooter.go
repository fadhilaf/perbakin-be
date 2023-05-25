package usecase

import (
	"context"
	"net/http"

	"github.com/FadhilAF/perbakin-be/internal/model"
	repositoryModel "github.com/FadhilAF/perbakin-be/internal/repository/postgres/sqlc"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (usecase *adminSuperUsecaseImpl) CreateShooter(req model.CreateShooterRequest) model.WebServiceResponse {
	shooter, err := usecase.Store.CreateShooter(context.Background(), repositoryModel.CreateShooterParams{
		ScorerID:  req.ScorerID,
		Name:      req.Body.Name,
		ImagePath: req.ImagePath,
		Province:  req.Body.Province,
		Club:      req.Body.Club,
	})
	if err != nil {
		return util.ToWebServiceResponse("Gagal membuat penembak: "+err.Error(), http.StatusInternalServerError, nil)
	}

	return util.ToWebServiceResponse("Berhasil membuat penembak", http.StatusCreated, gin.H{
		"shooter": model.Shooter{
			ID:        shooter.ID,
			ScorerID:  shooter.ScorerID,
			Name:      shooter.Name,
			ImagePath: shooter.ImagePath,
			Province:  shooter.Province,
			Club:      shooter.Club,
			CreatedAt: shooter.CreatedAt.Time,
			UpdatedAt: shooter.UpdatedAt.Time,
		}})
}
