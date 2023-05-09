package usecase

import (
	"context"
	"net/http"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
)

func (usecase *superUsecaseImpl) DeleteAdmin(req model.OperatorByIdRequest) model.WebServiceResponse {
	admin, err := usecase.Store.GetAdminById(context.Background(), req.ID)
	if err != nil {
		return util.ToWebServiceResponse("Tidak ditemukan admin dengan ID tersebut", http.StatusNotFound, nil)
	}

	if admin.ExamID != req.ExamID {
		return util.ToWebServiceResponse("Tidak diperbolehkan menghapus admin ujian lain", http.StatusUnauthorized, nil)
	}

	if err = usecase.Store.DeleteUser(context.Background(), admin.UserID); err != nil {
		return util.ToWebServiceResponse("Gagal menghapus admin: "+err.Error(), http.StatusInternalServerError, nil)
	}

	return util.ToWebServiceResponse("Berhasil menghapus admin", http.StatusOK, nil)
}
