package usecase

import (
	"context"
	"net/http"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (usecase *allUsecaseImpl) CreateResult(req model.ByShooterIdRequest) model.WebServiceResponse {
	if _, err := usecase.Store.GetResultRelationAndStatusByShooterId(context.Background(), req.ShooterID); err == nil {
		return util.ToWebServiceResponse("Hasil ujian sudah ada", http.StatusConflict, nil)
	}

	result, err := usecase.Store.CreateResult(context.Background(), req.ShooterID)
	if err != nil {
		return util.ToWebServiceResponse("Gagal membuat hasil ujian: "+err.Error(), http.StatusInternalServerError, nil)
	}

	return util.ToWebServiceResponse("Berhasil membuat hasil ujian", http.StatusCreated, gin.H{
		"result": model.Result{
			ID:        result.ID,
			ShooterID: result.ShooterID,
			Failed:    result.Failed,
			Stage:     string(result.Stage),
			CreatedAt: result.CreatedAt.Time,
			UpdatedAt: result.UpdatedAt.Time,
		}})
}
