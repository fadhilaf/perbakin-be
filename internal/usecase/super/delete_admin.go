package usecase

import (
	"context"
	"net/http"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
)

func (usecase *superUsecaseImpl) DeleteAdmin(req model.DeleteUserRequest) model.WebServiceResponse {
	admin, err := usecase.Store.GetAdmin(context.Background(), req.ID)
	if err != nil {
		return util.ToWebServiceResponse("Tidak ditemukan admin dengan ID tersebut", http.StatusNotFound, nil)
	}

	if err = usecase.Store.DeleteUser(context.Background(), admin.UserID); err != nil {
		return util.ToWebServiceResponse("Gagal menghapus admin", http.StatusInternalServerError, nil)
	}

	return util.ToWebServiceResponse("Berhasil menghapus admin", http.StatusOK, nil)
}
