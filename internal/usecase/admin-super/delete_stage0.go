package usecase

import (
	"context"
	"net/http"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
)

func (usecase *adminSuperUsecaseImpl) DeleteStage0(req model.ByIdRequest) model.WebServiceResponse {
	if err := usecase.Store.DeleteStage0(context.Background(), req.ID); err != nil {
		return util.ToWebServiceResponse("Gagal menghapus data ujian kualifikasi: "+err.Error(), http.StatusInternalServerError, nil)
	}

	return util.ToWebServiceResponse("Berhasil menghapus data ujian kualifikasi", http.StatusOK, nil)
}
