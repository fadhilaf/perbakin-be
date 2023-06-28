package usecase

import (
	"context"
	"net/http"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (usecase *adminUsecaseImpl) GetAdminByUserId(req model.UserByUserIdRequest) model.WebServiceResponse {
	admin, err := usecase.Store.GetAdminByUserId(context.Background(), req.UserID)
	if err != nil {
		return util.ToWebServiceResponse("User tidak terdaftar sebagai admin", http.StatusUnauthorized, nil)
	}

	return util.ToWebServiceResponse("Data admin ditemukan", http.StatusOK, gin.H{
		"admin": model.Admin{
			ID:     admin.ID,
			ExamID: admin.ExamID,
			User: model.User{
				ID:       admin.UserID,
				Username: admin.Username,
				Name:     admin.Name,
			},
		}})
}
