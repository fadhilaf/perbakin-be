package usecase

import (
	"context"
	"net/http"

	repositoryModel "github.com/FadhilAF/perbakin-be/internal/repository/postgres/sqlc"
	"github.com/gin-gonic/gin"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
)

func (usecase *adminSuperUsecaseImpl) UpdateStage6try2(req model.UpdateStage246try2Request) model.WebServiceResponse {
	newStage6, err := usecase.Store.UpdateStage6try2(context.Background(), repositoryModel.UpdateStage6try2Params{
		ID:             req.ID,
		Try1Status:     repositoryModel.Stage246Status(req.Try1.Status),
		Try1No1:        req.Try1.No1,
		Try1No2:        req.Try1.No2,
		Try1No3:        req.Try1.No3,
		Try1Checkmarks: req.Try1.Checkmarks,
		Try2Status:     repositoryModel.Stage246Status(req.Try2.Status),
		Try2No1:        req.Try2.No1,
		Try2No2:        req.Try2.No2,
		Try2No3:        req.Try2.No3,
		Try2Checkmarks: req.Try2.Checkmarks,
	})
	if err != nil {
		return util.ToWebServiceResponse("Gagal mengubah hasil ujian stage 6: "+err.Error(), http.StatusInternalServerError, nil)
	}

	return util.ToWebServiceResponse("Berhasil mengubah hasil ujian stage 6", http.StatusOK, gin.H{
		"stage_6": model.UpdateStage46try2Response{
			Try1: model.Stage46Try{
				Status:     string(newStage6.Try1Status),
				No1:        util.Stage46DatabaseNumbersToStruct(newStage6.Try1No1),
				No2:        util.Stage46DatabaseNumbersToStruct(newStage6.Try1No2),
				No3:        util.Stage46DatabaseNumbersToStruct(newStage6.Try1No3),
				Checkmarks: util.CheckmarksToBoolArray(newStage6.Try1Checkmarks),
			},
			Try2: model.Stage46Try{
				Status:     string(newStage6.Try2Status),
				No1:        util.Stage46DatabaseNumbersToStruct(newStage6.Try2No1),
				No2:        util.Stage46DatabaseNumbersToStruct(newStage6.Try2No2),
				No3:        util.Stage46DatabaseNumbersToStruct(newStage6.Try2No3),
				Checkmarks: util.CheckmarksToBoolArray(newStage6.Try2Checkmarks),
			},
			UpdatedAt: newStage6.UpdatedAt.Time,
		}})
}
