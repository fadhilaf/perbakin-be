package session

import (
	"net/http"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/alexedwards/scs/pgxstore"
	"github.com/alexedwards/scs/v2"

	"github.com/gin-gonic/gin"

	"github.com/FadhilAF/perbakin-be/common/env"
)

var SessionManager *scs.SessionManager

func SessionHandler(handler *gin.Engine, dbPool *pgxpool.Pool, config env.Config) http.Handler {
	SessionManager = scs.New()
	SessionManager.Store = pgxstore.New(dbPool)
	SessionManager.Lifetime = 24 * time.Hour
	SessionManager.IdleTimeout = 20 * time.Minute
	SessionManager.Cookie.Persist = true
	SessionManager.Cookie.HttpOnly = true
	SessionManager.Cookie.Path = "/"
	SessionManager.Cookie.Name = "session"

	SessionManager.Cookie.Domain = config.AppDomain

	if config.Env == env.EnvProd {
		SessionManager.Cookie.SameSite = http.SameSiteStrictMode
		SessionManager.Cookie.Secure = true
	} else {
		SessionManager.Cookie.SameSite = http.SameSiteNoneMode
		SessionManager.Cookie.Secure = false
	}

	return SessionManager.LoadAndSave(handler)
}
