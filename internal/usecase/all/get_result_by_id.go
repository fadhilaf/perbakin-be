package usecase

import (
	"context"
	"net/http"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (usecase *allUsecaseImpl) GetResultById(req model.ByIdRequest) model.WebServiceResponse {
	result, err := usecase.Store.GetResultById(context.Background(), req.ID)
	if err != nil {
		return util.ToWebServiceResponse("Hasil ujian tidak ada: "+err.Error(), http.StatusNotFound, nil)
	}

	return util.ToWebServiceResponse("Berhasil mendapatkan hasil ujian", http.StatusOK, gin.H{
		"result": model.Result{
			ID:        result.ID,
			ShooterID: result.ShooterID,
			Failed:    result.Failed,
			Stage:     string(result.Stage.Stages),
			CreatedAt: result.CreatedAt.Time,
			UpdatedAt: result.UpdatedAt.Time,
		}})
}
