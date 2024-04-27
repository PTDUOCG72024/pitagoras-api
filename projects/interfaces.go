package projects

import (
	"context"

	"github.com/gofiber/fiber/v2"
)

type Controller interface {
	CreateProject(ctx *fiber.Ctx) error
	GetProjectById(ctx *fiber.Ctx) error
	GetProjects(ctx *fiber.Ctx) error
	UpdateProject(ctx *fiber.Ctx) error
	DeleteProject(ctx *fiber.Ctx) error
	ActivateProject(ctx *fiber.Ctx) error
	DeactivateProject(ctx *fiber.Ctx) error
}

type Cases interface {
	CreateProject(ctx context.Context, project *Project) (Project, error)
	GetProjectById(ctx context.Context, id string) (Project, error)
	GetProjects(ctx context.Context) ([]Project, error)
	UpdateProject(ctx context.Context, id string, project *Project) error
	DeleteProject(ctx context.Context, id string) error
	ActivateProject(ctx context.Context, id string) error
	DeactivateProject(ctx context.Context, id string) error
}

type Repository interface {
	CreateProject(ctx context.Context, project *Project) error
	UpdateProject(ctx context.Context, project *Project) error
	GetProjectById(ctx context.Context, id string) (Project, error)
	GetProjects(ctx context.Context) ([]Project, error)
}
