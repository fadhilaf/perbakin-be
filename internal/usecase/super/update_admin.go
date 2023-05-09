package usecase

import (
	"context"
	"fmt"
	"net/http"

	respositoryModel "github.com/FadhilAF/perbakin-be/internal/repository/postgres/sqlc"
	"github.com/gin-gonic/gin"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
)

func (usecase *superUsecaseImpl) UpdateAdmin(req model.UpdateOperatorRequest) model.WebServiceResponse {
	admin, err := usecase.Store.GetAdminById(context.Background(), req.ID)
	if err != nil {
		return util.ToWebServiceResponse("Tidak ditemukan admin dengan ID tersebut", http.StatusNotFound, nil)
	}

	if admin.ExamID != req.ExamID {
		return util.ToWebServiceResponse("Tidak dapat mengubah admin ujian lain", http.StatusUnauthorized, nil)
	}

	passwordHash, err := util.HashPassword(req.Body.Password)
	if err != nil {
		return util.ToWebServiceResponse("Gagal proses hash password: "+err.Error(), http.StatusInternalServerError, nil)
	}

	newAdmin, err := usecase.Store.UpdateAdmin(context.Background(), respositoryModel.UpdateAdminParams{
		ID:       req.ID,
		Username: req.Body.Username,
		Password: passwordHash,
		Name:     req.Body.Name,
	})
	if err != nil {
		fmt.Println(err)
		return util.ToWebServiceResponse("Gagal mengubah admin", http.StatusInternalServerError, nil)
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
