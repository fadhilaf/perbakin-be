package app

import (
	"github.com/FadhilAF/perbakin-be/internal/route"
	"github.com/gin-gonic/gin"

	logoutController "github.com/FadhilAF/perbakin-be/internal/delivery/logout"
)

func (app *App) handlerV1(router *gin.RouterGroup) {
	router.POST("/logout", logoutController.Logout)

	superGroup := router.Group("/super")
	route.SuperRoutes(superGroup, app.delivery.super)

	adminGroup := router.Group("/admin")
	route.AdminRoutes(adminGroup, app.delivery.admin)

	adminSuperGroup := router.Group("/admin-super")
	route.AdminSuperRoutes(adminSuperGroup, app.delivery.adminSuper)
}
