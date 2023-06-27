package usecase

import (
	"context"
	"net/http"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
)

func (usecase *adminSuperUsecaseImpl) DeleteStage6(req model.ByIdRequest) model.WebServiceResponse {
	if err := usecase.Store.DeleteStage6(context.Background(), req.ID); err != nil {
		return util.ToWebServiceResponse("Gagal menghapus data ujian stage 6: "+err.Error(), http.StatusInternalServerError, nil)
	}

	return util.ToWebServiceResponse("Berhasil menghapus data ujian stage 6", http.StatusOK, nil)
}
