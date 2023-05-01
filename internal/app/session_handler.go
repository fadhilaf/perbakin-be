package app

import (
	"net/http"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/alexedwards/scs/pgxstore"
	"github.com/alexedwards/scs/v2"

	"github.com/gin-gonic/gin"
)

var SessionManager *scs.SessionManager

func SessionHandler(dbPool *pgxpool.Pool, handler *gin.Engine) http.Handler {
	SessionManager = scs.New()
	SessionManager.Store = pgxstore.New(dbPool)
	SessionManager.Lifetime = 24 * time.Hour
	SessionManager.IdleTimeout = 20 * time.Minute
	SessionManager.Cookie.Name = "session_id"

	// SessionManager.Cookie.Domain = "example.com"

	SessionManager.Cookie.HttpOnly = true

	//path cookie
	SessionManager.Cookie.Path = "/"

	SessionManager.Cookie.Persist = true

	// SessionManager.Cookie.SameSite = http.SameSiteStrictMode
	SessionManager.Cookie.SameSite = http.SameSiteNoneMode

	SessionManager.Cookie.Secure = true

	return SessionManager.LoadAndSave(handler)
}
