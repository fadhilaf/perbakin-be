package usecase

import (
	"context"
	"net/http"

	repositoryModel "github.com/FadhilAF/perbakin-be/internal/repository/postgres/sqlc"
	"github.com/gin-gonic/gin"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
)

func (usecase *adminSuperUsecaseImpl) UpdateStage1try2(req model.UpdateStage13try2Request) model.WebServiceResponse {
	newStage1, err := usecase.Store.UpdateStage1try2(context.Background(), repositoryModel.UpdateStage1try2Params{
		ID:             req.ID,
		Try1Status:     repositoryModel.Stage13Status(req.Try1.Status),
		Try1No1:        req.Try1.No1,
		Try1No2:        req.Try1.No2,
		Try1No3:        req.Try1.No3,
		Try1No4:        req.Try1.No4,
		Try1No5:        req.Try1.No5,
		Try1No6:        req.Try1.No6,
		Try1Checkmarks: req.Try1.Checkmarks,
		Try2Status:     repositoryModel.Stage13Status(req.Try2.Status),
		Try2No1:        req.Try2.No1,
		Try2No2:        req.Try2.No2,
		Try2No3:        req.Try2.No3,
		Try2No4:        req.Try2.No4,
		Try2No5:        req.Try2.No5,
		Try2No6:        req.Try2.No6,
		Try2Checkmarks: req.Try2.Checkmarks,
	})
	if err != nil {
		return util.ToWebServiceResponse("Gagal mengubah hasil ujian stage 1: "+err.Error(), http.StatusInternalServerError, nil)
	}

	return util.ToWebServiceResponse("Berhasil mengubah hasil ujian stage 1", http.StatusOK, gin.H{
		"stage_1": model.UpdateStage13try2Response{
			Try1: model.Stage13Try{
				Status:     string(newStage1.Try1Status),
				No1:        util.Stage123DatabaseNumbersToStruct(newStage1.Try1No1),
				No2:        util.Stage123DatabaseNumbersToStruct(newStage1.Try1No2),
				No3:        util.Stage123DatabaseNumbersToStruct(newStage1.Try1No3),
				No4:        util.Stage123DatabaseNumbersToStruct(newStage1.Try1No4),
				No5:        util.Stage123DatabaseNumbersToStruct(newStage1.Try1No5),
				No6:        util.Stage123DatabaseNumbersToStruct(newStage1.Try1No6),
				Checkmarks: util.CheckmarksToBoolArray(newStage1.Try1Checkmarks),
			},
			Try2: model.Stage13Try{
				Status:     string(newStage1.Try2Status),
				No1:        util.Stage123DatabaseNumbersToStruct(newStage1.Try2No1),
				No2:        util.Stage123DatabaseNumbersToStruct(newStage1.Try2No2),
				No3:        util.Stage123DatabaseNumbersToStruct(newStage1.Try2No3),
				No4:        util.Stage123DatabaseNumbersToStruct(newStage1.Try2No4),
				No5:        util.Stage123DatabaseNumbersToStruct(newStage1.Try2No5),
				No6:        util.Stage123DatabaseNumbersToStruct(newStage1.Try2No6),
				Checkmarks: util.CheckmarksToBoolArray(newStage1.Try2Checkmarks),
			},
			UpdatedAt: newStage1.UpdatedAt.Time,
		}})
}
