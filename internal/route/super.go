package route

import (
	adminSuperDelivery "github.com/FadhilAF/perbakin-be/internal/delivery/admin-super"
	allDelivery "github.com/FadhilAF/perbakin-be/internal/delivery/all"
	superDelivery "github.com/FadhilAF/perbakin-be/internal/delivery/super"

	"github.com/gin-gonic/gin"
)

func SuperRoutes(router *gin.RouterGroup, delivery superDelivery.SuperDelivery, allDelivery allDelivery.AllDelivery, adminSuperDelivery adminSuperDelivery.AdminSuperDelivery) {
	router.POST("/login", delivery.SuperLogin)

	superRouter := router.Group("/", delivery.MustSuperMiddleware())
	superRouter.GET("/", delivery.CheckSuperLogin)

	superRouter.GET("/exams", delivery.GetAllExams)                 //done
	superRouter.GET("/admins", delivery.GetAllAdmins)               //done
	superRouter.GET("/scorers", adminSuperDelivery.GetAllScorers)   //done
	superRouter.GET("/shooters", adminSuperDelivery.GetAllShooters) //done

	superRouter.GET("/exam", delivery.GetExamsBySuperId) //done
	superRouter.POST("/exam", delivery.CreateExam)       //done

	examRouter := superRouter.Group("/exam/:exam_id", delivery.MustExamMiddleware())
	examRouter.GET("/", delivery.GetExamById)   //done
	examRouter.PUT("/", delivery.UpdateExam)    //done
	examRouter.DELETE("/", delivery.DeleteExam) //done

	AdminSuperRoutes(examRouter, adminSuperDelivery, allDelivery)

	examRouter.GET("/admin", delivery.GetAdminsByExamId) //done
	examRouter.POST("/admin", delivery.CreateAdmin)      //done

	adminRouter := examRouter.Group("/admin/:admin_id", delivery.MustAdminMiddleware())
	adminRouter.GET("/", delivery.GetAdminById)   //done
	adminRouter.PUT("/", delivery.UpdateAdmin)    //done
	adminRouter.DELETE("/", delivery.DeleteAdmin) //done

}
