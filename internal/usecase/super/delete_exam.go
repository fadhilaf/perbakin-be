package usecase

import (
	"context"
	"net/http"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
)

func (usecase *superUsecaseImpl) DeleteExam(req model.DeleteExamRequest) model.WebServiceResponse {
	exam, err := usecase.Store.GetExamById(context.Background(), req.ID)
	if err != nil {
		return util.ToWebServiceResponse("Tidak ditemukan ujian dengan ID tersebut", http.StatusNotFound, nil)
	}

	if exam.SuperID != req.SuperID {
		return util.ToWebServiceResponse("Tidak dapat menghapus ujian super admin lain", http.StatusUnauthorized, nil)
	}

	if err = usecase.Store.DeleteExam(context.Background(), exam.ID); err != nil {
		return util.ToWebServiceResponse("Gagal menghapus ujian: "+err.Error(), http.StatusInternalServerError, nil)
	}

	return util.ToWebServiceResponse("Berhasil menghapus ujian", http.StatusOK, nil)
}
