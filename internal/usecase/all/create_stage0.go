package usecase

import (
	"context"
	"net/http"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (usecase *allUsecaseImpl) CreateStage0(req model.ByResultIdRequest) model.WebServiceResponse {
	if _, err := usecase.Store.GetStage0RelationByResultId(context.Background(), req.ResultID); err == nil {
		return util.ToWebServiceResponse("Hasil ujian kualifikasi sudah ada", http.StatusConflict, nil)
	}

	stage0, err := usecase.Store.CreateStage0(context.Background(), req.ResultID)
	if err != nil {
		return util.ToWebServiceResponse("Gagal membuat hasil ujian kualifikasi: "+err.Error(), http.StatusInternalServerError, nil)
	}

	return util.ToWebServiceResponse("Berhasil membuat hasil ujian kualifikasi", http.StatusCreated, gin.H{
		"stage_0": model.Stage0{
			ID:          stage0.ID,
			ResultID:    stage0.ResultID,
			Status:      string(stage0.Status),
			Series1:     util.ScoresToIntArray(stage0.Series1),
			Series2:     util.ScoresToIntArray(stage0.Series2),
			Series3:     util.ScoresToIntArray(stage0.Series3),
			Series4:     util.ScoresToIntArray(stage0.Series4),
			Series5:     util.ScoresToIntArray(stage0.Series5),
			ShooterSign: stage0.ShooterSign,
			ScorerSign:  stage0.ScorerSign,
			CreatedAt:   stage0.CreatedAt.Time,
			UpdatedAt:   stage0.UpdatedAt.Time,
		},
	})
}
