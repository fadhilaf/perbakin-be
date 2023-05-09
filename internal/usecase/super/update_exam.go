package usecase

import (
	"context"
	"fmt"
	"net/http"

	respositoryModel "github.com/FadhilAF/perbakin-be/internal/repository/postgres/sqlc"
	"github.com/gin-gonic/gin"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
)

func (usecase *superUsecaseImpl) UpdateExam(req model.UpdateExamRequest) model.WebServiceResponse {
	exam, err := usecase.Store.GetExamById(context.Background(), req.ID)
	if err != nil {
		return util.ToWebServiceResponse("Tidak ditemukan ujian dengan ID tersebut", http.StatusNotFound, nil)
	}

	if exam.SuperID != req.SuperID {
		return util.ToWebServiceResponse("Tidak dapat mengubah ujian super admin lain", http.StatusUnauthorized, nil)
	}

	newExam, err := usecase.Store.UpdateExam(context.Background(), respositoryModel.UpdateExamParams{
		ID:        req.ID,
		Name:      req.Body.Name,
		Location:  req.Body.Location,
		Organizer: req.Body.Organizer,
		Begin:     req.Body.Begin,
		Finish:    req.Body.Finish,
	})
	if err != nil {
		fmt.Println(err)
		return util.ToWebServiceResponse("Gagal mengubah ujian", http.StatusInternalServerError, nil)
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
