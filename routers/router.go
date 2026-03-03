package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/sidz111/employee-mapping-gorm/controller"
)

func RegisterEmployeeRoutes(r *gin.Engine, empController *controller.EmployeeController) {
	emp := r.Group("/employees")
	{
		emp.POST("/", empController.Create)
		emp.GET("/:id", empController.GetByID)
		emp.GET("/", empController.GetAll)
		emp.PUT("/:id", empController.Update)
		emp.DELETE("/:id", empController.Delete)
	}
}

func RegisterProjectRoutes(r *gin.Engine, proController *controller.ProjectController) {
	pro := r.Group("/projects")
	{
		pro.POST("/", proController.Create)
		pro.GET("/:id", proController.GetByID)
		pro.GET("/", proController.GetAllProjects)
		pro.PUT("/:id", proController.Update)
		pro.DELETE("/:id", proController.Delete)
	}
}
