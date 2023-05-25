package usecase

import (
	"context"
	"net/http"
	"strconv"

	repositoryModel "github.com/FadhilAF/perbakin-be/internal/repository/postgres/sqlc"
	"github.com/gin-gonic/gin"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
)

func (usecase *scorerUsecaseImpl) UpdateStage0NextSeries(req model.ByIdRequest) model.WebServiceResponse {
	status, _ := usecase.Store.GetStage0Status(context.Background(), req.ID)
	if status == "5" {
		return util.ToWebServiceResponse("Tidak dapat melanjutkan, batas seri kualifikasi hanya 5", http.StatusForbidden, nil)
	}

	intStatus, _ := strconv.Atoi(string(status))

	newStatus, err := usecase.Store.UpdateStage0NextSeries(context.Background(), repositoryModel.UpdateStage0NextSeriesParams{
		ID:     req.ID,
		Status: repositoryModel.Stage0Status(strconv.Itoa(intStatus + 1)),
	})
	if err != nil {
		return util.ToWebServiceResponse("Gagal melanjutkan seri kualifikasi: "+err.Error(), http.StatusInternalServerError, nil)
	}

	return util.ToWebServiceResponse("Berhasil melanjutkan seri kualifikasi ke seri "+string(newStatus), http.StatusOK, gin.H{
		"series": newStatus,
	})
}
