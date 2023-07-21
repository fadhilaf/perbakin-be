package usecase

import (
	"context"
	"net/http"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (usecase *adminSuperUsecaseImpl) GetResultsByExamId(req model.ByExamIdRequest) model.WebServiceResponse {
	results, err := usecase.Store.GetResultsByExamId(context.Background(), req.ExamID)
	if err != nil {
		return util.ToWebServiceResponse("Gagal mendapatkan hasil ujian semua penembak ujian ini: "+err.Error(), http.StatusNotFound, nil)
	}

	var resultsAndShooters []model.ResultAndShooter
	for _, result := range results {
		resultsAndShooters = append(resultsAndShooters, model.ResultAndShooter{
			ID:       result.ID,
			Name:     result.Name,
			Province: result.Province,
			Club:     result.Club,
			Failed:   result.Failed,
			Stage:    string(result.Stage),
		})
	}

	return util.ToWebServiceResponse("Berhasil mendapatkan hasil ujian semua penembak ujian ini", http.StatusOK, gin.H{
		"results": resultsAndShooters})
}
