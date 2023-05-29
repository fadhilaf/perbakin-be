package route

import (
	adminSuperDelivery "github.com/FadhilAF/perbakin-be/internal/delivery/admin-super"
	allDelivery "github.com/FadhilAF/perbakin-be/internal/delivery/all"

	"github.com/gin-gonic/gin"
)

// for adminRouter or examRouter (for super), return scorerRouter
func AdminSuperRoutes(adminExamRouter *gin.RouterGroup, delivery adminSuperDelivery.AdminSuperDelivery, allDelivery allDelivery.AllDelivery) {
	ExamAdminSuperRoutes(adminExamRouter, delivery)

	scorerRouter := ScorerAdminSuperRoutes(adminExamRouter, delivery)

	shooterRouter := ShooterAllRoutes(scorerRouter, allDelivery)
	ShooterAdminSuperRoutes(scorerRouter, delivery, allDelivery)

	resultRouter := ResultAllRoutes(shooterRouter, allDelivery)
	ResultAdminSuperRoutes(resultRouter, delivery)

	stage0Router := Stage0AllRoutes(resultRouter, allDelivery)
	Stage0AdminSuperRoutes(stage0Router, delivery)
}