package usecase

import (
	"context"
	"net/http"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (usecase *allUsecaseImpl) CreateStage4(req model.ByResultIdRequest) model.WebServiceResponse {
	if _, err := usecase.Store.GetStage4RelationByResultId(context.Background(), req.ResultID); err == nil {
		return util.ToWebServiceResponse("Hasil ujian stage 4 sudah ada", http.StatusConflict, nil)
	}

	stage4, err := usecase.Store.CreateStage4(context.Background(), req.ResultID)
	if err != nil {
		return util.ToWebServiceResponse("Gagal membuat hasil ujian stage 4: "+err.Error(), http.StatusInternalServerError, nil)
	}

	return util.ToWebServiceResponse("Berhasil membuat hasil ujian stage 4", http.StatusCreated, gin.H{
		"stage_4": model.Stage46{
			ID:       stage4.ID,
			ResultID: stage4.ResultID,
			Try1: model.Stage46Try{
				Status:     string(stage4.Status),
				No1:        util.Stage46DatabaseNumbersToStruct(stage4.No1),
				No2:        util.Stage46DatabaseNumbersToStruct(stage4.No2),
				No3:        util.Stage46DatabaseNumbersToStruct(stage4.No3),
				Checkmarks: util.CheckmarksToBoolArray(stage4.Checkmarks),
			},
			IsTry2:      stage4.IsTry2,
			ShooterSign: stage4.ShooterSign,
			ScorerSign:  stage4.ScorerSign,
			CreatedAt:   stage4.CreatedAt.Time,
			UpdatedAt:   stage4.UpdatedAt.Time,
		},
	})
}
