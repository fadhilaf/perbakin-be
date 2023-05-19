package usecase

import (
	"context"
	"net/http"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
)

func (usecase *adminSuperUsecaseImpl) DeleteScorer(req model.UserByUserIdRequest) model.WebServiceResponse {
	if err := usecase.Store.DeleteUser(context.Background(), req.UserID); err != nil {
		return util.ToWebServiceResponse("Gagal menghapus penguji: "+err.Error(), http.StatusInternalServerError, nil)
	}

	return util.ToWebServiceResponse("Berhasil menghapus penguji", http.StatusOK, nil)
}
