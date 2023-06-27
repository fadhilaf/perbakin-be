package usecase

import (
	"context"
	"net/http"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (usecase *allUsecaseImpl) CreateStage6try2(req model.ByIdRequest) model.WebServiceResponse {
	exists, err := usecase.Store.GetStage6try2ExistById(context.Background(), req.ID)
	if err != nil {
		return util.ToWebServiceResponse("Gagal mengecek keberadaan hasil ujian stage 6 percobaan 2: "+err.Error(), http.StatusInternalServerError, nil)
	}
	if exists {
		return util.ToWebServiceResponse("Hasil ujian stage 6 percobaan 2 sudah ada", http.StatusConflict, nil)
	}

	stage6, err := usecase.Store.CreateStage6try2(context.Background(), req.ID)
	if err != nil {
		return util.ToWebServiceResponse("Gagal membuat hasil ujian stage 6 percobaan 2: "+err.Error(), http.StatusInternalServerError, nil)
	}

	return util.ToWebServiceResponse("Berhasil membuat hasil ujian stage 6 percobaan 2", http.StatusCreated, gin.H{
		"stage_6": model.CreateStage46try2{
			Try2: model.Stage46Try{
				Status:     string(stage6.Status),
				No1:        util.Stage46DatabaseNumbersToStruct(stage6.No1),
				No2:        util.Stage46DatabaseNumbersToStruct(stage6.No2),
				No3:        util.Stage46DatabaseNumbersToStruct(stage6.No3),
				Checkmarks: util.CheckmarksToBoolArray(stage6.Checkmarks),
			},
			IsTry2: stage6.IsTry2,
		},
	})
}
