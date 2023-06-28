package usecase

import (
	"context"
	"net/http"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (usecase *adminSuperUsecaseImpl) GetScorersByExamId(req model.ByExamIdRequest) model.WebServiceResponse {
	scorers, err := usecase.Store.GetScorersByExamId(context.Background(), req.ExamID)
	if err != nil {
		return util.ToWebServiceResponse("Gagal mendapatkan data penguji", http.StatusInternalServerError, nil)
	}

	var scorersData []model.ScorerDisplayData
	for _, scorer := range scorers {
		scorersData = append(scorersData, model.ScorerDisplayData{
			ID:        scorer.ID,
			Name:      scorer.Name,
			ImagePath: scorer.ImagePath,
		})
	}

	return util.ToWebServiceResponse("Berhasil mendapatkan data penguji", http.StatusOK, gin.H{
		"scorers": scorersData,
	})
}
