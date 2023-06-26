package usecase

import (
	"context"
	"net/http"

	repositoryModel "github.com/FadhilAF/perbakin-be/internal/repository/postgres/sqlc"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
)

func (usecase *scorerUsecaseImpl) UpdateStage0Finish(req model.UpdateStageFinishRequest) model.WebServiceResponse {
	status, _ := usecase.Store.GetStage0Status(context.Background(), req.ID)
	if status != model.Stage0EndStatus {
		return util.ToWebServiceResponse("Tidak dapat menyelesaikan babak kualifikasi, masih pada seri ke-"+string(status), http.StatusForbidden, nil)
	}

	if req.Success {
		if err := usecase.Store.UpdateStage0FinishSuccess(context.Background(), repositoryModel.UpdateStage0FinishSuccessParams{
			ID:          req.ID,
			ShooterSign: req.ShooterSign,
			ScorerSign:  req.ScorerSign,
		}); err != nil {
			return util.ToWebServiceResponse("Gagal menyelesaikan babak kualifikasi menjadi berhasil: "+err.Error(), http.StatusInternalServerError, nil)
		}
	} else {
		if err := usecase.Store.UpdateStage0FinishFailed(context.Background(), repositoryModel.UpdateStage0FinishFailedParams{
			ID:          req.ID,
			ShooterSign: req.ShooterSign,
			ScorerSign:  req.ScorerSign,
		}); err != nil {
			return util.ToWebServiceResponse("Gagal menyelesaikan babak kualifikasi menjadi gagal: "+err.Error(), http.StatusInternalServerError, nil)
		}
	}

	var hasil string
	if req.Success {
		hasil = "berhasil"
	} else {
		hasil = "gagal"
	}

	return util.ToWebServiceResponse("Berhasil menyelesaikan pendataan kualifikasi menjadi "+hasil, http.StatusOK, nil)
}
