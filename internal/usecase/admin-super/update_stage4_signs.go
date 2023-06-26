package usecase

import (
	"context"
	"net/http"

	repositoryModel "github.com/FadhilAF/perbakin-be/internal/repository/postgres/sqlc"
	"github.com/gin-gonic/gin"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
)

func (usecase *adminSuperUsecaseImpl) UpdateStage4Signs(req model.UpdateStageSignsRequest) model.WebServiceResponse {
	updatedSign, err := usecase.Store.UpdateStage4Signs(context.Background(), repositoryModel.UpdateStage4SignsParams{
		ID:          req.ID,
		ShooterSign: req.ShooterSign,
		ScorerSign:  req.ScorerSign,
	})
	if err != nil {
		return util.ToWebServiceResponse("Gagal mengubah tanda tangan stage 4: "+err.Error(), http.StatusInternalServerError, nil)
	}

	return util.ToWebServiceResponse("Berhasil mengubah tanda tangan stage 4", http.StatusOK, gin.H{
		"stage_4": model.UpdateStageSignsResponse{
			ShooterSign: updatedSign.ShooterSign,
			ScorerSign:  updatedSign.ScorerSign,
			UpdatedAt:   updatedSign.UpdatedAt.Time,
		},
	})
}
