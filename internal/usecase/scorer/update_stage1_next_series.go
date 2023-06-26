package usecase

import (
	"context"
	"net/http"
	"strconv"

	repositoryModel "github.com/FadhilAF/perbakin-be/internal/repository/postgres/sqlc"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
)

func (usecase *scorerUsecaseImpl) UpdateStage1NextNo(req model.ByIdAndTryRequest) model.WebServiceResponse {
	var newStatus string

	switch req.Try {
	case "1":
		status, _ := usecase.Store.GetStage1try1Status(context.Background(), req.ID)
		if status == model.Stage13EndStatus {
			return util.ToWebServiceResponse("Tidak dapat melanjutkan, batas no stage 1 percobaan 1 hanya 6", http.StatusForbidden, nil)
		}

		intStatus, _ := strconv.Atoi(string(status))
		newStatus = strconv.Itoa(intStatus + 1)

		err := usecase.Store.UpdateStage1try1NextNo(context.Background(), repositoryModel.UpdateStage1try1NextNoParams{
			ID:     req.ID,
			Status: repositoryModel.Stage13Status(newStatus),
		})
		if err != nil {
			return util.ToWebServiceResponse("Gagal melanjutkan no stage 1 percobaan 1: "+err.Error(), http.StatusInternalServerError, nil)
		}
	case "2":
		status, _ := usecase.Store.GetStage1try2Status(context.Background(), req.ID)
		if status == model.Stage13EndStatus {
			return util.ToWebServiceResponse("Tidak dapat melanjutkan, batas no stage 1 percobaan 2 hanya 6", http.StatusForbidden, nil)
		}

		intStatus, _ := strconv.Atoi(string(status))
		newStatus = strconv.Itoa(intStatus + 1)

		err := usecase.Store.UpdateStage1try2NextNo(context.Background(), repositoryModel.UpdateStage1try2NextNoParams{
			ID:     req.ID,
			Status: repositoryModel.Stage13Status(newStatus),
		})
		if err != nil {
			return util.ToWebServiceResponse("Gagal melanjutkan no stage 1 percobaan 2: "+err.Error(), http.StatusInternalServerError, nil)
		}
	}

	return util.ToWebServiceResponse("Berhasil melanjutkan no stage 1 percobaan "+req.Try+" ke no "+newStatus, http.StatusOK, nil)
}
