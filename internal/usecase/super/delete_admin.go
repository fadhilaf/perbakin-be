package usecase

import (
	"context"
	"net/http"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
)

func (usecase *superUsecaseImpl) DeleteAdmin(req model.UserByUserIdRequest) model.WebServiceResponse {
	if err := usecase.Store.DeleteUser(context.Background(), req.UserID); err != nil {
		return util.ToWebServiceResponse("Gagal menghapus admin: "+err.Error(), http.StatusInternalServerError, nil)
	}

	return util.ToWebServiceResponse("Berhasil menghapus admin", http.StatusOK, nil)
}
