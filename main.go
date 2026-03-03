package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sidz111/employee-mapping-gorm/config"
	"github.com/sidz111/employee-mapping-gorm/controller"
	"github.com/sidz111/employee-mapping-gorm/model"
	"github.com/sidz111/employee-mapping-gorm/repository"
	"github.com/sidz111/employee-mapping-gorm/service"
)

func main() {
	if err := config.ConnectDB(); err != nil {
		panic(err)
	}
	config.DB.AutoMigrate(&model.Employee{}, &model.Project{})

	empRepo := repository.NewEmployeeRepository(config.DB)
	empServ := service.NewEmployeeService(empRepo)
	empController := controller.NewEmployeeController(empServ)

	proRepo := repository.NewProjectRepository(config.DB)
	proServ := service.NewProjectService(proRepo)
	proController := controller.NewProjectController(proServ)

	router := gin.Default()
	router.POST("/employees", empController.Create)
	router.GET("/employees/:id", empController.GetByID)
	router.GET("/employees", empController.GetAll)
	router.PUT("/employees", empController.Update)
	router.DELETE("/employees/:id", empController.Delete)

	router.POST("/projects", proController.Create)
	router.GET("/projects/:id", proController.GetByID)
	router.GET("/projects", proController.GetAllProjects)
	router.PUT("/projects", proController.Update)
	router.DELETE("/projects/:id", proController.Delete)

	router.Run(":8080")
}
