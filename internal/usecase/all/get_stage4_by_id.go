package usecase

import (
	"context"
	"net/http"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (usecase *allUsecaseImpl) GetStage4ById(req model.ByIdRequest) model.WebServiceResponse {
	stage4, err := usecase.Store.GetStage4ById(context.Background(), req.ID)
	if err != nil {
		return util.ToWebServiceResponse("Hasil ujian stage 4 tidak ada: "+err.Error(), http.StatusNotFound, nil)
	}

	var stage4Response gin.H

	if stage4.IsTry2 {
		stage4Response = gin.H{
			"stage_4": model.Stage46Full{
				ID:       stage4.ID,
				ResultID: stage4.ResultID,
				Try1: model.Stage46Try{
					Status:     string(stage4.Try1Status),
					No1:        util.Stage46DatabaseNumbersToStruct(stage4.Try1No1),
					No2:        util.Stage46DatabaseNumbersToStruct(stage4.Try1No2),
					No3:        util.Stage46DatabaseNumbersToStruct(stage4.Try1No3),
					Checkmarks: util.CheckmarksToBoolArray(stage4.Try1Checkmarks),
				},
				Try2: model.Stage46Try{
					Status:     string(stage4.Try2Status.Stage246Status),
					No1:        util.Stage46DatabaseNumbersToStruct(stage4.Try2No1.String),
					No2:        util.Stage46DatabaseNumbersToStruct(stage4.Try2No2.String),
					No3:        util.Stage46DatabaseNumbersToStruct(stage4.Try2No3.String),
					Checkmarks: util.CheckmarksToBoolArray(stage4.Try2Checkmarks.String),
				},
				IsTry2:      stage4.IsTry2,
				ShooterSign: stage4.ShooterSign,
				ScorerSign:  stage4.ScorerSign,
				CreatedAt:   stage4.CreatedAt.Time,
				UpdatedAt:   stage4.UpdatedAt.Time,
			}}
	} else {
		stage4Response = gin.H{
			"stage_4": model.Stage46{
				ID:       stage4.ID,
				ResultID: stage4.ResultID,
				Try1: model.Stage46Try{
					Status:     string(stage4.Try1Status),
					No1:        util.Stage46DatabaseNumbersToStruct(stage4.Try1No1),
					No2:        util.Stage46DatabaseNumbersToStruct(stage4.Try1No2),
					No3:        util.Stage46DatabaseNumbersToStruct(stage4.Try1No3),
					Checkmarks: util.CheckmarksToBoolArray(stage4.Try1Checkmarks),
				},
				IsTry2:      stage4.IsTry2,
				ShooterSign: stage4.ShooterSign,
				ScorerSign:  stage4.ScorerSign,
				CreatedAt:   stage4.CreatedAt.Time,
				UpdatedAt:   stage4.UpdatedAt.Time,
			}}
	}

	return util.ToWebServiceResponse("Berhasil mendapatkan hasil ujian stage 4", http.StatusOK, stage4Response)
}
