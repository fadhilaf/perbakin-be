package usecase

import (
	"context"
	"net/http"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (usecase *allUsecaseImpl) GetStage1ById(req model.ByIdRequest) model.WebServiceResponse {
	stage1, err := usecase.Store.GetStage1ById(context.Background(), req.ID)
	if err != nil {
		return util.ToWebServiceResponse("Hasil ujian stage 1 tidak ada: "+err.Error(), http.StatusNotFound, nil)
	}

	var stage1Response gin.H

	if stage1.IsTry2 {
		stage1Response = gin.H{
			"stage_1": model.Stage13Full{
				ID:       stage1.ID,
				ResultID: stage1.ResultID,
				Try1: model.Stage13Try{
					Status:     string(stage1.Try1Status),
					No1:        util.Stage123DatabaseNumbersToStruct(stage1.Try1No1),
					No2:        util.Stage123DatabaseNumbersToStruct(stage1.Try1No2),
					No3:        util.Stage123DatabaseNumbersToStruct(stage1.Try1No3),
					No4:        util.Stage123DatabaseNumbersToStruct(stage1.Try1No4),
					No5:        util.Stage123DatabaseNumbersToStruct(stage1.Try1No5),
					No6:        util.Stage123DatabaseNumbersToStruct(stage1.Try1No6),
					Checkmarks: util.CheckmarksToBoolArray(stage1.Try1Checkmarks),
				},
				Try2: model.Stage13Try{
					Status:     string(stage1.Try2Status.Stage13Status),
					No1:        util.Stage123DatabaseNumbersToStruct(stage1.Try2No1.String),
					No2:        util.Stage123DatabaseNumbersToStruct(stage1.Try2No2.String),
					No3:        util.Stage123DatabaseNumbersToStruct(stage1.Try2No3.String),
					No4:        util.Stage123DatabaseNumbersToStruct(stage1.Try2No4.String),
					No5:        util.Stage123DatabaseNumbersToStruct(stage1.Try2No5.String),
					No6:        util.Stage123DatabaseNumbersToStruct(stage1.Try2No6.String),
					Checkmarks: util.CheckmarksToBoolArray(stage1.Try2Checkmarks.String),
				},
				IsTry2:      stage1.IsTry2,
				ShooterSign: stage1.ShooterSign,
				ScorerSign:  stage1.ScorerSign,
				CreatedAt:   stage1.CreatedAt.Time,
				UpdatedAt:   stage1.UpdatedAt.Time,
			}}
	} else {
		stage1Response = gin.H{
			"stage_1": model.Stage13{
				ID:       stage1.ID,
				ResultID: stage1.ResultID,
				Try1: model.Stage13Try{
					Status:     string(stage1.Try1Status),
					No1:        util.Stage123DatabaseNumbersToStruct(stage1.Try1No1),
					No2:        util.Stage123DatabaseNumbersToStruct(stage1.Try1No2),
					No3:        util.Stage123DatabaseNumbersToStruct(stage1.Try1No3),
					No4:        util.Stage123DatabaseNumbersToStruct(stage1.Try1No4),
					No5:        util.Stage123DatabaseNumbersToStruct(stage1.Try1No5),
					No6:        util.Stage123DatabaseNumbersToStruct(stage1.Try1No6),
					Checkmarks: util.CheckmarksToBoolArray(stage1.Try1Checkmarks),
				},
				IsTry2:      stage1.IsTry2,
				ShooterSign: stage1.ShooterSign,
				ScorerSign:  stage1.ScorerSign,
				CreatedAt:   stage1.CreatedAt.Time,
				UpdatedAt:   stage1.UpdatedAt.Time,
			}}
	}

	return util.ToWebServiceResponse("Berhasil mendapatkan hasil ujian stage 1", http.StatusOK, stage1Response)
}
