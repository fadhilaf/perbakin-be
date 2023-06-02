package usecase

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
)

func (usecase *superUsecaseImpl) UpdateExamStatus(req model.ByIdRequest) model.WebServiceResponse {
	active, err := usecase.Store.UpdateExamStatus(context.Background(), req.ID)
	if err != nil {
		return util.ToWebServiceResponse("Gagal mengubah status ujian: "+err.Error(), http.StatusInternalServerError, nil)
	}

	return util.ToWebServiceResponse("Berhasil mengubah status ujian", http.StatusOK, gin.H{
		"active": active,
	})
}
