package middleware

import (
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func LoadSessionMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		util.GetUserIdFromSession(c)
	}
}
