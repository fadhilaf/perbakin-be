package usecase

import (
	"context"
	"net/http"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (usecase *allUsecaseImpl) CreateStage6(req model.ByResultIdRequest) model.WebServiceResponse {
	if _, err := usecase.Store.GetStage6RelationByResultId(context.Background(), req.ResultID); err == nil {
		return util.ToWebServiceResponse("Hasil ujian stage 6 sudah ada", http.StatusConflict, nil)
	}

	stage6, err := usecase.Store.CreateStage6(context.Background(), req.ResultID)
	if err != nil {
		return util.ToWebServiceResponse("Gagal membuat hasil ujian stage 6: "+err.Error(), http.StatusInternalServerError, nil)
	}

	return util.ToWebServiceResponse("Berhasil membuat hasil ujian stage 6", http.StatusCreated, gin.H{
		"stage_6": model.Stage46{
			ID:       stage6.ID,
			ResultID: stage6.ResultID,
			Try1: model.Stage46Try{
				Status:     string(stage6.Status),
				No1:        util.Stage46DatabaseNumbersToStruct(stage6.No1),
				No2:        util.Stage46DatabaseNumbersToStruct(stage6.No2),
				No3:        util.Stage46DatabaseNumbersToStruct(stage6.No3),
				Checkmarks: util.CheckmarksToBoolArray(stage6.Checkmarks),
			},
			IsTry2:      stage6.IsTry2,
			ShooterSign: stage6.ShooterSign,
			ScorerSign:  stage6.ScorerSign,
			CreatedAt:   stage6.CreatedAt.Time,
			UpdatedAt:   stage6.UpdatedAt.Time,
		},
	})
}
