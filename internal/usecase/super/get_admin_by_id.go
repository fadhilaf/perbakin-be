package usecase

import (
	"context"
	"net/http"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (usecase *superUsecaseImpl) GetAdminById(req model.OperatorByIdRequest) model.WebServiceResponse {
	admin, err := usecase.Store.GetAdminById(context.Background(), req.ID)
	if err != nil {
		return util.ToWebServiceResponse("Id tidak terdaftar sebagai admin", http.StatusUnauthorized, nil)
	}

	if admin.ExamID != req.ExamID {
		return util.ToWebServiceResponse("Tidak diperbolehkan menampilkan data akun admin ujian lain", http.StatusUnauthorized, nil)
	}

	return util.ToWebServiceResponse("Data admin ditemukan", http.StatusOK, gin.H{
		"admin": model.Operator{
			ID:     admin.ID,
			ExamID: admin.ExamID,
			User: model.User{
				ID:       admin.UserID,
				Username: admin.Username,
				Name:     admin.Name,
			},
		}})
}
