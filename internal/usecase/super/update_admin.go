package usecase

import (
	"context"
	"fmt"
	"net/http"

	respositoryModel "github.com/FadhilAF/perbakin-be/internal/repository/postgres/sqlc"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
)

func (usecase *superUsecaseImpl) UpdateAdmin(req model.UpdateUserRequest) model.WebServiceResponse {
	admin, err := usecase.Store.GetAdmin(context.Background(), req.ID)
	if err != nil {
		return util.ToWebServiceResponse("Tidak ditemukan admin dengan ID tersebut", http.StatusNotFound, nil)
	}

	passwordHash, err := util.HashPassword(req.Data.Password)
	if err != nil {
		return util.ToWebServiceResponse("Gagal proses hash password: "+err.Error(), http.StatusInternalServerError, nil)
	}

	err = usecase.Store.UpdateAdmin(context.Background(), respositoryModel.UpdateAdminParams{
		ID:       admin.ID,
		Username: req.Data.Username,
		Password: passwordHash,
		Name:     req.Data.Name,
	})
	if err != nil {
		fmt.Println(err)
		return util.ToWebServiceResponse("Gagal mengubah admin", http.StatusInternalServerError, nil)
	}

	return util.ToWebServiceResponse("Berhasil mengubah admin", http.StatusOK, nil)
}
