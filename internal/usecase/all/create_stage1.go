package usecase

import (
	"context"
	"net/http"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (usecase *allUsecaseImpl) CreateStage1(req model.ByResultIdRequest) model.WebServiceResponse {
	if _, err := usecase.Store.GetStage1RelationByResultId(context.Background(), req.ResultID); err == nil {
		return util.ToWebServiceResponse("Hasil ujian stage 1 sudah ada", http.StatusConflict, nil)
	}

	stage1, err := usecase.Store.CreateStage1(context.Background(), req.ResultID)
	if err != nil {
		return util.ToWebServiceResponse("Gagal membuat hasil ujian stage 1: "+err.Error(), http.StatusInternalServerError, nil)
	}

	return util.ToWebServiceResponse("Berhasil membuat hasil ujian stage 1", http.StatusCreated, gin.H{
		"stage_1": model.Stage1{
			ID:       stage1.ID,
			ResultID: stage1.ResultID,
			Status:   string(stage1.Status),
			// Scores1:     util.ScoresToIntArray(stage1.Scores1),
			// Duration1:   util.ScoresToIntArray(stage1.Duration1),
			// Scores2:     util.ScoresToIntArray(stage1.Scores2),
			// Duration2:   util.ScoresToIntArray(stage1.Duration2),
			// Scores3:     util.ScoresToIntArray(stage1.Scores3),
			// Duration3:   util.ScoresToIntArray(stage1.Duration3),
			// Scores4:     util.ScoresToIntArray(stage1.Scores4),
			// Duration4:   util.ScoresToIntArray(stage1.Duration4),
			// Scores5:     util.ScoresToIntArray(stage1.Scores5),
			// Duration5:   util.ScoresToIntArray(stage1.Duration5),
			// Scores6:     util.ScoresToIntArray(stage1.Scores6),
			// Duration6:   util.ScoresToIntArray(stage1.Duration6),
			ShooterSign: stage1.ShooterSign,
			ScorerSign:  stage1.ScorerSign,
			CreatedAt:   stage1.CreatedAt.Time,
			UpdatedAt:   stage1.UpdatedAt.Time,
		},
	})
}
