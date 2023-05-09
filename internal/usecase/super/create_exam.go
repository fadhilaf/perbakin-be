package usecase

import (
	"context"
	"net/http"

	"github.com/FadhilAF/perbakin-be/internal/model"
	repositoryModel "github.com/FadhilAF/perbakin-be/internal/repository/postgres/sqlc"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (usecase *superUsecaseImpl) CreateExam(req model.CreateExamRequest) model.WebServiceResponse {
	if _, err := usecase.Store.GetExamByName(context.Background(), req.Body.Name); err == nil {
		return util.ToWebServiceResponse("Nama ujian sudah digunakan", http.StatusConflict, nil)
	}

	exam, err := usecase.Store.CreateExam(context.Background(), repositoryModel.CreateExamParams{
		SuperID:   req.SuperID,
		Name:      req.Body.Name,
		Location:  req.Body.Location,
		Organizer: req.Body.Organizer,
		Begin:     req.Body.Begin,
		Finish:    req.Body.Finish,
	})
	if err != nil {
		return util.ToWebServiceResponse("Gagal membuat ujian: "+err.Error(), http.StatusInternalServerError, nil)
	}

	return util.ToWebServiceResponse("Berhasil membuat ujian", http.StatusCreated, gin.H{
		"exam": exam,
	})
}
