package util

import (
	"log"
	"net/http"

	"github.com/FadhilAF/perbakin-be/common/session"

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

func SetAuthStatusCookie(c *gin.Context, value string) {
	cookie := &http.Cookie{
		Name:     "auth",
		Value:    value,
		Path:     "/",
		Domain:   session.SessionManager.Cookie.Domain,
		HttpOnly: false,
		Secure:   session.SessionManager.Cookie.Secure,
		SameSite: session.SessionManager.Cookie.SameSite,
		MaxAge:   10800, //3 jam
	}

	http.SetCookie(c.Writer, cookie)
}

func DeleteAuthStatusCookie(c *gin.Context) {
	cookie := &http.Cookie{
		Name:     "auth",
		Value:    "",
		Path:     "/",
		Domain:   session.SessionManager.Cookie.Domain,
		HttpOnly: false,
		Secure:   session.SessionManager.Cookie.Secure,
		SameSite: session.SessionManager.Cookie.SameSite,
		MaxAge:   -1,
	}

	http.SetCookie(c.Writer, cookie)
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

func GetIdParam(c *gin.Context, param string) (uuid pgtype.UUID, ok bool) {
	id := c.Param(param)

	if err := uuid.Scan(id); err != nil {
		res := ToWebServiceResponse(param+" yang dimasukkan tidak valid", http.StatusBadRequest, nil)
		c.JSON(res.Status, res)
		return uuid, false
	}

	return uuid, true
}
