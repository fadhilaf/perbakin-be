package usecase

import (
	"context"
	"net/http"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (usecase *superUsecaseImpl) SuperLogin(req model.LoginRequest) model.WebServiceResponse {
	super, err := usecase.Store.GetSuperByUsername(context.Background(), req.Username)

	if err != nil {
		return util.ToWebServiceResponse("Username yang dimasukkan tidak ditemukan", http.StatusNotFound, nil)
	}

	if err := util.ComparePassword(req.Password, super.Password); err != nil {
		return util.ToWebServiceResponse("Password yang dimasukkan salah", http.StatusUnauthorized, nil)
	}

	return util.ToWebServiceResponse("Login berhasil", http.StatusOK, gin.H{
		"super": model.Super{
			ID: super.ID,
			User: model.User{
				ID:       super.UserID,
				Username: super.Username,
				Name:     super.Name,
			},
		}})
}
