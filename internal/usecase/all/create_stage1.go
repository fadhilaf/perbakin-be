package usecase

import (
	"context"
	"net/http"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (usecase *allUsecaseImpl) CreateStage1(req model.ByResultIdRequest) model.WebServiceResponse {
	if _, err := usecase.Store.GetStage1RelationByResultId(context.Background(), req.ResultID); err == nil {
		return util.ToWebServiceResponse("Hasil ujian stage 1 sudah ada", http.StatusConflict, nil)
	}

	stage1, err := usecase.Store.CreateStage1(context.Background(), req.ResultID)
	if err != nil {
		return util.ToWebServiceResponse("Gagal membuat hasil ujian stage 1: "+err.Error(), http.StatusInternalServerError, nil)
	}

	return util.ToWebServiceResponse("Berhasil membuat hasil ujian stage 1", http.StatusCreated, gin.H{
		"stage_1": model.Stage1{
			ID:       stage1.ID,
			ResultID: stage1.ResultID,
			Try1: model.Stage13Try{
				ID:         stage1.Try1ID,
				Status:     string(stage1.Status),
				No1:        util.NumbersToIntArrayArray(stage1.No1),
				No2:        util.NumbersToIntArrayArray(stage1.No2),
				No3:        util.NumbersToIntArrayArray(stage1.No3),
				No4:        util.NumbersToIntArrayArray(stage1.No4),
				No5:        util.NumbersToIntArrayArray(stage1.No5),
				No6:        util.NumbersToIntArrayArray(stage1.No6),
				Checkmarks: util.CheckmarksToBoolArray(stage1.Checkmarks),
			},
			IsTry2:      stage1.IsTry2,
			ShooterSign: stage1.ShooterSign,
			ScorerSign:  stage1.ScorerSign,
			CreatedAt:   stage1.CreatedAt.Time,
			UpdatedAt:   stage1.UpdatedAt.Time,
		},
	})
}
