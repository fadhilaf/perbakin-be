package usecase

import (
	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/repository"
)

type SuperUsecase interface {
	SuperLogin(model.LoginRequest) model.WebServiceResponse

	GetSuperRelationByUserId(model.UserByUserIdRequest) (model.SuperRelation, error)

	GetSuperByUserId(model.UserByUserIdRequest) model.WebServiceResponse

	CreateAdmin(model.CreateOperatorRequest) model.WebServiceResponse
	GetAllAdmins() model.WebServiceResponse
	GetAdminsByExamId(model.ByExamIdRequest) model.WebServiceResponse
	GetAdminById(model.OperatorByIdRequest) model.WebServiceResponse
	UpdateAdmin(model.UpdateOperatorRequest) model.WebServiceResponse
	DeleteAdmin(model.OperatorByIdRequest) model.WebServiceResponse

	GetExamById(model.ByIdRequest) model.WebServiceResponse
	GetAllExams() model.WebServiceResponse
	GetExamsBySuperId(model.GetExamsBySuperIdRequest) model.WebServiceResponse
	CreateExam(model.CreateExamRequest) model.WebServiceResponse
	UpdateExam(model.UpdateExamRequest) model.WebServiceResponse
	DeleteExam(model.DeleteExamRequest) model.WebServiceResponse

	GetExamRelationById(model.ByIdRequest) (model.ExamRelation, error)
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
