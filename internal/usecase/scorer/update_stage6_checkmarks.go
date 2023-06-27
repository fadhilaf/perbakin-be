package usecase

import (
	"context"
	"net/http"

	repositoryModel "github.com/FadhilAF/perbakin-be/internal/repository/postgres/sqlc"
	"github.com/gin-gonic/gin"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
)

func (usecase *scorerUsecaseImpl) UpdateStage6Checkmarks(req model.UpdateStage123456CheckmarksRequest) model.WebServiceResponse {
	var checkmarks string
	var err error

	switch req.Try {
	case "1":
		if status, _ := usecase.Store.GetStage6try1Status(context.Background(), req.ID); status == model.Stage246FinishedStatus {
			return util.ToWebServiceResponse("Gagal mengupdate centang hasil stage 6 percobaan 1, pencatatan data stage 6 percobaan 1 sudah selesai", http.StatusForbidden, nil)
		}

		checkmarks, err = usecase.Store.UpdateStage6try1Checkmarks(context.Background(), repositoryModel.UpdateStage6try1CheckmarksParams{
			ID:         req.ID,
			Checkmarks: req.Checkmarks,
		})
		if err != nil {
			return util.ToWebServiceResponse("Gagal mengupdate centang hasil stage 6 percobaan 1: "+err.Error(), http.StatusInternalServerError, nil)
		}
	case "2":
		if status, _ := usecase.Store.GetStage6try2Status(context.Background(), req.ID); status == model.Stage246FinishedStatus {
			return util.ToWebServiceResponse("Gagal mengupdate centang hasil stage 6 percobaan 2, pencatatan data stage 6 percobaan 2 sudah selesai", http.StatusForbidden, nil)
		}

		checkmarks, err = usecase.Store.UpdateStage6try2Checkmarks(context.Background(), repositoryModel.UpdateStage6try2CheckmarksParams{
			ID:         req.ID,
			Checkmarks: req.Checkmarks,
		})
		if err != nil {
			return util.ToWebServiceResponse("Gagal mengupdate centang hasil stage 6 percobaan 2: "+err.Error(), http.StatusInternalServerError, nil)
		}
	}

	return util.ToWebServiceResponse("Berhasil mengupdate centang hasil stage 6 percobaan "+req.Try, http.StatusOK, gin.H{
		"checkmarks": util.CheckmarksToBoolArray(checkmarks),
	})
}
