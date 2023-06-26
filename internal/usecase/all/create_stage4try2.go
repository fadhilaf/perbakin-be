package usecase

import (
	"context"
	"net/http"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (usecase *allUsecaseImpl) CreateStage4try2(req model.ByIdRequest) model.WebServiceResponse {
	exists, err := usecase.Store.GetStage4try2ExistById(context.Background(), req.ID)
	if err != nil {
		return util.ToWebServiceResponse("Gagal mengecek keberadaan hasil ujian stage 4 percobaan 2: "+err.Error(), http.StatusInternalServerError, nil)
	}
	if exists {
		return util.ToWebServiceResponse("Hasil ujian stage 4 percobaan 2 sudah ada", http.StatusConflict, nil)
	}

	stage4, err := usecase.Store.CreateStage4try2(context.Background(), req.ID)
	if err != nil {
		return util.ToWebServiceResponse("Gagal membuat hasil ujian stage 4 percobaan 2: "+err.Error(), http.StatusInternalServerError, nil)
	}

	return util.ToWebServiceResponse("Berhasil membuat hasil ujian stage 4 percobaan 2", http.StatusCreated, gin.H{
		"stage_4": model.CreateStage46try2{
			Try2: model.Stage46Try{
				Status:     string(stage4.Status),
				No1:        util.Stage46DatabaseNumbersToStruct(stage4.No1),
				No2:        util.Stage46DatabaseNumbersToStruct(stage4.No2),
				No3:        util.Stage46DatabaseNumbersToStruct(stage4.No3),
				Checkmarks: util.CheckmarksToBoolArray(stage4.Checkmarks),
			},
			IsTry2: stage4.IsTry2,
		},
	})
}
