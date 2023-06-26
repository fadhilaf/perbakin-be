package usecase

import (
	"context"
	"net/http"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
)

func (usecase *adminSuperUsecaseImpl) DeleteStage2try2(req model.ByIdRequest) model.WebServiceResponse {
	exists, err := usecase.Store.GetStage2try2ExistById(context.Background(), req.ID)
	if err != nil {
		return util.ToWebServiceResponse("Gagal mengecek keberadaan hasil ujian stage 2 percobaan 2: "+err.Error(), http.StatusInternalServerError, nil)
	}
	if !exists {
		return util.ToWebServiceResponse("Hasil ujian stage 2 percobaan 2 belum ada", http.StatusConflict, nil)
	}

	if err := usecase.Store.DeleteStage2try2(context.Background(), req.ID); err != nil {
		return util.ToWebServiceResponse("Gagal menghapus data ujian stage 2 percobaan 2: "+err.Error(), http.StatusInternalServerError, nil)
	}

	return util.ToWebServiceResponse("Berhasil menghapus data ujian stage 2 percobaan 2", http.StatusOK, nil)
}
