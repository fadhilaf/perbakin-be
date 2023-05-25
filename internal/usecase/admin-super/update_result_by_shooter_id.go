package usecase

import (
	"context"
	"net/http"

	repositoryModel "github.com/FadhilAF/perbakin-be/internal/repository/postgres/sqlc"
	"github.com/gin-gonic/gin"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
)

func (usecase *adminSuperUsecaseImpl) UpdateResult(req model.UpdateResultRequest) model.WebServiceResponse {
	newResult, err := usecase.Store.UpdateResult(context.Background(), repositoryModel.UpdateResultParams{
		ID:     req.ID,
		Failed: req.Body.Failed,
		Stage:  repositoryModel.NullStages{Stages: repositoryModel.Stages(req.Body.Stage), Valid: true},
	})
	if err != nil {
		return util.ToWebServiceResponse("Gagal mengubah hasil ujian: "+err.Error(), http.StatusInternalServerError, nil)
	}

	return util.ToWebServiceResponse("Berhasil mengubah hasil ujian", http.StatusOK, gin.H{
		"result": model.Result{
			ID:        newResult.ID,
			ShooterID: newResult.ShooterID,
			Failed:    newResult.Failed,
			Stage:     string(newResult.Stage.Stages),
			CreatedAt: newResult.CreatedAt.Time,
			UpdatedAt: newResult.UpdatedAt.Time,
		}})
}
