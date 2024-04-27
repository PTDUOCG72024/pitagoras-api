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
	projectResponse.Name = createProject.Name
	projectResponse.Description = createProject.Description
	projectResponse.StartDate = createProject.StartDate
	projectResponse.EndDate = createProject.EndDate
	projectResponse.IsActive = createProject.IsActive

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Proyecto creado correctamente",
		"project": projectResponse,
	})
}

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

func (c *controller) GetProjectById(ctx *fiber.Ctx) error {
	projectID := ctx.Params("id")

	project, err := c.cases.GetProjectById(ctx.UserContext(), projectID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error obteniendo el proyecto",
		})
	}

	projectResponse := new(GetProjectResponse)
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
	project.IsActive = requestBody.IsActive

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
