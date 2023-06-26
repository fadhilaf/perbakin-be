package usecase

import (
	"context"
	"net/http"

	repositoryModel "github.com/FadhilAF/perbakin-be/internal/repository/postgres/sqlc"
	"github.com/gin-gonic/gin"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
)

func (usecase *adminSuperUsecaseImpl) UpdateStage3try1(req model.UpdateStage13try1Request) model.WebServiceResponse {
	newStage3, err := usecase.Store.UpdateStage3try1(context.Background(), repositoryModel.UpdateStage3try1Params{
		ID:             req.ID,
		Try1Status:     repositoryModel.Stage13Status(req.Try1.Status),
		Try1No1:        req.Try1.No1,
		Try1No2:        req.Try1.No2,
		Try1No3:        req.Try1.No3,
		Try1No4:        req.Try1.No4,
		Try1No5:        req.Try1.No5,
		Try1No6:        req.Try1.No6,
		Try1Checkmarks: req.Try1.Checkmarks,
	})
	if err != nil {
		return util.ToWebServiceResponse("Gagal mengubah hasil ujian stage 3: "+err.Error(), http.StatusInternalServerError, nil)
	}

	return util.ToWebServiceResponse("Berhasil mengubah hasil ujian stage 3", http.StatusOK, gin.H{
		"stage_3": model.UpdateStage13try1Response{
			Try1: model.Stage13Try{
				Status:     string(newStage3.Try1Status),
				No1:        util.Stage123DatabaseNumbersToStruct(newStage3.Try1No1),
				No2:        util.Stage123DatabaseNumbersToStruct(newStage3.Try1No2),
				No3:        util.Stage123DatabaseNumbersToStruct(newStage3.Try1No3),
				No4:        util.Stage123DatabaseNumbersToStruct(newStage3.Try1No4),
				No5:        util.Stage123DatabaseNumbersToStruct(newStage3.Try1No5),
				No6:        util.Stage123DatabaseNumbersToStruct(newStage3.Try1No6),
				Checkmarks: util.CheckmarksToBoolArray(newStage3.Try1Checkmarks),
			},
			UpdatedAt: newStage3.UpdatedAt.Time,
		}})
}
