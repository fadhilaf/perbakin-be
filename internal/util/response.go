package util

import (
	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/gin-gonic/gin"
)

func ToWebServiceResponse(message string, status int, data gin.H) model.WebServiceResponse {
	return model.WebServiceResponse{
		Message: message,
		Status:  status,
		Data:    data,
	}
}
