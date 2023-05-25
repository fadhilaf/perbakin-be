package usecase

import (
	"context"
	"net/http"

	repositoryModel "github.com/FadhilAF/perbakin-be/internal/repository/postgres/sqlc"
	"github.com/gin-gonic/gin"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
)

func (usecase *superUsecaseImpl) UpdateAdmin(req model.UpdateOperatorRequest) model.WebServiceResponse {
	passwordHash, err := util.HashPassword(req.Body.Password)
	if err != nil {
		return util.ToWebServiceResponse("Gagal proses hash password: "+err.Error(), http.StatusInternalServerError, nil)
	}

	newAdmin, err := usecase.Store.UpdateAdmin(context.Background(), repositoryModel.UpdateAdminParams{
		ID:       req.ID,
		Username: req.Body.Username,
		Password: passwordHash,
		Name:     req.Body.Name,
	})
	if err != nil {
		return util.ToWebServiceResponse("Gagal mengubah admin: "+err.Error(), http.StatusInternalServerError, nil)
	}

	return util.ToWebServiceResponse("Berhasil mengubah admin", http.StatusOK, gin.H{
		"admin": model.Operator{
			ID:     newAdmin.ID,
			ExamID: newAdmin.ExamID,
			User: model.User{
				ID:       newAdmin.UserID,
				Username: newAdmin.Username,
				Name:     newAdmin.Name,
			},
		},
	})
}
