package usecase

import (
	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/repository"
)

type SuperUsecase interface {
	SuperLogin(model.LoginRequest) model.WebServiceResponse
	GetSuperByUserId(model.GetByUserIdRequest) model.WebServiceResponse

	CreateAdmin(model.CreateUserRequest) model.WebServiceResponse
	GetAdminById(model.ByIdRequest) model.WebServiceResponse
	GetAllAdmins() model.WebServiceResponse
	UpdateAdmin(model.UpdateUserRequest) model.WebServiceResponse
	DeleteAdmin(model.ByIdRequest) model.WebServiceResponse

	GetExamById(model.ByIdRequest) model.WebServiceResponse
	GetAllExams() model.WebServiceResponse
	CreateExam(model.CreateExamRequest) model.WebServiceResponse
	UpdateExam(model.UpdateExamRequest) model.WebServiceResponse
	DeleteExam(model.DeleteExamRequest) model.WebServiceResponse
}

var _ SuperUsecase = &superUsecaseImpl{}

func NewSuperUsecase(store repository.Store) SuperUsecase {
	return &superUsecaseImpl{
		Store: store,
	}
}

type superUsecaseImpl struct {
	repository.Store
}
