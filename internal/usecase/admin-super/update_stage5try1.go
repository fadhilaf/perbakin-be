package usecase

import (
	"context"
	"net/http"

	repositoryModel "github.com/FadhilAF/perbakin-be/internal/repository/postgres/sqlc"
	"github.com/gin-gonic/gin"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
)

func (usecase *adminSuperUsecaseImpl) UpdateStage5try1(req model.UpdateStage5try1Request) model.WebServiceResponse {
	newStage5, err := usecase.Store.UpdateStage5try1(context.Background(), repositoryModel.UpdateStage5try1Params{
		ID:             req.ID,
		Try1Status:     repositoryModel.Stage5Status(req.Try1.Status),
		Try1No1:        req.Try1.No1,
		Try1No2:        req.Try1.No2,
		Try1Checkmarks: req.Try1.Checkmarks,
	})
	if err != nil {
		return util.ToWebServiceResponse("Gagal mengubah hasil ujian stage 5: "+err.Error(), http.StatusInternalServerError, nil)
	}

	return util.ToWebServiceResponse("Berhasil mengubah hasil ujian stage 5", http.StatusOK, gin.H{
		"stage_5": model.UpdateStage5try1Response{
			Try1: model.Stage5Try{
				Status:     string(newStage5.Try1Status),
				No1:        util.Stage5DatabaseNumbersToStruct(newStage5.Try1No1),
				No2:        util.Stage5DatabaseNumbersToStruct(newStage5.Try1No2),
				Checkmarks: util.CheckmarksToBoolArray(newStage5.Try1Checkmarks),
			},
			UpdatedAt: newStage5.UpdatedAt.Time,
		}})
}
