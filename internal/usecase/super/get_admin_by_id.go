package usecase

import (
	"context"
	"net/http"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (usecase *superUsecaseImpl) GetAdminById(req model.ByIdRequest) model.WebServiceResponse {
	admin, err := usecase.Store.GetAdminById(context.Background(), req.ID)
	if err != nil {
		return util.ToWebServiceResponse("Id tidak terdaftar sebagai admin", http.StatusUnauthorized, nil)
	}

	return util.ToWebServiceResponse("Data admin ditemukan", http.StatusOK, gin.H{
		"admin": model.Admin{
			ID: admin.ID,
			User: model.User{
				ID:       admin.UserID,
				Username: admin.Username,
				Name:     admin.Name,
			},
		}})
}
