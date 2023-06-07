package usecase

import (
	"context"
	"net/http"

	repositoryModel "github.com/FadhilAF/perbakin-be/internal/repository/postgres/sqlc"
	"github.com/gin-gonic/gin"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
)

func (usecase *scorerUsecaseImpl) UpdateStage1Checkmarks(req model.UpdateStage123456CheckmarksRequest) model.WebServiceResponse {
	var checkmarks string
	var err error

	switch req.Try {
	case "1":
		if status, _ := usecase.Store.GetStage1try1Status(context.Background(), req.ID); status == "7" {
			return util.ToWebServiceResponse("Gagal mengupdate stage 1 percobaan 1 centang hasil, pencatatan data stage 1 percobaan 1 sudah selesai", http.StatusForbidden, nil)
		}

		checkmarks, err = usecase.Store.UpdateStage1try1Checkmarks(context.Background(), repositoryModel.UpdateStage1try1CheckmarksParams{
			ID:         req.ID,
			Checkmarks: req.Checkmarks,
		})
		if err != nil {
			return util.ToWebServiceResponse("Gagal mengupdate stage 1 percobaan 1 centang hasil: "+err.Error(), http.StatusInternalServerError, nil)
		}
	case "2":
		if status, _ := usecase.Store.GetStage1try2Status(context.Background(), req.ID); status == "7" {
			return util.ToWebServiceResponse("Gagal mengupdate stage 1 percobaan 2 centang hasil, pencatatan data stage 1 percobaan 2 sudah selesai", http.StatusForbidden, nil)
		}

		checkmarks, err = usecase.Store.UpdateStage1try2Checkmarks(context.Background(), repositoryModel.UpdateStage1try2CheckmarksParams{
			ID:         req.ID,
			Checkmarks: req.Checkmarks,
		})
		if err != nil {
			return util.ToWebServiceResponse("Gagal mengupdate stage 1 percobaan 2 centang hasil: "+err.Error(), http.StatusInternalServerError, nil)
		}
	}

	return util.ToWebServiceResponse("Berhasil mengupdate stage 1 percobaan "+req.Try+" centang hasil", http.StatusOK, gin.H{
		"scores": util.CheckmarksToBoolArray(checkmarks),
	})
}
