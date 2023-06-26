package usecase

import (
	"context"
	"net/http"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (usecase *allUsecaseImpl) CreateStage3(req model.ByResultIdRequest) model.WebServiceResponse {
	if _, err := usecase.Store.GetStage3RelationByResultId(context.Background(), req.ResultID); err == nil {
		return util.ToWebServiceResponse("Hasil ujian stage 3 sudah ada", http.StatusConflict, nil)
	}

	stage3, err := usecase.Store.CreateStage3(context.Background(), req.ResultID)
	if err != nil {
		return util.ToWebServiceResponse("Gagal membuat hasil ujian stage 3: "+err.Error(), http.StatusInternalServerError, nil)
	}

	return util.ToWebServiceResponse("Berhasil membuat hasil ujian stage 3", http.StatusCreated, gin.H{
		"stage_3": model.Stage13{
			ID:       stage3.ID,
			ResultID: stage3.ResultID,
			Try1: model.Stage13Try{
				Status:     string(stage3.Status),
				No1:        util.Stage123DatabaseNumbersToStruct(stage3.No1),
				No2:        util.Stage123DatabaseNumbersToStruct(stage3.No2),
				No3:        util.Stage123DatabaseNumbersToStruct(stage3.No3),
				No4:        util.Stage123DatabaseNumbersToStruct(stage3.No4),
				No5:        util.Stage123DatabaseNumbersToStruct(stage3.No5),
				No6:        util.Stage123DatabaseNumbersToStruct(stage3.No6),
				Checkmarks: util.CheckmarksToBoolArray(stage3.Checkmarks),
			},
			IsTry2:      stage3.IsTry2,
			ShooterSign: stage3.ShooterSign,
			ScorerSign:  stage3.ScorerSign,
			CreatedAt:   stage3.CreatedAt.Time,
			UpdatedAt:   stage3.UpdatedAt.Time,
		},
	})
}
