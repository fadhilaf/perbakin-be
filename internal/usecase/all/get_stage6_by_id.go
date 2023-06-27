package usecase

import (
	"context"
	"net/http"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (usecase *allUsecaseImpl) GetStage6ById(req model.ByIdRequest) model.WebServiceResponse {
	stage6, err := usecase.Store.GetStage6ById(context.Background(), req.ID)
	if err != nil {
		return util.ToWebServiceResponse("Hasil ujian stage 6 tidak ada: "+err.Error(), http.StatusNotFound, nil)
	}

	var stage6Response gin.H

	if stage6.IsTry2 {
		stage6Response = gin.H{
			"stage_6": model.Stage46Full{
				ID:       stage6.ID,
				ResultID: stage6.ResultID,
				Try1: model.Stage46Try{
					Status:     string(stage6.Try1Status),
					No1:        util.Stage46DatabaseNumbersToStruct(stage6.Try1No1),
					No2:        util.Stage46DatabaseNumbersToStruct(stage6.Try1No2),
					No3:        util.Stage46DatabaseNumbersToStruct(stage6.Try1No3),
					Checkmarks: util.CheckmarksToBoolArray(stage6.Try1Checkmarks),
				},
				Try2: model.Stage46Try{
					Status:     string(stage6.Try2Status.Stage246Status),
					No1:        util.Stage46DatabaseNumbersToStruct(stage6.Try2No1.String),
					No2:        util.Stage46DatabaseNumbersToStruct(stage6.Try2No2.String),
					No3:        util.Stage46DatabaseNumbersToStruct(stage6.Try2No3.String),
					Checkmarks: util.CheckmarksToBoolArray(stage6.Try2Checkmarks.String),
				},
				IsTry2:      stage6.IsTry2,
				ShooterSign: stage6.ShooterSign,
				ScorerSign:  stage6.ScorerSign,
				CreatedAt:   stage6.CreatedAt.Time,
				UpdatedAt:   stage6.UpdatedAt.Time,
			}}
	} else {
		stage6Response = gin.H{
			"stage_6": model.Stage46{
				ID:       stage6.ID,
				ResultID: stage6.ResultID,
				Try1: model.Stage46Try{
					Status:     string(stage6.Try1Status),
					No1:        util.Stage46DatabaseNumbersToStruct(stage6.Try1No1),
					No2:        util.Stage46DatabaseNumbersToStruct(stage6.Try1No2),
					No3:        util.Stage46DatabaseNumbersToStruct(stage6.Try1No3),
					Checkmarks: util.CheckmarksToBoolArray(stage6.Try1Checkmarks),
				},
				IsTry2:      stage6.IsTry2,
				ShooterSign: stage6.ShooterSign,
				ScorerSign:  stage6.ScorerSign,
				CreatedAt:   stage6.CreatedAt.Time,
				UpdatedAt:   stage6.UpdatedAt.Time,
			}}
	}

	return util.ToWebServiceResponse("Berhasil mendapatkan hasil ujian stage 6", http.StatusOK, stage6Response)
}
