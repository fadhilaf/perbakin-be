package usecase

import (
	"context"
	"net/http"

	repositoryModel "github.com/FadhilAF/perbakin-be/internal/repository/postgres/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
)

func (usecase *superUsecaseImpl) UpdateAdmin(req model.UpdateOperatorRequest) model.WebServiceResponse {
	if existingAdmin, err := usecase.Store.GetUserByUsername(context.Background(), req.Username); err == nil {
		if existingAdmin != req.ID {
			return util.ToWebServiceResponse("Username sudah digunakan", http.StatusConflict, nil)
		}
	}

	var query repositoryModel.UpdateAdminParams

	if req.Password.Valid {
		passwordHash, err := util.HashPassword(req.Password.String)
		if err != nil {
			return util.ToWebServiceResponse("Gagal proses hash password: "+err.Error(), http.StatusInternalServerError, nil)
		}

		var passwordHashText pgtype.Text
		passwordHashText.Scan(passwordHash)

		query = repositoryModel.UpdateAdminParams{
			ID:       req.ID,
			Username: req.Username,
			Password: passwordHashText,
			Name:     req.Name,
		}
	} else {
		query = repositoryModel.UpdateAdminParams{
			ID:       req.ID,
			Username: req.Username,
			Name:     req.Name,
		}
	}

	newAdmin, err := usecase.Store.UpdateAdmin(context.Background(), query)
	if err != nil {
		return util.ToWebServiceResponse("Gagal mengubah admin: "+err.Error(), http.StatusInternalServerError, nil)
	}

	return util.ToWebServiceResponse("Berhasil mengubah admin", http.StatusOK, gin.H{
		"admin": model.Admin{
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
