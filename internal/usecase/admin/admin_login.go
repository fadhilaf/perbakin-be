package usecase

import (
	"context"
	"net/http"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (usecase *adminUsecaseImpl) AdminLogin(req model.LoginRequest) model.WebServiceResponse {
	admin, err := usecase.Store.GetAdminByUsername(context.Background(), req.Username)

	if err != nil {
		return util.ToWebServiceResponse("Username yang dimasukkan salah", http.StatusNotFound, nil)
	}

	if err := util.ComparePassword(req.Password, admin.Password); err != nil {
		return util.ToWebServiceResponse("Password yang dimasukkan salah", http.StatusUnauthorized, nil)
	}

	return util.ToWebServiceResponse("Login berhasil", http.StatusOK, gin.H{
		"admin": model.Admin{
			ID: admin.ID,
			User: model.User{
				ID:       admin.UserID,
				Username: admin.Username,
				Name:     admin.Name,
			},
		}})
}
