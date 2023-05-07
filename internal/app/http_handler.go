package app

import (
	"github.com/FadhilAF/perbakin-be/internal/route"
	"github.com/gin-gonic/gin"
)

func (app *App) handlerV1(router *gin.RouterGroup) {
	AuthGroup := router.Group("/auth")
	route.AuthRoutes(AuthGroup, app.delivery.auth)

	SuperGroup := router.Group("/super")
	route.SuperRoutes(SuperGroup, app.delivery.super)

	AdminGroup := router.Group("/admin")
	route.AdminRoutes(AdminGroup, app.delivery.admin)
}
