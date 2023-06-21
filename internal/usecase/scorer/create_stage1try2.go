package usecase

import (
	"context"
	"net/http"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (usecase *scorerUsecaseImpl) CreateStage1try2(req model.ByIdRequest) model.WebServiceResponse {
	exists, err := usecase.Store.GetStage1try2ExistById(context.Background(), req.ID)
	if err != nil {
		return util.ToWebServiceResponse("Gagal mengecek keberadaan hasil ujian stage 1 percobaan 2: "+err.Error(), http.StatusInternalServerError, nil)
	}
	if exists {
		return util.ToWebServiceResponse("Hasil ujian stage 1 percobaan 2 sudah ada", http.StatusConflict, nil)
	}

	status, _ := usecase.Store.GetStage1try1Status(context.Background(), req.ID)
	if status != "6" {
		return util.ToWebServiceResponse("Tidak dapat membuat stage 1 percobaan 2. Percobaan 1 belum selesai, masih pada nomor ke-"+string(status), http.StatusForbidden, nil)
	}

	stage1, err := usecase.Store.CreateStage1try2(context.Background(), req.ID)
	if err != nil {
		return util.ToWebServiceResponse("Gagal membuat hasil ujian stage 1 percobaan 2: "+err.Error(), http.StatusInternalServerError, nil)
	}

	return util.ToWebServiceResponse("Berhasil membuat hasil ujian stage 1 percobaan 2", http.StatusCreated, gin.H{
		"stage_1": model.CreateStage1try2{
			Try2: model.Stage13Try{
				Status:     string(stage1.Status),
				No1:        util.NumbersToIntArrayArray(stage1.No1),
				No2:        util.NumbersToIntArrayArray(stage1.No2),
				No3:        util.NumbersToIntArrayArray(stage1.No3),
				No4:        util.NumbersToIntArrayArray(stage1.No4),
				No5:        util.NumbersToIntArrayArray(stage1.No5),
				No6:        util.NumbersToIntArrayArray(stage1.No6),
				Checkmarks: util.CheckmarksToBoolArray(stage1.Checkmarks),
			},
			IsTry2: stage1.IsTry2,
		},
	})
}
