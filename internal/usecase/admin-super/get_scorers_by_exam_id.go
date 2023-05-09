package usecase

import (
	"context"
	"net/http"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (usecase *adminSuperUsecaseImpl) GetScorersByExamId(req model.GetOperatorsByExamIdRequest) model.WebServiceResponse {
	scorers, err := usecase.Store.GetScorersByExamId(context.Background(), req.ExamID)
	if err != nil {
		return util.ToWebServiceResponse("Gagal mendapatkan data penguji", http.StatusInternalServerError, nil)
	}

	var scorersData []model.Operator
	for _, scorer := range scorers {
		scorersData = append(scorersData, model.Operator{
			ID:     scorer.ID,
			ExamID: scorer.ExamID,
			User: model.User{
				ID:       scorer.UserID,
				Username: scorer.Username,
				Name:     scorer.Name,
			},
		})
	}

	return util.ToWebServiceResponse("Berhasil mendapatkan data penguji", http.StatusOK, gin.H{
		"scorers": scorersData,
	})
}
