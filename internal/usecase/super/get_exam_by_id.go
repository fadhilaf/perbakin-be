package usecase

import (
	"context"
	"net/http"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (usecase *superUsecaseImpl) GetExamById(req model.ByIdRequest) model.WebServiceResponse {
	exam, err := usecase.Store.GetExamById(context.Background(), req.ID)
	if err != nil {
		return util.ToWebServiceResponse("Data ujian tidak ditemukan", http.StatusUnauthorized, nil)
	}

	return util.ToWebServiceResponse("Data ujian ditemukan", http.StatusOK, gin.H{
		"exam": model.Exam{
			ID:        exam.ID,
			SuperID:   exam.SuperID,
			Name:      exam.Name,
			Location:  exam.Location,
			Organizer: exam.Organizer,
			Begin:     exam.Begin,
			Finish:    exam.Finish,
			CreatedAt: exam.CreatedAt.Time,
			UpdatedAt: exam.UpdatedAt.Time,
		}})
}
