package usecase

import (
	"context"
	"net/http"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (usecase *superUsecaseImpl) GetExamsByUserId(req model.GetExamsBySuperIdRequest) model.WebServiceResponse {
	exams, err := usecase.Store.GetExamsBySuperId(context.Background(), req.SuperID)
	if err != nil {
		return util.ToWebServiceResponse("Data ujian tidak ditemukan", http.StatusUnauthorized, nil)
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
