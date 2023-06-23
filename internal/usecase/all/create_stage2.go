package usecase

import (
	"context"
	"net/http"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (usecase *allUsecaseImpl) CreateStage2(req model.ByResultIdRequest) model.WebServiceResponse {
	if _, err := usecase.Store.GetStage2RelationByResultId(context.Background(), req.ResultID); err == nil {
		return util.ToWebServiceResponse("Hasil ujian stage 2 sudah ada", http.StatusConflict, nil)
	}

	stage2, err := usecase.Store.CreateStage2(context.Background(), req.ResultID)
	if err != nil {
		return util.ToWebServiceResponse("Gagal membuat hasil ujian stage 2: "+err.Error(), http.StatusInternalServerError, nil)
	}

	return util.ToWebServiceResponse("Berhasil membuat hasil ujian stage 2", http.StatusCreated, gin.H{
		"stage_2": model.Stage2{
			ID:       stage2.ID,
			ResultID: stage2.ResultID,
			Try1: model.Stage2Try{
				Status:     string(stage2.Status),
				No1:        util.NumbersToIntArrayArray(stage2.No1),
				No2:        util.NumbersToIntArrayArray(stage2.No2),
				No3:        util.NumbersToIntArrayArray(stage2.No3),
				Checkmarks: util.CheckmarksToBoolArray(stage2.Checkmarks),
			},
			IsTry2:      stage2.IsTry2,
			ShooterSign: stage2.ShooterSign,
			ScorerSign:  stage2.ScorerSign,
			CreatedAt:   stage2.CreatedAt.Time,
			UpdatedAt:   stage2.UpdatedAt.Time,
		},
	})
}
