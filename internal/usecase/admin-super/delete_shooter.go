package usecase

import (
	"context"
	"net/http"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
)

func (usecase *adminSuperUsecaseImpl) DeleteShooter(req model.ByIdRequest) model.WebServiceResponse {
	if err := usecase.Store.DeleteShooter(context.Background(), req.ID); err != nil {
		return util.ToWebServiceResponse("Gagal menghapus penembak: "+err.Error(), http.StatusInternalServerError, nil)
	}

	return util.ToWebServiceResponse("Berhasil menghapus penembak", http.StatusOK, nil)
}
