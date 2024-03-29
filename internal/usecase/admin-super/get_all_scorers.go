package usecase

import (
	"context"
	"net/http"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (usecase *adminSuperUsecaseImpl) GetAllScorers() model.WebServiceResponse {
	scorers, err := usecase.Store.GetAllScorers(context.Background())
	if err != nil {
		return util.ToWebServiceResponse("Gagal mendapatkan data penguji", http.StatusInternalServerError, nil)
	}

	var scorersDisplayData []model.ScorerDisplayExamData
	for _, scorer := range scorers {
		scorersDisplayData = append(scorersDisplayData, model.ScorerDisplayExamData{
			Exam:      scorer.Exam,
			Name:      scorer.Name,
			ImagePath: scorer.ImagePath,
		})
	}

	return util.ToWebServiceResponse("Berhasil mendapatkan data penguji", http.StatusOK, gin.H{
		"scorers": scorersDisplayData,
	})
}
