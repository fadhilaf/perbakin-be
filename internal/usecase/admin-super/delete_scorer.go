package usecase

import (
	"context"
	"net/http"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
)

func (usecase *adminSuperUsecaseImpl) DeleteScorer(req model.DeleteUserRequest) model.WebServiceResponse {
	scorer, err := usecase.Store.GetScorer(context.Background(), req.ID)
	if err != nil {
		return util.ToWebServiceResponse("Tidak ditemukan scorer dengan ID tersebut", http.StatusNotFound, nil)
	}

	if err = usecase.Store.DeleteUser(context.Background(), scorer.UserID); err != nil {
		return util.ToWebServiceResponse("Gagal menghapus scorer", http.StatusInternalServerError, nil)
	}

	return util.ToWebServiceResponse("Berhasil menghapus scorer", http.StatusOK, nil)
}
