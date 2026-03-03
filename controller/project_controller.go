package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sidz111/employee-mapping-gorm/model"
	"github.com/sidz111/employee-mapping-gorm/service"
)

type ProjectController struct {
	service service.ProjectService
}

func NewProjectController(service service.ProjectService) *ProjectController {
	return &ProjectController{service: service}
}

func (c *ProjectController) Create(ctx *gin.Context) {
	var project model.Project
	if err := ctx.ShouldBindJSON(&project); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := c.service.Create(ctx.Request.Context(), &project); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(201, project)
}
func (c *ProjectController) GetByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}
	project, err := c.service.GetByID(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(404, gin.H{"error": "Project not found"})
		return
	}
	ctx.JSON(200, project)
}
func (c *ProjectController) GetAllProjects(ctx *gin.Context) {
	projects, err := c.service.GetAllProjects(ctx.Request.Context())
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, projects)
}
func (c *ProjectController) Update(ctx *gin.Context) {
	var project model.Project
	if err := ctx.ShouldBindJSON(&project); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := c.service.Update(ctx.Request.Context(), &project); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, project)
}
func (c *ProjectController) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}
	if err := c.service.Delete(ctx.Request.Context(), id); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"message": "Project deleted successfully"})
}
