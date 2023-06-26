package usecase

import (
	"context"
	"net/http"
	"strconv"

	repositoryModel "github.com/FadhilAF/perbakin-be/internal/repository/postgres/sqlc"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
)

func (usecase *scorerUsecaseImpl) UpdateStage4NextNo(req model.ByIdAndTryRequest) model.WebServiceResponse {
	var newStatus string

	switch req.Try {
	case "1":
		status, _ := usecase.Store.GetStage4try1Status(context.Background(), req.ID)
		if status == model.Stage246EndStatus {
			return util.ToWebServiceResponse("Tidak dapat melanjutkan, batas no stage 4 percobaan 1 hanya 3", http.StatusForbidden, nil)
		}

		intStatus, _ := strconv.Atoi(string(status))
		newStatus = strconv.Itoa(intStatus + 1)

		err := usecase.Store.UpdateStage4try1NextNo(context.Background(), repositoryModel.UpdateStage4try1NextNoParams{
			ID:     req.ID,
			Status: repositoryModel.Stage246Status(newStatus),
		})
		if err != nil {
			return util.ToWebServiceResponse("Gagal melanjutkan no stage 4 percobaan 1: "+err.Error(), http.StatusInternalServerError, nil)
		}
	case "2":
		status, _ := usecase.Store.GetStage4try2Status(context.Background(), req.ID)
		if status == model.Stage246EndStatus {
			return util.ToWebServiceResponse("Tidak dapat melanjutkan, batas no stage 4 percobaan 2 hanya 3", http.StatusForbidden, nil)
		}

		intStatus, _ := strconv.Atoi(string(status))
		newStatus = strconv.Itoa(intStatus + 1)

		err := usecase.Store.UpdateStage4try2NextNo(context.Background(), repositoryModel.UpdateStage4try2NextNoParams{
			ID:     req.ID,
			Status: repositoryModel.Stage246Status(newStatus),
		})
		if err != nil {
			return util.ToWebServiceResponse("Gagal melanjutkan no stage 4 percobaan 2: "+err.Error(), http.StatusInternalServerError, nil)
		}
	}

	return util.ToWebServiceResponse("Berhasil melanjutkan no stage 4 percobaan "+req.Try+" ke no "+newStatus, http.StatusOK, nil)
}
