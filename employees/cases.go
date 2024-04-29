package employees

import (
	"context"
	"errors"
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

// ActivateEmployee implements Cases.
func (c *cases) ActivateEmployee(ctx context.Context, id string) error {
	objID, _ := primitive.ObjectIDFromHex(id)
	employee, err := c.repository.GetEmployeeById(ctx, objID)

	if err != nil {
		return err
	}

	employee.IsActive = true
	employee.UpdatedAt = time.Now()

	return c.repository.UpdateEmployee(ctx, employee)
}

// ActivateNationality implements Cases.
func (c *cases) ActivateNationality(ctx context.Context, id string) error {
	objID, _ := primitive.ObjectIDFromHex(id)
	nationality, err := c.repository.GetNationalityById(ctx, objID)

	if err != nil {
		return err
	}

	nationality.IsActive = true
	nationality.UpdatedAt = time.Now()

	return c.repository.UpdateNationality(ctx, nationality)
}

// ActivatePosition implements Cases.
func (c *cases) ActivatePosition(ctx context.Context, id string) error {
	objID, _ := primitive.ObjectIDFromHex(id)
	position, err := c.repository.GetPositionById(ctx, objID)

	if err != nil {
		return err
	}

	position.IsActive = true
	position.UpdatedAt = time.Now()

	return c.repository.UpdatePosition(ctx, position)
}

// ActivateSupervisor implements Cases.
func (c *cases) ActivateSupervisor(ctx context.Context, id string) error {
	objID, _ := primitive.ObjectIDFromHex(id)
	supervisor, err := c.repository.GetSupervisorById(ctx, objID)

	if err != nil {
		return err
	}

	supervisor.IsActive = true
	supervisor.UpdatedAt = time.Now()

	return c.repository.UpdateSupervisor(ctx, supervisor)
}

// CreateEmployee implements Cases.
func (c *cases) CreateEmployee(ctx context.Context, employee *Employee) (*Employee, error) {
	emp, _ := c.repository.GetEmployeeByIdentificationNumber(ctx, employee.IdentificationNumber)

	if emp != nil {
		return nil, errors.New("el empleado ya existe")
	}

	supervisor, err := c.repository.GetSupervisorById(ctx, employee.Supervisor.ID)

	if err != nil {
		return nil, errors.New("el supervisor no existe")
	}

	employee.Supervisor.Name = supervisor.Name
	employee.Supervisor.IsActive = supervisor.IsActive
	employee.Supervisor.IsDeleted = supervisor.IsDeleted
	employee.Supervisor.CreatedAt = supervisor.CreatedAt
	employee.Supervisor.UpdatedAt = supervisor.UpdatedAt
	employee.Supervisor.DeleteAt = supervisor.DeleteAt

	position, err := c.repository.GetPositionById(ctx, employee.Position.ID)

	if err != nil {
		return nil, errors.New("el cargo no existe")
	}

	employee.Position.Name = position.Name
	employee.Position.IsActive = position.IsActive
	employee.Position.IsDeleted = position.IsDeleted
	employee.Position.CreatedAt = position.CreatedAt
	employee.Position.UpdatedAt = position.UpdatedAt
	employee.Position.DeleteAt = position.DeleteAt

	nationality, err := c.repository.GetNationalityById(ctx, employee.Nationality.ID)

	if err != nil {
		return nil, errors.New("la nacionalidad no existe")
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

	return c.repository.CreateEmployee(ctx, employee)
}

// CreateNationality implements Cases.
func (c *cases) CreateNationality(ctx context.Context, nationality *Nationality) (*Nationality, error) {
	nat, _ := c.repository.GetNationalityByName(ctx, nationality.Name)

	if nat != nil {
		return nil, errors.New("la nacionalidad ya existe")
	}

	nationality.CreatedAt = time.Now()
	nationality.UpdatedAt = time.Now()
	nationality.IsActive = true

	return c.repository.CreateNationality(ctx, nationality)
}

// CreatePosition implements Cases.
func (c *cases) CreatePosition(ctx context.Context, position *Position) (*Position, error) {
	pos, err := c.repository.GetPositionByName(ctx, position.Name)

	if pos != nil {
		return nil, err
	}

	position.CreatedAt = time.Now()
	position.UpdatedAt = time.Now()
	position.IsActive = true

	return c.repository.CreatePosition(ctx, position)
}

// CreateSupervisor implements Cases.
func (c *cases) CreateSupervisor(ctx context.Context, supervisor *Supervisor) (*Supervisor, error) {
	sup, _ := c.repository.GetSupervisorByName(ctx, supervisor.Name)

	if sup != nil {
		return nil, errors.New("el supervisor ya existe")
	}

	supervisor.CreatedAt = time.Now()
	supervisor.UpdatedAt = time.Now()
	supervisor.IsActive = true

	return c.repository.CreateSupervisor(ctx, supervisor)
}

// DeactivateEmployee implements Cases.
func (c *cases) DeactivateEmployee(ctx context.Context, id string) error {
	objID, _ := primitive.ObjectIDFromHex(id)
	employee, err := c.repository.GetEmployeeById(ctx, objID)

	if err != nil {
		return err
	}

	employee.IsActive = false
	employee.UpdatedAt = time.Now()

	return c.repository.UpdateEmployee(ctx, employee)
}

// DeactivateNationality implements Cases.
func (c *cases) DeactivateNationality(ctx context.Context, id string) error {
	objID, _ := primitive.ObjectIDFromHex(id)
	nationality, err := c.repository.GetNationalityById(ctx, objID)

	if err != nil {
		return err
	}

	nationality.IsActive = false
	nationality.UpdatedAt = time.Now()

	return c.repository.UpdateNationality(ctx, nationality)
}

// DeactivatePosition implements Cases.
func (c *cases) DeactivatePosition(ctx context.Context, id string) error {
	objID, _ := primitive.ObjectIDFromHex(id)
	position, err := c.repository.GetPositionById(ctx, objID)

	if err != nil {
		return err
	}

	position.IsActive = false
	position.UpdatedAt = time.Now()

	return c.repository.UpdatePosition(ctx, position)
}

// DeactivateSupervisor implements Cases.
func (c *cases) DeactivateSupervisor(ctx context.Context, id string) error {
	objID, _ := primitive.ObjectIDFromHex(id)
	supervisor, err := c.repository.GetSupervisorById(ctx, objID)

	if err != nil {
		return err
	}

	supervisor.IsActive = false
	supervisor.UpdatedAt = time.Now()

	return c.repository.UpdateSupervisor(ctx, supervisor)
}

// DeleteEmployee implements Cases.
func (c *cases) DeleteEmployee(ctx context.Context, id string) error {
	objID, _ := primitive.ObjectIDFromHex(id)
	employee, err := c.repository.GetEmployeeById(ctx, objID)

	if err != nil {
		return err
	}

	employee.IsDeleted = true
	employee.UpdatedAt = time.Now()
	employee.DeleteAt = time.Now()

	return c.repository.UpdateEmployee(ctx, employee)
}

// DeleteNationality implements Cases.
func (c *cases) DeleteNationality(ctx context.Context, id string) error {
	objID, _ := primitive.ObjectIDFromHex(id)
	nationality, err := c.repository.GetNationalityById(ctx, objID)

	if err != nil {
		return err
	}

	nationality.IsDeleted = true
	nationality.UpdatedAt = time.Now()
	nationality.DeleteAt = time.Now()

	return c.repository.UpdateNationality(ctx, nationality)
}

// DeletePosition implements Cases.
func (c *cases) DeletePosition(ctx context.Context, id string) error {
	objID, _ := primitive.ObjectIDFromHex(id)
	position, err := c.repository.GetPositionById(ctx, objID)

	if err != nil {
		return err
	}

	position.IsDeleted = true
	position.UpdatedAt = time.Now()
	position.DeleteAt = time.Now()

	return c.repository.UpdatePosition(ctx, position)
}

// DeleteSupervisor implements Cases.
func (c *cases) DeleteSupervisor(ctx context.Context, id string) error {
	objID, _ := primitive.ObjectIDFromHex(id)
	supervisor, err := c.repository.GetSupervisorById(ctx, objID)

	if err != nil {
		return err
	}

	supervisor.IsDeleted = true
	supervisor.UpdatedAt = time.Now()
	supervisor.DeleteAt = time.Now()

	return c.repository.UpdateSupervisor(ctx, supervisor)
}

// GetEmployeeById implements Cases.
func (c *cases) GetEmployeeById(ctx context.Context, id string) (*Employee, error) {
	objID, _ := primitive.ObjectIDFromHex(id)

	return c.repository.GetEmployeeById(ctx, objID)
}

// GetEmployees implements Cases.
func (c *cases) GetEmployees(ctx context.Context) ([]Employee, error) {
	return c.repository.GetEmployees(ctx)
}

// GetNationalities implements Cases.
func (c *cases) GetNationalities(ctx context.Context) ([]Nationality, error) {
	return c.repository.GetNationalities(ctx)
}

// GetNationalityById implements Cases.
func (c *cases) GetNationalityById(ctx context.Context, id string) (*Nationality, error) {
	objID, _ := primitive.ObjectIDFromHex(id)

	return c.repository.GetNationalityById(ctx, objID)
}

// GetPositionById implements Cases.
func (c *cases) GetPositionById(ctx context.Context, id string) (*Position, error) {
	objID, _ := primitive.ObjectIDFromHex(id)

	return c.repository.GetPositionById(ctx, objID)
}

// GetPositions implements Cases.
func (c *cases) GetPositions(ctx context.Context) ([]Position, error) {
	return c.repository.GetPositions(ctx)
}

// GetSupervisorById implements Cases.
func (c *cases) GetSupervisorById(ctx context.Context, id string) (*Supervisor, error) {
	objID, _ := primitive.ObjectIDFromHex(id)

	return c.repository.GetSupervisorById(ctx, objID)
}

// GetSupervisors implements Cases.
func (c *cases) GetSupervisors(ctx context.Context) ([]Supervisor, error) {
	return c.repository.GetSupervisors(ctx)
}

// UpdateEmployee implements Cases.
func (c *cases) UpdateEmployee(ctx context.Context, id string, employee *Employee) error {
	objID, _ := primitive.ObjectIDFromHex(id)
	result, err := c.repository.GetEmployeeById(ctx, objID)

	if err != nil {
		return err
	}

	if _, err = c.repository.GetSupervisorById(ctx, employee.Supervisor.ID); err != nil {
		return err
	}

	if _, err = c.repository.GetPositionById(ctx, employee.Position.ID); err != nil {
		return err
	}

	if _, err = c.repository.GetNationalityById(ctx, employee.Nationality.ID); err != nil {
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

	return c.repository.UpdateEmployee(ctx, result)
}

// UpdateNationality implements Cases.
func (c *cases) UpdateNationality(ctx context.Context, id string, nationality *Nationality) error {
	objID, _ := primitive.ObjectIDFromHex(id)
	result, err := c.repository.GetNationalityById(ctx, objID)

	if err != nil {
		return err
	}

	if nationality.Name != result.Name {
		result.Name = nationality.Name
	}

	result.UpdatedAt = time.Now()

	return c.repository.UpdateNationality(ctx, result)
}

// UpdatePosition implements Cases.
func (c *cases) UpdatePosition(ctx context.Context, id string, position *Position) error {
	objID, _ := primitive.ObjectIDFromHex(id)
	result, err := c.repository.GetPositionById(ctx, objID)

	if err != nil {
		return err
	}

	if position.Name != result.Name {
		result.Name = position.Name
	}

	result.UpdatedAt = time.Now()

	return c.repository.UpdatePosition(ctx, result)
}

// UpdateSupervisor implements Cases.
func (c *cases) UpdateSupervisor(ctx context.Context, id string, supervisor *Supervisor) error {
	objID, _ := primitive.ObjectIDFromHex(id)
	result, err := c.repository.GetSupervisorById(ctx, objID)

	if err != nil {
		return err
	}

	if supervisor.Name != result.Name {
		result.Name = supervisor.Name
	}

	result.UpdatedAt = time.Now()

	return c.repository.UpdateSupervisor(ctx, result)
}
