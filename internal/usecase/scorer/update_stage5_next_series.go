package usecase

import (
	"context"
	"net/http"
	"strconv"

	repositoryModel "github.com/FadhilAF/perbakin-be/internal/repository/postgres/sqlc"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
)

func (usecase *scorerUsecaseImpl) UpdateStage5NextNo(req model.ByIdAndTryRequest) model.WebServiceResponse {
	var newStatus string

	switch req.Try {
	case "1":
		status, _ := usecase.Store.GetStage5try1Status(context.Background(), req.ID)
		if status == model.Stage5EndStatus {
			return util.ToWebServiceResponse("Tidak dapat melanjutkan, batas no stage 5 percobaan 1 hanya 3", http.StatusForbidden, nil)
		}

		intStatus, _ := strconv.Atoi(string(status))
		newStatus = strconv.Itoa(intStatus + 1)

		err := usecase.Store.UpdateStage5try1NextNo(context.Background(), repositoryModel.UpdateStage5try1NextNoParams{
			ID:     req.ID,
			Status: repositoryModel.Stage5Status(newStatus),
		})
		if err != nil {
			return util.ToWebServiceResponse("Gagal melanjutkan no stage 5 percobaan 1: "+err.Error(), http.StatusInternalServerError, nil)
		}
	case "2":
		status, _ := usecase.Store.GetStage5try2Status(context.Background(), req.ID)
		if status == model.Stage5EndStatus {
			return util.ToWebServiceResponse("Tidak dapat melanjutkan, batas no stage 5 percobaan 2 hanya 3", http.StatusForbidden, nil)
		}

		intStatus, _ := strconv.Atoi(string(status))
		newStatus = strconv.Itoa(intStatus + 1)

		err := usecase.Store.UpdateStage5try2NextNo(context.Background(), repositoryModel.UpdateStage5try2NextNoParams{
			ID:     req.ID,
			Status: repositoryModel.Stage5Status(newStatus),
		})
		if err != nil {
			return util.ToWebServiceResponse("Gagal melanjutkan no stage 5 percobaan 2: "+err.Error(), http.StatusInternalServerError, nil)
		}
	}

	return util.ToWebServiceResponse("Berhasil melanjutkan no stage 5 percobaan "+req.Try+" ke no "+newStatus, http.StatusOK, nil)
}
