package app

import (
	"github.com/FadhilAF/perbakin-be/internal/route"
	"github.com/gin-gonic/gin"

	logoutController "github.com/FadhilAF/perbakin-be/internal/delivery/logout"
)

func (app *App) handlerV1(router *gin.RouterGroup) {
	router.POST("/logout", logoutController.Logout)

	SuperGroup := router.Group("/super")
	route.SuperRoutes(SuperGroup, app.delivery.super)

	AdminGroup := router.Group("/admin")
	route.AdminRoutes(AdminGroup, app.delivery.admin)
}
