package usecase

import (
	"context"
	"net/http"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (usecase *allUsecaseImpl) GetStage0ById(req model.ByIdRequest) model.WebServiceResponse {
	stage0, err := usecase.Store.GetStage0ById(context.Background(), req.ID)
	if err != nil {
		return util.ToWebServiceResponse("Hasil ujian kualifikasi tidak ada: "+err.Error(), http.StatusNotFound, nil)
	}

	return util.ToWebServiceResponse("Berhasil mendapatkan hasil ujian kualifikasi", http.StatusOK, gin.H{
		"stage_0": model.Stage0{
			ID:          stage0.ID,
			ResultID:    stage0.ResultID,
			Status:      string(stage0.Status),
			Series1:     util.ScoresToIntArray(stage0.Series1),
			Series2:     util.ScoresToIntArray(stage0.Series2),
			Series3:     util.ScoresToIntArray(stage0.Series3),
			Series4:     util.ScoresToIntArray(stage0.Series4),
			Series5:     util.ScoresToIntArray(stage0.Series5),
			Checkmarks:  util.CheckmarksToArray(stage0.Checkmarks),
			ShooterSign: stage0.ShooterSign,
			ScorerSign:  stage0.ScorerSign,
			CreatedAt:   stage0.CreatedAt.Time,
			UpdatedAt:   stage0.UpdatedAt.Time,
		}})
}
