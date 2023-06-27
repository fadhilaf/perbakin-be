package usecase

import (
	"context"
	"net/http"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (usecase *allUsecaseImpl) CreateStage5try2(req model.ByIdRequest) model.WebServiceResponse {
	exists, err := usecase.Store.GetStage5try2ExistById(context.Background(), req.ID)
	if err != nil {
		return util.ToWebServiceResponse("Gagal mengecek keberadaan hasil ujian stage 5 percobaan 2: "+err.Error(), http.StatusInternalServerError, nil)
	}
	if exists {
		return util.ToWebServiceResponse("Hasil ujian stage 5 percobaan 2 sudah ada", http.StatusConflict, nil)
	}

	stage5, err := usecase.Store.CreateStage5try2(context.Background(), req.ID)
	if err != nil {
		return util.ToWebServiceResponse("Gagal membuat hasil ujian stage 5 percobaan 2: "+err.Error(), http.StatusInternalServerError, nil)
	}

	return util.ToWebServiceResponse("Berhasil membuat hasil ujian stage 5 percobaan 2", http.StatusCreated, gin.H{
		"stage_5": model.CreateStage5try2{
			Try2: model.Stage5Try{
				Status:     string(stage5.Status),
				No1:        util.Stage5DatabaseNumbersToStruct(stage5.No1),
				No2:        util.Stage5DatabaseNumbersToStruct(stage5.No2),
				Checkmarks: util.CheckmarksToBoolArray(stage5.Checkmarks),
			},
			IsTry2: stage5.IsTry2,
		},
	})
}
