package employees

import (
	"context"
	"time"

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

// ActivateEmployee implements Cases.
func (c *cases) ActivateEmployee(ctx context.Context, id string) (*Employee, error) {
	employee, err := c.repository.GetEmployeeById(ctx, id)

	if err != nil {
		return nil, err
	}

	employee.IsActive = true
	employee.UpdatedAt = time.Now()

	if _, err := c.repository.UpdateEmployee(ctx, employee); err != nil {
		return nil, err
	}

	return employee, nil
}

// ActivateNationality implements Cases.
func (c *cases) ActivateNationality(ctx context.Context, id string) (*Nationality, error) {
	nationality, err := c.repository.GetNationalityById(ctx, id)

	if err != nil {
		return nil, err
	}

	nationality.IsActive = true
	nationality.UpdatedAt = time.Now()

	if _, err := c.repository.UpdateNationality(ctx, nationality); err != nil {
		return nil, err
	}

	return nationality, nil
}

// ActivatePosition implements Cases.
func (c *cases) ActivatePosition(ctx context.Context, id string) (*Position, error) {
	position, err := c.repository.GetPositionById(ctx, id)

	if err != nil {
		return nil, err
	}

	position.IsActive = true
	position.UpdatedAt = time.Now()

	if _, err := c.repository.UpdatePosition(ctx, position); err != nil {
		return nil, err
	}

	return position, nil
}

// ActivateSupervisor implements Cases.
func (c *cases) ActivateSupervisor(ctx context.Context, id string) (*Supervisor, error) {
	supervisor, err := c.repository.GetSupervisorById(ctx, id)

	if err != nil {
		return nil, err
	}

	supervisor.IsActive = true
	supervisor.UpdatedAt = time.Now()

	if _, err := c.repository.UpdateSupervisor(ctx, supervisor); err != nil {
		return nil, err
	}

	return supervisor, nil
}

// CreateEmployee implements Cases.
func (c *cases) CreateEmployee(ctx context.Context, employee *Employee) (*Employee, error) {
	panic("unimplemented")
}

// CreateNationality implements Cases.
func (c *cases) CreateNationality(ctx context.Context, nationality *Nationality) (*Nationality, error) {
	panic("unimplemented")
}

// CreatePosition implements Cases.
func (c *cases) CreatePosition(ctx context.Context, position *Position) (*Position, error) {
	panic("unimplemented")
}

// CreateSupervisor implements Cases.
func (c *cases) CreateSupervisor(ctx context.Context, supervisor *Supervisor) (*Supervisor, error) {
	panic("unimplemented")
}

// DeactivateEmployee implements Cases.
func (c *cases) DeactivateEmployee(ctx context.Context, id string) (*Employee, error) {
	panic("unimplemented")
}

// DeactivateNationality implements Cases.
func (c *cases) DeactivateNationality(ctx context.Context, id string) (*Nationality, error) {
	panic("unimplemented")
}

// DeactivatePosition implements Cases.
func (c *cases) DeactivatePosition(ctx context.Context, id string) (*Position, error) {
	panic("unimplemented")
}

// DeactivateSupervisor implements Cases.
func (c *cases) DeactivateSupervisor(ctx context.Context, id string) (*Supervisor, error) {
	panic("unimplemented")
}

// DeleteEmployee implements Cases.
func (c *cases) DeleteEmployee(ctx context.Context, id string) error {
	panic("unimplemented")
}

// DeleteNationality implements Cases.
func (c *cases) DeleteNationality(ctx context.Context, id string) error {
	panic("unimplemented")
}

// DeletePosition implements Cases.
func (c *cases) DeletePosition(ctx context.Context, id string) error {
	panic("unimplemented")
}

// DeleteSupervisor implements Cases.
func (c *cases) DeleteSupervisor(ctx context.Context, id string) error {
	panic("unimplemented")
}

// GetEmployeeById implements Cases.
func (c *cases) GetEmployeeById(ctx context.Context, id string) (*Employee, error) {
	panic("unimplemented")
}

// GetEmployees implements Cases.
func (c *cases) GetEmployees(ctx context.Context) ([]Employee, error) {
	panic("unimplemented")
}

// GetNationalities implements Cases.
func (c *cases) GetNationalities(ctx context.Context) ([]Nationality, error) {
	panic("unimplemented")
}

// GetNationalityById implements Cases.
func (c *cases) GetNationalityById(ctx context.Context, id string) (*Nationality, error) {
	panic("unimplemented")
}

// GetPositionById implements Cases.
func (c *cases) GetPositionById(ctx context.Context, id string) (*Position, error) {
	panic("unimplemented")
}

// GetPositions implements Cases.
func (c *cases) GetPositions(ctx context.Context) ([]Position, error) {
	panic("unimplemented")
}

// GetSupervisorById implements Cases.
func (c *cases) GetSupervisorById(ctx context.Context, id string) (*Supervisor, error) {
	panic("unimplemented")
}

// GetSupervisors implements Cases.
func (c *cases) GetSupervisors(ctx context.Context) ([]Supervisor, error) {
	panic("unimplemented")
}

// UpdateEmployee implements Cases.
func (c *cases) UpdateEmployee(ctx context.Context, id string, employee *Employee) (*Employee, error) {
	panic("unimplemented")
}

// UpdateNationality implements Cases.
func (c *cases) UpdateNationality(ctx context.Context, id string, nationality *Nationality) (*Nationality, error) {
	panic("unimplemented")
}

// UpdatePosition implements Cases.
func (c *cases) UpdatePosition(ctx context.Context, id string, position *Position) (*Position, error) {
	panic("unimplemented")
}

// UpdateSupervisor implements Cases.
func (c *cases) UpdateSupervisor(ctx context.Context, id string, supervisor *Supervisor) (*Supervisor, error) {
	panic("unimplemented")
}
