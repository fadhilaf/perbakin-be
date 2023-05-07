package usecase

import (
	"context"
	"net/http"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (usecase *adminUsecaseImpl) GetAdminByUserId(req model.GetByUserIdRequest) model.WebServiceResponse {
	admin, err := usecase.Store.GetAdminByUserId(context.Background(), req.UserID)
	if err != nil {
		return util.ToWebServiceResponse("User tidak terdaftar sebagai admin", http.StatusUnauthorized, nil)
	}

	return util.ToWebServiceResponse("User adalah admin", http.StatusOK, gin.H{
		"admin": model.Admin{
			ID: admin.ID,
			User: model.User{
				ID:       admin.UserID,
				Username: admin.Username,
				Name:     admin.Name,
			},
		}})
}
