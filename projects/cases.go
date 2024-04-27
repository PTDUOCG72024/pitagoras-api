package projects

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

type cases struct {
	logger     *zap.Logger
	repository Repository
}

func NewCases(logger *zap.Logger, repository Repository) Cases {
	return &cases{
		logger:     logger,
		repository: repository,
	}
}

func (c *cases) CreateProject(ctx context.Context, project *Project) (Project, error) {
	project.CreatedAt = time.Now()
	project.UpdatedAt = time.Now()
	project.IsActive = true

	if err := c.repository.CreateProject(ctx, project); err != nil {
		return Project{}, err
	}

	return *project, nil
}

func (c *cases) DeleteProject(ctx context.Context, id string) error {
	project, err := c.repository.GetProjectById(ctx, id)

	if err != nil {
		return err
	}

	project.DeleteAt = time.Now()
	project.UpdatedAt = time.Now()
	project.IsDeleted = true

	if err := c.repository.UpdateProject(ctx, &project); err != nil {
		return err
	}

	return nil
}

func (c *cases) GetProjectById(ctx context.Context, id string) (Project, error) {
	project, err := c.repository.GetProjectById(ctx, id)

	if err != nil {
		return Project{}, err
	}

	return project, nil
}

func (c *cases) GetProjects(ctx context.Context) ([]Project, error) {
	projects, err := c.repository.GetProjects(ctx)

	if err != nil {
		return []Project{}, err
	}

	return projects, nil
}

func (c *cases) UpdateProject(ctx context.Context, id string, project *Project) error {
	_, err := c.repository.GetProjectById(ctx, id)

	if err != nil {
		return err
	}

	project.ID, _ = primitive.ObjectIDFromHex(id)
	project.UpdatedAt = time.Now()

	if err := c.repository.UpdateProject(ctx, project); err != nil {
		return err
	}

	return nil
}

func (c *cases) ActivateProject(ctx context.Context, id string) error {
	project, err := c.repository.GetProjectById(ctx, id)

	if err != nil {
		return err
	}

	project.IsActive = true
	project.UpdatedAt = time.Now()

	if err := c.repository.UpdateProject(ctx, &project); err != nil {
		return err
	}

	return nil
}

func (c *cases) DeactivateProject(ctx context.Context, id string) error {
	project, err := c.repository.GetProjectById(ctx, id)

	if err != nil {
		return err
	}

	project.IsActive = false
	project.UpdatedAt = time.Now()

	if err := c.repository.UpdateProject(ctx, &project); err != nil {
		return err
	}

	return nil
}
