package usecase

import (
	"context"
	"net/http"

	repositoryModel "github.com/FadhilAF/perbakin-be/internal/repository/postgres/sqlc"
	"github.com/gin-gonic/gin"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
)

func (usecase *scorerUsecaseImpl) UpdateStage0Checkmarks(req model.UpdateStage0CheckmarksRequest) model.WebServiceResponse {
	if status, _ := usecase.Store.GetStage0Status(context.Background(), req.ID); status == "6" {
		return util.ToWebServiceResponse("Gagal mengupdate kualifikasi centang hasil, pencatatan data kualifikasi sudah selesai", http.StatusForbidden, nil)
	}

	checkmarks, err := usecase.Store.UpdateStage0Checkmarks(context.Background(), repositoryModel.UpdateStage0CheckmarksParams{
		ID:         req.ID,
		Checkmarks: req.Checkmarks,
	})
	if err != nil {
		return util.ToWebServiceResponse("Gagal mengupdate kualifikasi centang hasil: "+err.Error(), http.StatusInternalServerError, nil)
	}

	return util.ToWebServiceResponse("Berhasil mengupdate kualifikasi centang hasil", http.StatusOK, gin.H{
		"scores": util.CheckmarksToBoolArray(checkmarks),
	})
}
