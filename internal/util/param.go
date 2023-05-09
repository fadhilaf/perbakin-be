package util

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
)

func GetIdParam(c *gin.Context) (uuid pgtype.UUID, ok bool) {
	id := c.Param("id")

	if err := uuid.Scan(id); err != nil {
		res := ToWebServiceResponse("id yang dimasukkan tidak valid", http.StatusBadRequest, nil)
		c.JSON(res.Status, res)
		return uuid, false
	}

	return uuid, true
}
