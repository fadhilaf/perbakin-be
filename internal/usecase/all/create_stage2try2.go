package usecase

import (
	"context"
	"net/http"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (usecase *allUsecaseImpl) CreateStage2try2(req model.ByIdRequest) model.WebServiceResponse {
	exists, err := usecase.Store.GetStage2try2ExistById(context.Background(), req.ID)
	if err != nil {
		return util.ToWebServiceResponse("Gagal mengecek keberadaan hasil ujian stage 1 percobaan 2: "+err.Error(), http.StatusInternalServerError, nil)
	}
	if exists {
		return util.ToWebServiceResponse("Hasil ujian stage 1 percobaan 2 sudah ada", http.StatusConflict, nil)
	}

	stage2, err := usecase.Store.CreateStage2try2(context.Background(), req.ID)
	if err != nil {
		return util.ToWebServiceResponse("Gagal membuat hasil ujian stage 1 percobaan 2: "+err.Error(), http.StatusInternalServerError, nil)
	}

	return util.ToWebServiceResponse("Berhasil membuat hasil ujian stage 1 percobaan 2", http.StatusCreated, gin.H{
		"stage_1": model.CreateStage2try2{
			Try2: model.Stage2Try{
				Status:     string(stage2.Status),
				No1:        util.Stage123DatabaseNumbersToStruct(stage2.No1),
				No2:        util.Stage123DatabaseNumbersToStruct(stage2.No2),
				No3:        util.Stage123DatabaseNumbersToStruct(stage2.No3),
				Checkmarks: util.CheckmarksToBoolArray(stage2.Checkmarks),
			},
			IsTry2: stage2.IsTry2,
		},
	})
}
