package usecase

import (
	"context"
	"net/http"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (usecase *superUsecaseImpl) GetSuperByUserId(req model.UserByUserIdRequest) model.WebServiceResponse {
	super, err := usecase.Store.GetSuperByUserId(context.Background(), req.UserID)
	if err != nil {
		return util.ToWebServiceResponse("User tidak terdaftar sebagai super admin", http.StatusUnauthorized, nil)
	}

	return util.ToWebServiceResponse("Data super admin ditemukan", http.StatusOK, gin.H{
		"super": model.Super{
			ID: super.ID,
			User: model.User{
				ID:       super.UserID,
				Username: super.Username,
				Name:     super.Name,
			},
		}})
}
