package usecase

import (
	"context"
	"net/http"

	repositoryModel "github.com/FadhilAF/perbakin-be/internal/repository/postgres/sqlc"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
)

func (usecase *scorerUsecaseImpl) UpdateStage4Finish(req model.UpdateStage123456FinishRequest) model.WebServiceResponse {
	switch req.Try {
	case "1":
		status, _ := usecase.Store.GetStage4try1Status(context.Background(), req.ID)
		if status != model.Stage246EndStatus {
			return util.ToWebServiceResponse("Tidak dapat menyelesaikan stage 4 percobaan 1, masih pada nomor ke-"+string(status), http.StatusForbidden, nil)
		}

		if req.Success {
			if err := usecase.Store.UpdateStage4try1FinishSuccess(context.Background(), repositoryModel.UpdateStage4try1FinishSuccessParams{
				ID:          req.ID,
				ShooterSign: req.ShooterSign,
				ScorerSign:  req.ScorerSign,
			}); err != nil {
				return util.ToWebServiceResponse("Gagal menyelesaikan stage 4 percobaan 1 menjadi berhasil: "+err.Error(), http.StatusInternalServerError, nil)
			}
		} else {
			if err := usecase.Store.UpdateStage4try1FinishFailed(context.Background(), repositoryModel.UpdateStage4try1FinishFailedParams{
				ID:          req.ID,
				ShooterSign: req.ShooterSign,
				ScorerSign:  req.ScorerSign,
			}); err != nil {
				return util.ToWebServiceResponse("Gagal menyelesaikan stage 4 percobaan 1 menjadi gagal: "+err.Error(), http.StatusInternalServerError, nil)
			}
		}
	case "2":
		status, _ := usecase.Store.GetStage4try2Status(context.Background(), req.ID)
		if status != model.Stage246EndStatus {
			return util.ToWebServiceResponse("Tidak dapat menyelesaikan stage 4 percobaan 2, masih pada nomor ke-"+string(status), http.StatusForbidden, nil)
		}

		if req.Success {
			if err := usecase.Store.UpdateStage4try2FinishSuccess(context.Background(), repositoryModel.UpdateStage4try2FinishSuccessParams{
				ID:          req.ID,
				ShooterSign: req.ShooterSign,
				ScorerSign:  req.ScorerSign,
			}); err != nil {
				return util.ToWebServiceResponse("Gagal menyelesaikan stage 4 percobaan 2 menjadi berhasil: "+err.Error(), http.StatusInternalServerError, nil)
			}
		} else {
			if err := usecase.Store.UpdateStage4try2FinishFailed(context.Background(), repositoryModel.UpdateStage4try2FinishFailedParams{
				ID:          req.ID,
				ShooterSign: req.ShooterSign,
				ScorerSign:  req.ScorerSign,
			}); err != nil {
				return util.ToWebServiceResponse("Gagal menyelesaikan stage 4 percobaan 2 menjadi gagal: "+err.Error(), http.StatusInternalServerError, nil)
			}
		}
	}

	var hasil string
	if req.Success {
		hasil = "berhasil"
	} else {
		hasil = "gagal"
	}

	return util.ToWebServiceResponse("Berhasil menyelesaikan pendataan stage 4 percobaan "+req.Try+" menjadi "+hasil, http.StatusOK, nil)
}
