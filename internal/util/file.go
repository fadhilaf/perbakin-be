package util

import (
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func SaveFileFromForm(ctx *gin.Context, field string, path string) (filename string, ok bool) {
	file, err := ctx.FormFile(field)
	if err != nil {
		res := ToWebServiceResponse("Tidak ada file '"+field+"' dari form", http.StatusBadRequest, nil)
		ctx.JSON(http.StatusBadRequest, res)
		ctx.Abort()

		return "", false
	}

	extension := filepath.Base(file.Filename)

	newFileName := uuid.New().String() + extension

	if err := ctx.SaveUploadedFile(file, path+newFileName); err != nil {
		res := ToWebServiceResponse("Gagal menyimpan file ke folder", http.StatusInternalServerError, nil)
		ctx.JSON(http.StatusInternalServerError, res)
		ctx.Abort()
		return "", false
	}

	return newFileName, true
}
