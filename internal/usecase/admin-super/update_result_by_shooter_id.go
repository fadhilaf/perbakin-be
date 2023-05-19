package usecase

import (
	"context"
	"net/http"

	respositoryModel "github.com/FadhilAF/perbakin-be/internal/repository/postgres/sqlc"
	"github.com/gin-gonic/gin"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
)

func (usecase *adminSuperUsecaseImpl) UpdateResultByShooterId(req model.UpdateResultByShooterIdRequest) model.WebServiceResponse {
	newResult, err := usecase.Store.UpdateResultByShooterId(context.Background(), respositoryModel.UpdateResultByShooterIdParams{
		ShooterID: req.ShooterID,
		Failed:    req.Body.Failed,
		Stage:     req.Body.Stage,
	})
	if err != nil {
		return util.ToWebServiceResponse("Gagal mengubah hasil ujian: "+err.Error(), http.StatusInternalServerError, nil)
	}

	return util.ToWebServiceResponse("Berhasil mengubah hasil ujian", http.StatusOK, gin.H{
		"result": model.Result{
			ID:        newResult.ID,
			ShooterID: newResult.ShooterID,
			Failed:    newResult.Failed,
			Stage:     newResult.Stage.(string),
			CreatedAt: newResult.CreatedAt.Time,
			UpdatedAt: newResult.UpdatedAt.Time,
		}})
}
