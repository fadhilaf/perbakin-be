package usecase

import (
	"context"
	"net/http"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (usecase *superUsecaseImpl) GetAdminsByExamId(req model.ByExamIdRequest) model.WebServiceResponse {
	admins, err := usecase.Store.GetAdminsByExamId(context.Background(), req.ExamID)
	if err != nil {
		return util.ToWebServiceResponse("Gagal mendapatkan data admin: "+err.Error(), http.StatusInternalServerError, nil)
	}

	var adminsData []model.AdminDisplayData
	for _, admin := range admins {
		adminsData = append(adminsData, model.AdminDisplayData{
			ID:   admin.ID,
			Name: admin.Name,
		})
	}

	return util.ToWebServiceResponse("Berhasil mendapatkan data admin", http.StatusOK, gin.H{
		"admins": adminsData,
	})
}
