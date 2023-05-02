package middleware

import (
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func SaveAndLoadSessionMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := util.GetUserIdFromSession(c)

		if userId.Valid {
			c.Set("user_id", userId.Bytes)
		}
	}
}
