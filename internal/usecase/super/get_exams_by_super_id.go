package usecase

import (
	"context"
	"net/http"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (usecase *superUsecaseImpl) GetExamsBySuperId(req model.GetExamsBySuperIdRequest) model.WebServiceResponse {
	exams, err := usecase.Store.GetExamsBySuperId(context.Background(), req.SuperID)
	if err != nil {
		return util.ToWebServiceResponse("Data ujian tidak ditemukan", http.StatusUnauthorized, nil)
	}

	var examsData []model.ExamDisplayData
	for _, exam := range exams {
		examsData = append(examsData, model.ExamDisplayData{
			ID:        exam.ID,
			Name:      exam.Name,
			Location:  exam.Location,
			Organizer: exam.Organizer,
			Begin:     exam.Begin,
			Finish:    exam.Finish,
		})
	}

	return util.ToWebServiceResponse("Berhasil mendapatkan data ujian", http.StatusOK, gin.H{
		"exams": examsData,
	})
}
