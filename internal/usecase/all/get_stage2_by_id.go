package usecase

import (
	"context"
	"net/http"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (usecase *allUsecaseImpl) GetStage2ById(req model.ByIdRequest) model.WebServiceResponse {
	stage2, err := usecase.Store.GetStage2ById(context.Background(), req.ID)
	if err != nil {
		return util.ToWebServiceResponse("Hasil ujian stage 2 tidak ada: "+err.Error(), http.StatusNotFound, nil)
	}

	var stage2Response gin.H

	if stage2.IsTry2 {
		stage2Response = gin.H{
			"stage_2": model.Stage2Full{
				ID:       stage2.ID,
				ResultID: stage2.ResultID,
				Try1: model.Stage2Try{
					Status:     string(stage2.Try1Status),
					No1:        util.NumbersToIntArrayArray(stage2.Try1No1),
					No2:        util.NumbersToIntArrayArray(stage2.Try1No2),
					No3:        util.NumbersToIntArrayArray(stage2.Try1No3),
					Checkmarks: util.CheckmarksToBoolArray(stage2.Try1Checkmarks),
				},
				Try2: model.Stage2Try{
					Status:     string(stage2.Try2Status.Stage246Status),
					No1:        util.NumbersToIntArrayArray(stage2.Try2No1.String),
					No2:        util.NumbersToIntArrayArray(stage2.Try2No2.String),
					No3:        util.NumbersToIntArrayArray(stage2.Try2No3.String),
					Checkmarks: util.CheckmarksToBoolArray(stage2.Try2Checkmarks.String),
				},
				IsTry2:      stage2.IsTry2,
				ShooterSign: stage2.ShooterSign,
				ScorerSign:  stage2.ScorerSign,
				CreatedAt:   stage2.CreatedAt.Time,
				UpdatedAt:   stage2.UpdatedAt.Time,
			}}
	} else {
		stage2Response = gin.H{
			"stage_2": model.Stage2{
				ID:       stage2.ID,
				ResultID: stage2.ResultID,
				Try1: model.Stage2Try{
					Status:     string(stage2.Try1Status),
					No1:        util.NumbersToIntArrayArray(stage2.Try1No1),
					No2:        util.NumbersToIntArrayArray(stage2.Try1No2),
					No3:        util.NumbersToIntArrayArray(stage2.Try1No3),
					Checkmarks: util.CheckmarksToBoolArray(stage2.Try1Checkmarks),
				},
				IsTry2:      stage2.IsTry2,
				ShooterSign: stage2.ShooterSign,
				ScorerSign:  stage2.ScorerSign,
				CreatedAt:   stage2.CreatedAt.Time,
				UpdatedAt:   stage2.UpdatedAt.Time,
			}}
	}

	return util.ToWebServiceResponse("Berhasil mendapatkan hasil ujian stage 2", http.StatusOK, stage2Response)
}
