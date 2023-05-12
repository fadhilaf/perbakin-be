package usecase

import (
	"context"
	"net/http"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
)

func (usecase *adminSuperUsecaseImpl) DeleteShooter(req model.ShooterByIdRequest) model.WebServiceResponse {
	shooter, err := usecase.Store.GetShooterById(context.Background(), req.ID)
	if err != nil {
		return util.ToWebServiceResponse("Tidak ditemukan penembak dengan ID tersebut", http.StatusNotFound, nil)
	}

	if shooter.ScorerID != req.ScorerID {
		return util.ToWebServiceResponse("Tidak diperbohkan menghapus penembak ujian lain", http.StatusUnauthorized, nil)
	}

	if err = usecase.Store.DeleteShooter(context.Background(), shooter.ID); err != nil {
		return util.ToWebServiceResponse("Gagal menghapus penembak: "+err.Error(), http.StatusInternalServerError, nil)
	}

	return util.ToWebServiceResponse("Berhasil menghapus penembak", http.StatusOK, nil)
}
