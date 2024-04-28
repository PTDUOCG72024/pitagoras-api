package projects

import (
	"github.com/gofiber/fiber/v2"
	"github.com/xbizzybone/pitagoras-api/utils"
)

type controller struct {
	cases Cases
}

func NewController(cases Cases) Controller {
	return &controller{cases}
}

// CreateProject godoc
// @Summary Create a new project
// @Description Create a new project
// @Tags projects
// @Accept json
// @Produce json
// @Param project body CreateProjectRequest true "Project object that needs to be created"
// @Success 200 {object} CreateProjectResponse
// @Failure 400 {string} string "Error validando el cuerpo de la petición"
// @Failure 500 {string} string "Error creando el proyecto"
// @Router /projects [post]
func (c *controller) CreateProject(ctx *fiber.Ctx) error {
	requestBody := new(CreateProjectRequest)

	if err := utils.BodyParser(ctx, requestBody); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error validando el cuerpo de la petición :" + err.Error(),
		})
	}

	project := new(Project)
	project.Name = requestBody.Name
	project.Description = requestBody.Description
	project.StartDate = requestBody.StartDate
	project.EndDate = requestBody.EndDate

	createProject, err := c.cases.CreateProject(ctx.UserContext(), project)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error creando el proyecto",
		})
	}

	projectResponse := new(CreateProjectResponse)
	projectResponse.ID = createProject.ID
	projectResponse.Name = createProject.Name
	projectResponse.Description = createProject.Description
	projectResponse.StartDate = createProject.StartDate
	projectResponse.EndDate = createProject.EndDate
	projectResponse.IsActive = createProject.IsActive
	projectResponse.CreatedAt = createProject.CreatedAt

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Proyecto creado correctamente",
		"project": projectResponse,
	})
}

// DeleteProject godoc
// @Summary Delete a project
// @Description Delete a project
// @Tags projects
// @Accept json
// @Produce json
// @Param id path string true "Project ID"
// @Success 200 {string} string "Proyecto eliminado correctamente"
// @Failure 500 {string} string "Error eliminando el proyecto"
// @Router /projects/{id} [delete]
func (c *controller) DeleteProject(ctx *fiber.Ctx) error {
	projectID := ctx.Params("id")

	err := c.cases.DeleteProject(ctx.UserContext(), projectID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error eliminando el proyecto",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Proyecto eliminado correctamente",
	})
}

// GetProjectById godoc
// @Summary Get a project by id
// @Description Get a project by id
// @Tags projects
// @Accept json
// @Produce json
// @Param id path string true "Project ID"
// @Success 200 {object} GetProjectResponse
// @Failure 500 {string} string "Error obteniendo el proyecto"
// @Router /projects/{id} [get]
func (c *controller) GetProjectById(ctx *fiber.Ctx) error {
	projectID := ctx.Params("id")

	project, err := c.cases.GetProjectById(ctx.UserContext(), projectID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error obteniendo el proyecto",
		})
	}

	projectResponse := new(GetProjectResponse)
	projectResponse.ID = project.ID
	projectResponse.Name = project.Name
	projectResponse.Description = project.Description
	projectResponse.StartDate = project.StartDate
	projectResponse.EndDate = project.EndDate
	projectResponse.IsActive = project.IsActive
	projectResponse.IsDeleted = project.IsDeleted
	projectResponse.CreatedAt = project.CreatedAt
	projectResponse.UpdatedAt = project.UpdatedAt
	projectResponse.DeleteAt = project.DeleteAt

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Proyecto obtenido correctamente",
		"project": projectResponse,
	})
}

// GetProjects godoc
// @Summary Get all projects
// @Description Get all projects
// @Tags projects
// @Accept json
// @Produce json
// @Success 200 {object} []GetProjectResponse
// @Failure 500 {string} string "Error obteniendo los proyectos"
// @Router /projects [get]
func (c *controller) GetProjects(ctx *fiber.Ctx) error {
	projects, err := c.cases.GetProjects(ctx.UserContext())
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error obteniendo los proyectos",
		})
	}

	projectsResponse := make([]GetProjectResponse, 0)
	for _, project := range projects {
		projectResponse := new(GetProjectResponse)
		projectResponse.ID = project.ID
		projectResponse.Name = project.Name
		projectResponse.Description = project.Description
		projectResponse.StartDate = project.StartDate
		projectResponse.EndDate = project.EndDate
		projectResponse.IsActive = project.IsActive
		projectResponse.IsDeleted = project.IsDeleted
		projectResponse.CreatedAt = project.CreatedAt
		projectResponse.UpdatedAt = project.UpdatedAt
		projectResponse.DeleteAt = project.DeleteAt

		projectsResponse = append(projectsResponse, *projectResponse)
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":  "Proyectos obtenidos correctamente",
		"projects": projectsResponse,
	})
}

// UpdateProject godoc
// @Summary Update a project
// @Description Update a project
// @Tags projects
// @Accept json
// @Produce json
// @Param id path string true "Project ID"
// @Param project body UpdateProjectRequest true "Project object that needs to be updated"
// @Success 200 {string} string "Proyecto actualizado correctamente"
// @Failure 400 {string} string "Error validando el cuerpo de la petición"
// @Failure 500 {string} string "Error actualizando el proyecto"
// @Router /projects/{id} [put]
func (c *controller) UpdateProject(ctx *fiber.Ctx) error {
	projectID := ctx.Params("id")

	requestBody := new(UpdateProjectRequest)

	if err := utils.BodyParser(ctx, requestBody); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error validando el cuerpo de la petición :" + err.Error(),
		})
	}

	project := new(Project)
	project.Name = requestBody.Name
	project.Description = requestBody.Description
	project.StartDate = requestBody.StartDate
	project.EndDate = requestBody.EndDate

	err := c.cases.UpdateProject(ctx.UserContext(), projectID, project)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error actualizando el proyecto",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Proyecto actualizado correctamente",
	})
}

// ActivateProject godoc
// @Summary Activate a project
// @Description Activate a project
// @Tags projects
// @Accept json
// @Produce json
// @Param id path string true "Project ID"
// @Success 200 {string} string "Proyecto activado correctamente"
// @Failure 500 {string} string "Error activando el proyecto"
// @Router /projects/activate/{id} [put]
func (c *controller) ActivateProject(ctx *fiber.Ctx) error {
	projectID := ctx.Params("id")

	err := c.cases.ActivateProject(ctx.UserContext(), projectID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error activando el proyecto",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Proyecto activado correctamente",
	})
}

// DeactivateProject godoc
// @Summary Deactivate a project
// @Description Deactivate a project
// @Tags projects
// @Accept json
// @Produce json
// @Param id path string true "Project ID"
// @Success 200 {string} string "Proyecto desactivado correctamente"
// @Failure 500 {string} string "Error desactivando el proyecto"
// @Router /projects/deactivate/{id} [put]
func (c *controller) DeactivateProject(ctx *fiber.Ctx) error {
	projectID := ctx.Params("id")

	err := c.cases.DeactivateProject(ctx.UserContext(), projectID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error desactivando el proyecto",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Proyecto desactivado correctamente",
	})
}
