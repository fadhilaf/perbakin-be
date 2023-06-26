package usecase

import (
	"context"
	"net/http"

	repositoryModel "github.com/FadhilAF/perbakin-be/internal/repository/postgres/sqlc"
	"github.com/gin-gonic/gin"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
)

func (usecase *adminSuperUsecaseImpl) UpdateStage2try1(req model.UpdateStage246try1Request) model.WebServiceResponse {
	newStage2, err := usecase.Store.UpdateStage2try1(context.Background(), repositoryModel.UpdateStage2try1Params{
		ID:             req.ID,
		Try1Status:     repositoryModel.Stage246Status(req.Try1.Status),
		Try1No1:        req.Try1.No1,
		Try1No2:        req.Try1.No2,
		Try1No3:        req.Try1.No3,
		Try1Checkmarks: req.Try1.Checkmarks,
	})
	if err != nil {
		return util.ToWebServiceResponse("Gagal mengubah hasil ujian stage 2: "+err.Error(), http.StatusInternalServerError, nil)
	}

	return util.ToWebServiceResponse("Berhasil mengubah hasil ujian stage 2", http.StatusOK, gin.H{
		"stage_2": model.UpdateStage2try1Response{
			Try1: model.Stage2Try{
				Status:     string(newStage2.Try1Status),
				No1:        util.Stage123DatabaseNumbersToStruct(newStage2.Try1No1),
				No2:        util.Stage123DatabaseNumbersToStruct(newStage2.Try1No2),
				No3:        util.Stage123DatabaseNumbersToStruct(newStage2.Try1No3),
				Checkmarks: util.CheckmarksToBoolArray(newStage2.Try1Checkmarks),
			},
			UpdatedAt: newStage2.UpdatedAt.Time,
		}})
}
