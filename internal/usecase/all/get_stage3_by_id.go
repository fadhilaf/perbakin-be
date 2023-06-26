package usecase

import (
	"context"
	"net/http"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (usecase *allUsecaseImpl) GetStage3ById(req model.ByIdRequest) model.WebServiceResponse {
	stage3, err := usecase.Store.GetStage3ById(context.Background(), req.ID)
	if err != nil {
		return util.ToWebServiceResponse("Hasil ujian stage 3 tidak ada: "+err.Error(), http.StatusNotFound, nil)
	}

	var stage3Response gin.H

	if stage3.IsTry2 {
		stage3Response = gin.H{
			"stage_3": model.Stage13Full{
				ID:       stage3.ID,
				ResultID: stage3.ResultID,
				Try1: model.Stage13Try{
					Status:     string(stage3.Try1Status),
					No1:        util.Stage123DatabaseNumbersToStruct(stage3.Try1No1),
					No2:        util.Stage123DatabaseNumbersToStruct(stage3.Try1No2),
					No3:        util.Stage123DatabaseNumbersToStruct(stage3.Try1No3),
					No4:        util.Stage123DatabaseNumbersToStruct(stage3.Try1No4),
					No5:        util.Stage123DatabaseNumbersToStruct(stage3.Try1No5),
					No6:        util.Stage123DatabaseNumbersToStruct(stage3.Try1No6),
					Checkmarks: util.CheckmarksToBoolArray(stage3.Try1Checkmarks),
				},
				Try2: model.Stage13Try{
					Status:     string(stage3.Try2Status.Stage13Status),
					No1:        util.Stage123DatabaseNumbersToStruct(stage3.Try2No1.String),
					No2:        util.Stage123DatabaseNumbersToStruct(stage3.Try2No2.String),
					No3:        util.Stage123DatabaseNumbersToStruct(stage3.Try2No3.String),
					No4:        util.Stage123DatabaseNumbersToStruct(stage3.Try2No4.String),
					No5:        util.Stage123DatabaseNumbersToStruct(stage3.Try2No5.String),
					No6:        util.Stage123DatabaseNumbersToStruct(stage3.Try2No6.String),
					Checkmarks: util.CheckmarksToBoolArray(stage3.Try2Checkmarks.String),
				},
				IsTry2:      stage3.IsTry2,
				ShooterSign: stage3.ShooterSign,
				ScorerSign:  stage3.ScorerSign,
				CreatedAt:   stage3.CreatedAt.Time,
				UpdatedAt:   stage3.UpdatedAt.Time,
			}}
	} else {
		stage3Response = gin.H{
			"stage_3": model.Stage13{
				ID:       stage3.ID,
				ResultID: stage3.ResultID,
				Try1: model.Stage13Try{
					Status:     string(stage3.Try1Status),
					No1:        util.Stage123DatabaseNumbersToStruct(stage3.Try1No1),
					No2:        util.Stage123DatabaseNumbersToStruct(stage3.Try1No2),
					No3:        util.Stage123DatabaseNumbersToStruct(stage3.Try1No3),
					No4:        util.Stage123DatabaseNumbersToStruct(stage3.Try1No4),
					No5:        util.Stage123DatabaseNumbersToStruct(stage3.Try1No5),
					No6:        util.Stage123DatabaseNumbersToStruct(stage3.Try1No6),
					Checkmarks: util.CheckmarksToBoolArray(stage3.Try1Checkmarks),
				},
				IsTry2:      stage3.IsTry2,
				ShooterSign: stage3.ShooterSign,
				ScorerSign:  stage3.ScorerSign,
				CreatedAt:   stage3.CreatedAt.Time,
				UpdatedAt:   stage3.UpdatedAt.Time,
			}}
	}

	return util.ToWebServiceResponse("Berhasil mendapatkan hasil ujian stage 3", http.StatusOK, stage3Response)
}
