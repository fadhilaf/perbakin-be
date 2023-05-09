package usecase

import (
	"context"
	"net/http"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (usecase *superUsecaseImpl) GetAllExams() model.WebServiceResponse {
	exams, err := usecase.Store.GetAllExams(context.Background())
	if err != nil {
		return util.ToWebServiceResponse("Gagal mendapatkan data ujian", http.StatusInternalServerError, nil)
	}

	var examsData []model.Exam
	for _, exam := range exams {
		examsData = append(examsData, model.Exam{
			ID:        exam.ID,
			SuperID:   exam.SuperID,
			Name:      exam.Name,
			Location:  exam.Location,
			Organizer: exam.Organizer,
			Begin:     exam.Begin,
			Finish:    exam.Finish,
			CreatedAt: exam.CreatedAt.Time,
			UpdatedAt: exam.UpdatedAt.Time,
		})
	}

	return util.ToWebServiceResponse("Berhasil mendapatkan data ujian", http.StatusOK, gin.H{
		"exams": examsData,
	})
}
