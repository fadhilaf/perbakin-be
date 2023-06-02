package usecase

import (
	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/repository"
)

type SuperUsecase interface {
	SuperLogin(model.LoginRequest) model.WebServiceResponse

	GetSuperRelationByUserId(model.UserByUserIdRequest) (model.SuperRelation, error)

	GetSuperByUserId(model.UserByUserIdRequest) model.WebServiceResponse

	GetAdminRelationById(model.ByIdRequest) (model.OperatorRelation, error)

	CreateAdmin(model.CreateOperatorRequest) model.WebServiceResponse
	GetAllAdmins() model.WebServiceResponse
	GetAdminsByExamId(model.ByExamIdRequest) model.WebServiceResponse
	GetAdminById(model.ByIdRequest) model.WebServiceResponse
	UpdateAdmin(model.UpdateOperatorRequest) model.WebServiceResponse
	DeleteAdmin(model.UserByUserIdRequest) model.WebServiceResponse

	GetExamById(model.ByIdRequest) model.WebServiceResponse
	GetAllExams() model.WebServiceResponse
	GetExamsBySuperId(model.GetExamsBySuperIdRequest) model.WebServiceResponse
	CreateExam(model.CreateExamRequest) model.WebServiceResponse
	UpdateExam(model.UpdateExamRequest) model.WebServiceResponse
	UpdateExamStatus(model.ByIdRequest) model.WebServiceResponse
	DeleteExam(model.ByIdRequest) model.WebServiceResponse

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
