package usecase

import (
	"context"
	"net/http"

	repositoryModel "github.com/FadhilAF/perbakin-be/internal/repository/postgres/sqlc"
	"github.com/gin-gonic/gin"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
)

func (usecase *adminSuperUsecaseImpl) UpdateStage0(req model.UpdateStage0Request) model.WebServiceResponse {
	newStage0, err := usecase.Store.UpdateStage0(context.Background(), repositoryModel.UpdateStage0Params{
		ID:         req.ID,
		Status:     repositoryModel.Stage0Status(req.Status),
		Series1:    req.Series1,
		Series2:    req.Series2,
		Series3:    req.Series3,
		Series4:    req.Series4,
		Series5:    req.Series5,
		Checkmarks: req.Checkmarks,
	})
	if err != nil {
		return util.ToWebServiceResponse("Gagal mengubah hasil ujian: "+err.Error(), http.StatusInternalServerError, nil)
	}

	return util.ToWebServiceResponse("Berhasil mengubah hasil ujian", http.StatusOK, gin.H{
		"stage_0": model.UpdateStage0Response{
			ID:         newStage0.ID,
			ResultID:   newStage0.ResultID,
			Status:     string(newStage0.Status),
			Series1:    util.ScoresToIntArray(newStage0.Series1),
			Series2:    util.ScoresToIntArray(newStage0.Series2),
			Series3:    util.ScoresToIntArray(newStage0.Series3),
			Series4:    util.ScoresToIntArray(newStage0.Series4),
			Series5:    util.ScoresToIntArray(newStage0.Series5),
			Checkmarks: util.CheckmarksToBoolArray(newStage0.Checkmarks),
			CreatedAt:  newStage0.CreatedAt.Time,
			UpdatedAt:  newStage0.UpdatedAt.Time,
		}})
}
