package usecase

import (
	"context"
	"net/http"

	repositoryModel "github.com/FadhilAF/perbakin-be/internal/repository/postgres/sqlc"
	"github.com/gin-gonic/gin"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
)

func (usecase *adminSuperUsecaseImpl) UpdateStage0Signs(req model.UpdateStageSignsRequest) model.WebServiceResponse {
	updatedSign, err := usecase.Store.UpdateStage0Signs(context.Background(), repositoryModel.UpdateStage0SignsParams{
		ID:          req.ID,
		ShooterSign: req.ShooterSign,
		ScorerSign:  req.ScorerSign,
	})
	if err != nil {
		return util.ToWebServiceResponse("Gagal mengubah tanda tangan babak kualifikasi: "+err.Error(), http.StatusInternalServerError, nil)
	}

	return util.ToWebServiceResponse("Berhasil mengubah tanda tangan babak kualifikasi", http.StatusOK, gin.H{
		"stage_0": model.UpdateStageSignsResponse{
			ShooterSign: updatedSign.ShooterSign,
			ScorerSign:  updatedSign.ScorerSign,
			UpdatedAt:   updatedSign.UpdatedAt.Time,
		},
	})
}
