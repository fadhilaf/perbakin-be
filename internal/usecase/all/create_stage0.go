package usecase

import (
	"context"
	"net/http"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (usecase *allUsecaseImpl) CreateStage0(req model.ByResultIdRequest) model.WebServiceResponse {
	if _, err := usecase.Store.GetStage0(context.Background(), req.ResultID); err == nil {
		return util.ToWebServiceResponse("Hasil ujian kualifikasi sudah ada", http.StatusConflict, nil)
	}

	stage0, err := usecase.Store.CreateStage0(context.Background(), req.ResultID)
	if err != nil {
		return util.ToWebServiceResponse("Gagal membuat hasil ujian kualifikasi: "+err.Error(), http.StatusInternalServerError, nil)
	}

	return util.ToWebServiceResponse("Berhasil membuat hasil ujian kualifikasi", http.StatusCreated, gin.H{
		"stage0": model.Stage0{
			ID:        stage0.ID,
			ResultID:  stage0.ResultID,
			Status:    stage0.Status.(string),
			Series1:   stage0.Series1.([]int),
			Series2:   stage0.Series2.([]int),
			Series3:   stage0.Series3.([]int),
			Series4:   stage0.Series4.([]int),
			Series5:   stage0.Series5.([]int),
			CreatedAt: stage0.CreatedAt.Time,
			UpdatedAt: stage0.UpdatedAt.Time,
		},
	})
}
