package usecase

import (
	"context"
	"net/http"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
)

func (usecase *adminSuperUsecaseImpl) DeleteScorer(req model.OperatorByIdRequest) model.WebServiceResponse {
	scorer, err := usecase.Store.GetScorerById(context.Background(), req.ID)
	if err != nil {
		return util.ToWebServiceResponse("Tidak ditemukan penguji dengan ID tersebut", http.StatusNotFound, nil)
	}

	if scorer.ExamID != req.ExamID {
		return util.ToWebServiceResponse("Tidak diperbohkan menghapus penguji ujian lain", http.StatusUnauthorized, nil)
	}

	if err = usecase.Store.DeleteUser(context.Background(), scorer.UserID); err != nil {
		return util.ToWebServiceResponse("Gagal menghapus penguji: "+err.Error(), http.StatusInternalServerError, nil)
	}

	return util.ToWebServiceResponse("Berhasil menghapus penguji", http.StatusOK, nil)
}
