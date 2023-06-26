package usecase

import (
	"context"
	"net/http"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
)

func (usecase *adminSuperUsecaseImpl) DeleteStage4(req model.ByIdRequest) model.WebServiceResponse {
	if err := usecase.Store.DeleteStage4(context.Background(), req.ID); err != nil {
		return util.ToWebServiceResponse("Gagal menghapus data ujian stage 4: "+err.Error(), http.StatusInternalServerError, nil)
	}

	return util.ToWebServiceResponse("Berhasil menghapus data ujian stage 4", http.StatusOK, nil)
}
