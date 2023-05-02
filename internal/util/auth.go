package util

import (
	"fmt"

	"github.com/FadhilAF/perbakin-be/common/session"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
)

func SaveUserToSession(c *gin.Context, id pgtype.UUID) {
	session.SessionManager.Put(c, "user_id", id.Bytes)
}

func GetUserIdFromSession(c *gin.Context) pgtype.UUID {
	var uuid pgtype.UUID

	userIdSession := session.SessionManager.GetBytes(c, "user_id")

	if userIdSession != nil {
		if err := uuid.Scan(userIdSession); err != nil {
			fmt.Println("Error ketika scan uuid:", err)
		}
	}

	return uuid
}

func GetUserIdFromContext(c *gin.Context) pgtype.UUID {
	var uuid pgtype.UUID

	userId, exist := c.Get("user_id")
	if exist {
		if parsedUserId, ok := userId.([16]byte); ok {
			if err := uuid.Scan(parsedUserId); err != nil {
				fmt.Println("Error ketika scan uuid:", err)
			}
		}
	}
	return uuid
}

func RemoveAuthSession(c *gin.Context) {
	if session.SessionManager.Exists(c, "user_id") {
		session.SessionManager.Remove(c, "user_id")
	}
}
