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
	SessionManager.IdleTimeout = 3 * time.Hour
	SessionManager.Cookie.Persist = true
	SessionManager.Cookie.Domain = config.AppDomain
	SessionManager.Cookie.Path = "/"
	SessionManager.Cookie.Name = "session"
	SessionManager.Cookie.HttpOnly = true
	SessionManager.Cookie.Secure = true
	// kalo mau same site none, harus pake secure
	// kalo mau httpOnly true, harus pake secure

	// jadi tidak ada cara lain selain pake secure cookie.

	//kalo meskipun masih development, tetep pake secure cookie
	//soalnyo kalo engga, cookie engga bakal kebaca di browser
	//jika same site none dan httpOnly false, maka secure harus true (aturanny)
	//kita ada salah satu cookie yg mesti bisa dibaca di browser, yaitu cookie "auth"
	//tapi, kalo kito pake secure cookie, cookie auth engga bakal kebaca di browser, karena localhost engga pake https TT-TT

	if config.Env == env.EnvProd {
		SessionManager.Cookie.SameSite = http.SameSiteStrictMode
	} else {
		SessionManager.Cookie.SameSite = http.SameSiteNoneMode
	}

	return SessionManager.LoadAndSave(handler)
}
