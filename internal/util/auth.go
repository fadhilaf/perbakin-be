package util

import (
	"log"
	"net/http"

	"github.com/FadhilAF/perbakin-be/common/session"
	"github.com/FadhilAF/perbakin-be/internal/model"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
)

func SaveUserToSession(c *gin.Context, uuid pgtype.UUID) {
	userId, err := uuid.Value()
	if err != nil {
		log.Println("Error ketika membaca pgtype UUID:", err)
	}
	session.SessionManager.Put(c.Request.Context(), "user_id", userId)
}

func GetUserIdFromSession(c *gin.Context) {
	var uuid pgtype.UUID

	userIdSession := session.SessionManager.GetString(c.Request.Context(), "user_id")

	if userIdSession != "" {
		if err := uuid.Scan(userIdSession); err == nil {
			c.Set("user_id", userIdSession)
		} else {
			log.Println("Error ketika scan uuid:", err)
		}
	}

}

func GetUserIdFromContext(c *gin.Context) pgtype.UUID {
	var uuid pgtype.UUID

	userIdSession, exist := c.Get("user_id")
	if exist {
		if parsedUserIdSession, ok := userIdSession.(string); ok {
			if err := uuid.Scan(parsedUserIdSession); err != nil {
				log.Println("Error ketika scan uuid:", err)
			}
		}
	}
	return uuid
}

func RemoveAuthSession(c *gin.Context) {
	session.SessionManager.Remove(c.Request.Context(), "user_id")
}

func MustGetSuper(c *gin.Context) (super model.Super, ok bool) {
	super, ok = c.MustGet("super").(model.Super)
	if !ok {
		res := ToWebServiceResponse("Error ketika parsing data super", http.StatusInternalServerError, nil)
		c.JSON(res.Status, res)
	}

	return super, ok
}
