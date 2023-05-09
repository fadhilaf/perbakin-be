package usecase

import (
	"context"
	"net/http"

	"github.com/FadhilAF/perbakin-be/internal/model"
	repositoryModel "github.com/FadhilAF/perbakin-be/internal/repository/postgres/sqlc"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (usecase *superUsecaseImpl) CreateAdmin(req model.CreateOperatorRequest) model.WebServiceResponse {
	if _, err := usecase.Store.GetUserByUsername(context.Background(), req.Body.Username); err == nil {
		return util.ToWebServiceResponse("Username sudah digunakan", http.StatusConflict, nil)
	}

	passwordHash, err := util.HashPassword(req.Body.Password)
	if err != nil {
		return util.ToWebServiceResponse("Gagal proses hash password: "+err.Error(), http.StatusInternalServerError, nil)
	}

	admin, err := usecase.Store.CreateAdmin(context.Background(), repositoryModel.CreateAdminParams{
		ExamID:   req.ExamID,
		Username: req.Body.Username,
		Password: passwordHash,
		Name:     req.Body.Name,
	})
	if err != nil {
		return util.ToWebServiceResponse("Gagal membuat admin: "+err.Error(), http.StatusInternalServerError, nil)
	}

	return util.ToWebServiceResponse("Berhasil membuat admin", http.StatusCreated, gin.H{
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
