package usecase

import (
	"context"
	"net/http"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (usecase *superUsecaseImpl) GetAllAdmins() model.WebServiceResponse {
	admins, err := usecase.Store.GetAllAdmins(context.Background())
	if err != nil {
		return util.ToWebServiceResponse("Gagal mendapatkan data admin", http.StatusInternalServerError, nil)
	}

	var adminsData []model.OperatorDisplayData
	for _, admin := range admins {
		adminsData = append(adminsData, model.OperatorDisplayData{
			ID:        admin.ID,
			Exam:      admin.Exam,
			Name:      admin.Name,
			CreatedAt: admin.CreatedAt.Time,
			UpdatedAt: admin.UpdatedAt.Time,
		})
	}

	return util.ToWebServiceResponse("Berhasil mendapatkan data admin", http.StatusOK, gin.H{
		"admins": adminsData,
	})
}
