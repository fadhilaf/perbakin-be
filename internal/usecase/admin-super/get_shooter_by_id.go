package usecase

import (
	"context"
	"net/http"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (usecase *adminSuperUsecaseImpl) GetShooterById(req model.ByIdRequest) model.WebServiceResponse {
	shooter, _ := usecase.Store.GetShooterById(context.Background(), req.ID)

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
