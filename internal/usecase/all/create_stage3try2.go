package usecase

import (
	"context"
	"net/http"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (usecase *allUsecaseImpl) CreateStage3try2(req model.ByIdRequest) model.WebServiceResponse {
	exists, err := usecase.Store.GetStage3try2ExistById(context.Background(), req.ID)
	if err != nil {
		return util.ToWebServiceResponse("Gagal mengecek keberadaan hasil ujian stage 3 percobaan 2: "+err.Error(), http.StatusInternalServerError, nil)
	}
	if exists {
		return util.ToWebServiceResponse("Hasil ujian stage 3 percobaan 2 sudah ada", http.StatusConflict, nil)
	}

	stage3, err := usecase.Store.CreateStage3try2(context.Background(), req.ID)
	if err != nil {
		return util.ToWebServiceResponse("Gagal membuat hasil ujian stage 3 percobaan 2: "+err.Error(), http.StatusInternalServerError, nil)
	}

	return util.ToWebServiceResponse("Berhasil membuat hasil ujian stage 3 percobaan 2", http.StatusCreated, gin.H{
		"stage_3": model.CreateStage13try2{
			Try2: model.Stage13Try{
				Status:     string(stage3.Status),
				No1:        util.Stage123DatabaseNumbersToStruct(stage3.No1),
				No2:        util.Stage123DatabaseNumbersToStruct(stage3.No2),
				No3:        util.Stage123DatabaseNumbersToStruct(stage3.No3),
				No4:        util.Stage123DatabaseNumbersToStruct(stage3.No4),
				No5:        util.Stage123DatabaseNumbersToStruct(stage3.No5),
				No6:        util.Stage123DatabaseNumbersToStruct(stage3.No6),
				Checkmarks: util.CheckmarksToBoolArray(stage3.Checkmarks),
			},
			IsTry2: stage3.IsTry2,
		},
	})
}
