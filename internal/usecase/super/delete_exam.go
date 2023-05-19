package usecase

import (
	"context"
	"net/http"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
)

func (usecase *superUsecaseImpl) DeleteExam(req model.ByIdRequest) model.WebServiceResponse {
	if err := usecase.Store.DeleteExam(context.Background(), req.ID); err != nil {
		return util.ToWebServiceResponse("Gagal menghapus ujian: "+err.Error(), http.StatusInternalServerError, nil)
	}

	return util.ToWebServiceResponse("Berhasil menghapus ujian", http.StatusOK, nil)
}
