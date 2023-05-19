package usecase

import (
	"context"
	"net/http"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (usecase *allUsecaseImpl) GetResultByShooterId(req model.ByShooterIdRequest) model.WebServiceResponse {
	result, err := usecase.Store.GetResultByShooterId(context.Background(), req.ShooterID)
	if err != nil {
		return util.ToWebServiceResponse("Gagal mendapatkan hasil ujian", http.StatusInternalServerError, nil)
	}

	return util.ToWebServiceResponse("Berhasil mendapatkan hasil ujian", http.StatusOK, gin.H{
		"result": model.Result{
			ID:        result.ID,
			ShooterID: result.ShooterID,
			Failed:    result.Failed,
			Stage:     result.Stage.(string),
			CreatedAt: result.CreatedAt.Time,
			UpdatedAt: result.UpdatedAt.Time,
		}})
}
