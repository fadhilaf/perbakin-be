package usecase

import (
	"context"
	"net/http"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
)

func (usecase *adminSuperUsecaseImpl) DeleteResultByShooterId(req model.ByShooterIdRequest) model.WebServiceResponse {
	result, err := usecase.Store.GetResultByShooterId(context.Background(), req.ShooterID)
	if err != nil {
		return util.ToWebServiceResponse("Tidak ditemukan hasil ujian dengan ID yang diberikan", http.StatusNotFound, nil)
	}

	if err = usecase.Store.DeleteResultByShooterId(context.Background(), result.ID); err != nil {
		return util.ToWebServiceResponse("Gagal menghapus hasil ujian: "+err.Error(), http.StatusInternalServerError, nil)
	}

	return util.ToWebServiceResponse("Berhasil menghapus hasil ujian", http.StatusOK, nil)
}
