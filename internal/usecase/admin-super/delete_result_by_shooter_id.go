package usecase

import (
	"context"
	"net/http"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
)

func (usecase *adminSuperUsecaseImpl) DeleteResultByShooterId(req model.ByShooterIdRequest) model.WebServiceResponse {
	if err := usecase.Store.DeleteResultByShooterId(context.Background(), req.ShooterID); err != nil {
		return util.ToWebServiceResponse("Gagal menghapus hasil ujian: "+err.Error(), http.StatusInternalServerError, nil)
	}

	return util.ToWebServiceResponse("Berhasil menghapus hasil ujian", http.StatusOK, nil)
}
