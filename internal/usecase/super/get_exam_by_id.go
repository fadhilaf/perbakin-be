package usecase

import (
	"context"
	"net/http"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (usecase *superUsecaseImpl) GetExamById(req model.ByIdRequest) model.WebServiceResponse {
	exam, _ := usecase.Store.GetExamById(context.Background(), req.ID)
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
