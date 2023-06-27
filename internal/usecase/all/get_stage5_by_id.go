package usecase

import (
	"context"
	"net/http"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (usecase *allUsecaseImpl) GetStage5ById(req model.ByIdRequest) model.WebServiceResponse {
	stage5, err := usecase.Store.GetStage5ById(context.Background(), req.ID)
	if err != nil {
		return util.ToWebServiceResponse("Hasil ujian stage 5 tidak ada: "+err.Error(), http.StatusNotFound, nil)
	}

	var stage5Response gin.H

	if stage5.IsTry2 {
		stage5Response = gin.H{
			"stage_5": model.Stage5Full{
				ID:       stage5.ID,
				ResultID: stage5.ResultID,
				Try1: model.Stage5Try{
					Status:     string(stage5.Try1Status),
					No1:        util.Stage5DatabaseNumbersToStruct(stage5.Try1No1),
					No2:        util.Stage5DatabaseNumbersToStruct(stage5.Try1No2),
					Checkmarks: util.CheckmarksToBoolArray(stage5.Try1Checkmarks),
				},
				Try2: model.Stage5Try{
					Status:     string(stage5.Try2Status.Stage5Status),
					No1:        util.Stage5DatabaseNumbersToStruct(stage5.Try2No1.String),
					No2:        util.Stage5DatabaseNumbersToStruct(stage5.Try2No2.String),
					Checkmarks: util.CheckmarksToBoolArray(stage5.Try2Checkmarks.String),
				},
				IsTry2:      stage5.IsTry2,
				ShooterSign: stage5.ShooterSign,
				ScorerSign:  stage5.ScorerSign,
				CreatedAt:   stage5.CreatedAt.Time,
				UpdatedAt:   stage5.UpdatedAt.Time,
			}}
	} else {
		stage5Response = gin.H{
			"stage_5": model.Stage5{
				ID:       stage5.ID,
				ResultID: stage5.ResultID,
				Try1: model.Stage5Try{
					Status:     string(stage5.Try1Status),
					No1:        util.Stage5DatabaseNumbersToStruct(stage5.Try1No1),
					No2:        util.Stage5DatabaseNumbersToStruct(stage5.Try1No2),
					Checkmarks: util.CheckmarksToBoolArray(stage5.Try1Checkmarks),
				},
				IsTry2:      stage5.IsTry2,
				ShooterSign: stage5.ShooterSign,
				ScorerSign:  stage5.ScorerSign,
				CreatedAt:   stage5.CreatedAt.Time,
				UpdatedAt:   stage5.UpdatedAt.Time,
			}}
	}

	return util.ToWebServiceResponse("Berhasil mendapatkan hasil ujian stage 5", http.StatusOK, stage5Response)
}
