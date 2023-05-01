package app

import (
	"github.com/FadhilAF/perbakin-be/internal/route"
	"github.com/gin-gonic/gin"
)

func (app *App) handlerV1(router *gin.RouterGroup) {
	AdminGroup := router.Group("/admin")
	route.AdminRoutes(AdminGroup, app.delivery.admin)
}
