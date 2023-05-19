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
		return util.ToWebServiceResponse("Gagal mendapatkan data admin: "+err.Error(), http.StatusInternalServerError, nil)
	}

	var adminsData []model.OperatorDisplayExamData
	for _, admin := range admins {
		adminsData = append(adminsData, model.OperatorDisplayExamData{
			Exam: admin.Exam,
			Name: admin.Name,
		})
	}

	return util.ToWebServiceResponse("Berhasil mendapatkan data admin", http.StatusOK, gin.H{
		"admins": adminsData,
	})
}
