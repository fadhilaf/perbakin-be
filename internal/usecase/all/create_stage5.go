package usecase

import (
	"context"
	"net/http"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (usecase *allUsecaseImpl) CreateStage5(req model.ByResultIdRequest) model.WebServiceResponse {
	if _, err := usecase.Store.GetStage5RelationByResultId(context.Background(), req.ResultID); err == nil {
		return util.ToWebServiceResponse("Hasil ujian stage 5 sudah ada", http.StatusConflict, nil)
	}

	stage5, err := usecase.Store.CreateStage5(context.Background(), req.ResultID)
	if err != nil {
		return util.ToWebServiceResponse("Gagal membuat hasil ujian stage 5: "+err.Error(), http.StatusInternalServerError, nil)
	}

	return util.ToWebServiceResponse("Berhasil membuat hasil ujian stage 5", http.StatusCreated, gin.H{
		"stage_5": model.Stage5{
			ID:       stage5.ID,
			ResultID: stage5.ResultID,
			Try1: model.Stage5Try{
				Status:     string(stage5.Status),
				No1:        util.Stage5DatabaseNumbersToStruct(stage5.No1),
				No2:        util.Stage5DatabaseNumbersToStruct(stage5.No2),
				Checkmarks: util.CheckmarksToBoolArray(stage5.Checkmarks),
			},
			IsTry2:      stage5.IsTry2,
			ShooterSign: stage5.ShooterSign,
			ScorerSign:  stage5.ScorerSign,
			CreatedAt:   stage5.CreatedAt.Time,
			UpdatedAt:   stage5.UpdatedAt.Time,
		},
	})
}
