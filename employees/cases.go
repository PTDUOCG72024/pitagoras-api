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
func (c *cases) ActivateEmployee(ctx context.Context, id string) error {
	employee, err := c.repository.GetEmployeeById(ctx, id)

	if err != nil {
		return err
	}

	employee.IsActive = true
	employee.UpdatedAt = time.Now()

	if err := c.repository.UpdateEmployee(ctx, employee); err != nil {
		return err
	}

	return nil
}

// ActivateNationality implements Cases.
func (c *cases) ActivateNationality(ctx context.Context, id string) error {
	nationality, err := c.repository.GetNationalityById(ctx, id)

	if err != nil {
		return err
	}

	nationality.IsActive = true
	nationality.UpdatedAt = time.Now()

	if err := c.repository.UpdateNationality(ctx, nationality); err != nil {
		return err
	}

	return nil
}

// ActivatePosition implements Cases.
func (c *cases) ActivatePosition(ctx context.Context, id string) error {
	position, err := c.repository.GetPositionById(ctx, id)

	if err != nil {
		return err
	}

	position.IsActive = true
	position.UpdatedAt = time.Now()

	if err := c.repository.UpdatePosition(ctx, position); err != nil {
		return err
	}

	return nil
}

// ActivateSupervisor implements Cases.
func (c *cases) ActivateSupervisor(ctx context.Context, id string) error {
	supervisor, err := c.repository.GetSupervisorById(ctx, id)

	if err != nil {
		return err
	}

	supervisor.IsActive = true
	supervisor.UpdatedAt = time.Now()

	if err := c.repository.UpdateSupervisor(ctx, supervisor); err != nil {
		return err
	}

	return nil
}

// CreateEmployee implements Cases.
func (c *cases) CreateEmployee(ctx context.Context, employee *Employee) (*Employee, error) {
	_, err := c.repository.GetEmployeeByIdentificationNumber(ctx, employee.IdentificationNumber)

	if err != nil {
		return nil, err
	}

	supervisor, err := c.repository.GetSupervisorById(ctx, employee.Supervisor.ID.Hex())

	if err != nil {
		return nil, err
	}

	employee.Supervisor.Name = supervisor.Name
	employee.Supervisor.IsActive = supervisor.IsActive
	employee.Supervisor.IsDeleted = supervisor.IsDeleted
	employee.Supervisor.CreatedAt = supervisor.CreatedAt
	employee.Supervisor.UpdatedAt = supervisor.UpdatedAt
	employee.Supervisor.DeleteAt = supervisor.DeleteAt

	position, err := c.repository.GetPositionById(ctx, employee.Position.ID.Hex())

	if err != nil {
		return nil, err
	}

	employee.Position.Name = position.Name
	employee.Position.IsActive = position.IsActive
	employee.Position.IsDeleted = position.IsDeleted
	employee.Position.CreatedAt = position.CreatedAt
	employee.Position.UpdatedAt = position.UpdatedAt
	employee.Position.DeleteAt = position.DeleteAt

	nationality, err := c.repository.GetNationalityById(ctx, employee.Nationality.ID.Hex())

	if err != nil {
		return nil, err
	}

	employee.Nationality.Name = nationality.Name
	employee.Nationality.IsActive = nationality.IsActive
	employee.Nationality.IsDeleted = nationality.IsDeleted
	employee.Nationality.CreatedAt = nationality.CreatedAt
	employee.Nationality.UpdatedAt = nationality.UpdatedAt
	employee.Nationality.DeleteAt = nationality.DeleteAt

	employee.CreatedAt = time.Now()
	employee.UpdatedAt = time.Now()
	employee.IsActive = true

	result, err := c.repository.CreateEmployee(ctx, employee)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// CreateNationality implements Cases.
func (c *cases) CreateNationality(ctx context.Context, nationality *Nationality) (*Nationality, error) {
	_, err := c.repository.GetNationalityByName(ctx, nationality.Name)

	if err != nil {
		return nil, err
	}

	nationality.CreatedAt = time.Now()
	nationality.UpdatedAt = time.Now()
	nationality.IsActive = true

	result, err := c.repository.CreateNationality(ctx, nationality)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// CreatePosition implements Cases.
func (c *cases) CreatePosition(ctx context.Context, position *Position) (*Position, error) {
	_, err := c.repository.GetPositionByName(ctx, position.Name)

	if err != nil {
		return nil, err
	}

	position.CreatedAt = time.Now()
	position.UpdatedAt = time.Now()
	position.IsActive = true

	result, err := c.repository.CreatePosition(ctx, position)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// CreateSupervisor implements Cases.
func (c *cases) CreateSupervisor(ctx context.Context, supervisor *Supervisor) (*Supervisor, error) {
	_, err := c.repository.GetSupervisorByName(ctx, supervisor.Name)

	if err != nil {
		return nil, err
	}

	supervisor.CreatedAt = time.Now()
	supervisor.UpdatedAt = time.Now()
	supervisor.IsActive = true

	result, err := c.repository.CreateSupervisor(ctx, supervisor)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// DeactivateEmployee implements Cases.
func (c *cases) DeactivateEmployee(ctx context.Context, id string) error {
	employee, err := c.repository.GetEmployeeById(ctx, id)

	if err != nil {
		return err
	}

	employee.IsActive = false
	employee.UpdatedAt = time.Now()

	if err := c.repository.UpdateEmployee(ctx, employee); err != nil {
		return err
	}

	return nil
}

// DeactivateNationality implements Cases.
func (c *cases) DeactivateNationality(ctx context.Context, id string) error {
	nationality, err := c.repository.GetNationalityById(ctx, id)

	if err != nil {
		return err
	}

	nationality.IsActive = false
	nationality.UpdatedAt = time.Now()

	if err := c.repository.UpdateNationality(ctx, nationality); err != nil {
		return err
	}

	return nil
}

// DeactivatePosition implements Cases.
func (c *cases) DeactivatePosition(ctx context.Context, id string) error {
	position, err := c.repository.GetPositionById(ctx, id)

	if err != nil {
		return err
	}

	position.IsActive = false
	position.UpdatedAt = time.Now()

	if err := c.repository.UpdatePosition(ctx, position); err != nil {
		return err
	}

	return nil
}

// DeactivateSupervisor implements Cases.
func (c *cases) DeactivateSupervisor(ctx context.Context, id string) error {
	supervisor, err := c.repository.GetSupervisorById(ctx, id)

	if err != nil {
		return err
	}

	supervisor.IsActive = false
	supervisor.UpdatedAt = time.Now()

	if err := c.repository.UpdateSupervisor(ctx, supervisor); err != nil {
		return err
	}

	return nil
}

// DeleteEmployee implements Cases.
func (c *cases) DeleteEmployee(ctx context.Context, id string) error {
	employee, err := c.repository.GetEmployeeById(ctx, id)

	if err != nil {
		return err
	}

	employee.IsDeleted = true
	employee.UpdatedAt = time.Now()
	employee.DeleteAt = time.Now()

	if err := c.repository.UpdateEmployee(ctx, employee); err != nil {
		return err
	}

	return nil
}

// DeleteNationality implements Cases.
func (c *cases) DeleteNationality(ctx context.Context, id string) error {
	nationality, err := c.repository.GetNationalityById(ctx, id)

	if err != nil {
		return err
	}

	nationality.IsDeleted = true
	nationality.UpdatedAt = time.Now()
	nationality.DeleteAt = time.Now()

	if err := c.repository.UpdateNationality(ctx, nationality); err != nil {
		return err
	}

	return nil
}

// DeletePosition implements Cases.
func (c *cases) DeletePosition(ctx context.Context, id string) error {
	position, err := c.repository.GetPositionById(ctx, id)

	if err != nil {
		return err
	}

	position.IsDeleted = true
	position.UpdatedAt = time.Now()
	position.DeleteAt = time.Now()

	if err := c.repository.UpdatePosition(ctx, position); err != nil {
		return err
	}

	return nil
}

// DeleteSupervisor implements Cases.
func (c *cases) DeleteSupervisor(ctx context.Context, id string) error {
	supervisor, err := c.repository.GetSupervisorById(ctx, id)

	if err != nil {
		return err
	}

	supervisor.IsDeleted = true
	supervisor.UpdatedAt = time.Now()
	supervisor.DeleteAt = time.Now()

	if err := c.repository.UpdateSupervisor(ctx, supervisor); err != nil {
		return err
	}

	return nil
}

// GetEmployeeById implements Cases.
func (c *cases) GetEmployeeById(ctx context.Context, id string) (*Employee, error) {
	employee, err := c.repository.GetEmployeeById(ctx, id)

	if err != nil {
		return nil, err
	}

	return employee, nil
}

// GetEmployees implements Cases.
func (c *cases) GetEmployees(ctx context.Context) ([]Employee, error) {
	employees, err := c.repository.GetEmployees(ctx)

	if err != nil {
		return nil, err
	}

	return employees, nil
}

// GetNationalities implements Cases.
func (c *cases) GetNationalities(ctx context.Context) ([]Nationality, error) {
	nationalities, err := c.repository.GetNationalities(ctx)

	if err != nil {
		return nil, err
	}

	return nationalities, nil
}

// GetNationalityById implements Cases.
func (c *cases) GetNationalityById(ctx context.Context, id string) (*Nationality, error) {
	nationality, err := c.repository.GetNationalityById(ctx, id)

	if err != nil {
		return nil, err
	}

	return nationality, nil
}

// GetPositionById implements Cases.
func (c *cases) GetPositionById(ctx context.Context, id string) (*Position, error) {
	position, err := c.repository.GetPositionById(ctx, id)

	if err != nil {
		return nil, err
	}

	return position, nil
}

// GetPositions implements Cases.
func (c *cases) GetPositions(ctx context.Context) ([]Position, error) {
	positions, err := c.repository.GetPositions(ctx)

	if err != nil {
		return nil, err
	}

	return positions, nil
}

// GetSupervisorById implements Cases.
func (c *cases) GetSupervisorById(ctx context.Context, id string) (*Supervisor, error) {
	supervisor, err := c.repository.GetSupervisorById(ctx, id)

	if err != nil {
		return nil, err
	}

	return supervisor, nil
}

// GetSupervisors implements Cases.
func (c *cases) GetSupervisors(ctx context.Context) ([]Supervisor, error) {
	supervisors, err := c.repository.GetSupervisors(ctx)

	if err != nil {
		return nil, err
	}

	return supervisors, nil
}

// UpdateEmployee implements Cases.
func (c *cases) UpdateEmployee(ctx context.Context, id string, employee *Employee) error {
	result, err := c.repository.GetEmployeeById(ctx, id)

	if err != nil {
		return err
	}

	_, err = c.repository.GetSupervisorById(ctx, employee.Supervisor.ID.Hex())

	if err != nil {
		return err
	}

	_, err = c.repository.GetPositionById(ctx, employee.Position.ID.Hex())

	if err != nil {
		return err
	}

	_, err = c.repository.GetNationalityById(ctx, employee.Nationality.ID.Hex())

	if err != nil {
		return err
	}

	if employee.Name != result.Name {
		result.Name = employee.Name
	}

	if employee.Email != result.Email {
		result.Email = employee.Email
	}

	if employee.InLaw != result.InLaw {
		result.InLaw = employee.InLaw
	}

	if employee.ContractType != result.ContractType {
		result.ContractType = employee.ContractType
	}

	if employee.StartDate != result.StartDate {
		result.StartDate = employee.StartDate
	}

	if employee.EndDate != result.EndDate {
		result.EndDate = employee.EndDate
	}

	employee.UpdatedAt = time.Now()

	if employee.Supervisor.Name != result.Supervisor.Name {
		result.Supervisor.Name = employee.Supervisor.Name
	}

	employee.Supervisor.UpdatedAt = time.Now()

	if employee.Position.Name != result.Position.Name {
		result.Position.Name = employee.Position.Name
	}

	employee.Position.UpdatedAt = time.Now()

	if employee.Nationality.Name != result.Nationality.Name {
		result.Nationality.Name = employee.Nationality.Name
	}

	employee.Nationality.UpdatedAt = time.Now()

	if err := c.repository.UpdateEmployee(ctx, result); err != nil {
		return err
	}

	return nil
}

// UpdateNationality implements Cases.
func (c *cases) UpdateNationality(ctx context.Context, id string, nationality *Nationality) error {
	result, err := c.repository.GetNationalityById(ctx, id)

	if err != nil {
		return err
	}

	if nationality.Name != result.Name {
		result.Name = nationality.Name
	}

	result.UpdatedAt = time.Now()

	if err := c.repository.UpdateNationality(ctx, result); err != nil {
		return err
	}

	return nil
}

// UpdatePosition implements Cases.
func (c *cases) UpdatePosition(ctx context.Context, id string, position *Position) error {
	result, err := c.repository.GetPositionById(ctx, id)

	if err != nil {
		return err
	}

	if position.Name != result.Name {
		result.Name = position.Name
	}

	result.UpdatedAt = time.Now()

	if err := c.repository.UpdatePosition(ctx, result); err != nil {
		return err
	}

	return nil
}

// UpdateSupervisor implements Cases.
func (c *cases) UpdateSupervisor(ctx context.Context, id string, supervisor *Supervisor) error {
	result, err := c.repository.GetSupervisorById(ctx, id)

	if err != nil {
		return err
	}

	if supervisor.Name != result.Name {
		result.Name = supervisor.Name
	}

	result.UpdatedAt = time.Now()

	if err := c.repository.UpdateSupervisor(ctx, result); err != nil {
		return err
	}

	return nil
}
