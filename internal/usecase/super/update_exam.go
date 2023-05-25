package usecase

import (
	"context"
	"net/http"

	repositoryModel "github.com/FadhilAF/perbakin-be/internal/repository/postgres/sqlc"
	"github.com/gin-gonic/gin"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
)

func (usecase *superUsecaseImpl) UpdateExam(req model.UpdateExamRequest) model.WebServiceResponse {
	newExam, err := usecase.Store.UpdateExam(context.Background(), repositoryModel.UpdateExamParams{
		ID:        req.ID,
		Name:      req.Body.Name,
		Location:  req.Body.Location,
		Organizer: req.Body.Organizer,
		Begin:     req.Body.Begin,
		Finish:    req.Body.Finish,
	})
	if err != nil {
		return util.ToWebServiceResponse("Gagal mengubah ujian: "+err.Error(), http.StatusInternalServerError, nil)
	}

	return util.ToWebServiceResponse("Berhasil mengubah ujian", http.StatusOK, gin.H{
		"exam": model.Exam{
			ID:        newExam.ID,
			SuperID:   newExam.SuperID,
			Name:      newExam.Name,
			Location:  newExam.Location,
			Organizer: newExam.Organizer,
			Begin:     newExam.Begin,
			Finish:    newExam.Finish,
		},
	})
}
