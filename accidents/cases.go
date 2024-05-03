package accidents

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

// ActivateClassification implements Cases.
func (c *cases) ActivateClassification(ctx context.Context, id string) error {
	objID, _ := primitive.ObjectIDFromHex(id)
	classification, err := c.repository.GetClassificationById(ctx, objID)
	if err != nil {
		return err
	}

	classification.IsActive = true
	classification.UpdatedAt = classification.CreatedAt

	return c.repository.UpdateClassification(ctx, classification)
}

// ActivateGravity implements Cases.
func (c *cases) ActivateGravity(ctx context.Context, id string) error {
	objID, _ := primitive.ObjectIDFromHex(id)
	gravity, err := c.repository.GetGravityById(ctx, objID)
	if err != nil {
		return err
	}

	gravity.IsActive = true
	gravity.UpdatedAt = gravity.CreatedAt

	return c.repository.UpdateGravity(ctx, gravity)
}

// ActivateInjuredPart implements Cases.
func (c *cases) ActivateInjuredPart(ctx context.Context, id string) error {
	objID, _ := primitive.ObjectIDFromHex(id)
	injuredPart, err := c.repository.GetInjuredPartById(ctx, objID)
	if err != nil {
		return err
	}

	injuredPart.IsActive = true
	injuredPart.UpdatedAt = injuredPart.CreatedAt

	return c.repository.UpdateInjuredPart(ctx, injuredPart)
}

// CreateAccident implements Cases.
func (c *cases) CreateAccident(ctx context.Context, accident *Accident) (*Accident, error) {
	classification, err := c.repository.GetClassificationById(ctx, accident.Classification.ID)

	if err != nil {
		return nil, errors.New("clasificación no encontrada")
	}

	accident.Classification = *classification

	gravity, err := c.repository.GetGravityById(ctx, accident.Gravity.ID)

	if err != nil {
		return nil, errors.New("gravedad no encontrada")
	}

	accident.Gravity = *gravity

	injuredPart, err := c.repository.GetInjuredPartById(ctx, accident.InjuredPart.ID)

	if err != nil {
		return nil, errors.New("parte lesionada no encontrada")
	}

	accident.InjuredPart = *injuredPart

	project, err := c.repository.GetProjectById(ctx, accident.Project.ID)

	if err != nil {
		return nil, errors.New("proyecto no encontrado")
	}

	accident.Project = *project

	employee, err := c.repository.GetEmployeeById(ctx, accident.Employee.ID)

	if err != nil {
		return nil, errors.New("empleado no encontrado")
	}

	accident.Employee = *employee

	accident.CreatedAt = time.Now()
	accident.UpdatedAt = time.Now()

	return c.repository.CreateAccident(ctx, accident)

}

// CreateClassification implements Cases.
func (c *cases) CreateClassification(ctx context.Context, classification *Classification) (*Classification, error) {
	cla, _ := c.repository.GetClassificationByName(ctx, classification.Name)

	if cla != nil {
		return nil, errors.New("clasificación ya existe")
	}

	classification.CreatedAt = time.Now()
	classification.UpdatedAt = time.Now()
	classification.IsActive = true

	return c.repository.CreateClassification(ctx, classification)
}

// CreateGravity implements Cases.
func (c *cases) CreateGravity(ctx context.Context, gravity *Gravity) (*Gravity, error) {
	grav, _ := c.repository.GetGravityByName(ctx, gravity.Name)

	if grav != nil {
		return nil, errors.New("gravedad ya existe")
	}

	gravity.CreatedAt = time.Now()
	gravity.UpdatedAt = time.Now()
	gravity.IsActive = true

	return c.repository.CreateGravity(ctx, gravity)
}

// CreateInjuredPart implements Cases.
func (c *cases) CreateInjuredPart(ctx context.Context, injuredPart *InjuredPart) (*InjuredPart, error) {
	inj, _ := c.repository.GetInjuredPartByName(ctx, injuredPart.Name)

	if inj != nil {
		return nil, errors.New("parte lesionada ya existe")
	}

	injuredPart.CreatedAt = time.Now()
	injuredPart.UpdatedAt = time.Now()
	injuredPart.IsActive = true

	return c.repository.CreateInjuredPart(ctx, injuredPart)
}

// DeactivateClassification implements Cases.
func (c *cases) DeactivateClassification(ctx context.Context, id string) error {
	objID, _ := primitive.ObjectIDFromHex(id)
	classification, err := c.repository.GetClassificationById(ctx, objID)
	if err != nil {
		return err
	}

	classification.IsActive = false
	classification.UpdatedAt = classification.CreatedAt

	return c.repository.UpdateClassification(ctx, classification)
}

// DeactivateGravity implements Cases.
func (c *cases) DeactivateGravity(ctx context.Context, id string) error {
	objID, _ := primitive.ObjectIDFromHex(id)
	gravity, err := c.repository.GetGravityById(ctx, objID)
	if err != nil {
		return err
	}

	gravity.IsActive = false
	gravity.UpdatedAt = gravity.CreatedAt

	return c.repository.UpdateGravity(ctx, gravity)
}

// DeactivateInjuredPart implements Cases.
func (c *cases) DeactivateInjuredPart(ctx context.Context, id string) error {
	objID, _ := primitive.ObjectIDFromHex(id)
	injuredPart, err := c.repository.GetInjuredPartById(ctx, objID)
	if err != nil {
		return err
	}

	injuredPart.IsActive = false
	injuredPart.UpdatedAt = injuredPart.CreatedAt

	return c.repository.UpdateInjuredPart(ctx, injuredPart)
}

// DeleteAccident implements Cases.
func (c *cases) DeleteAccident(ctx context.Context, id string) error {
	objID, _ := primitive.ObjectIDFromHex(id)
	accident, err := c.repository.GetAccidentById(ctx, objID)
	if err != nil {
		return err
	}

	accident.IsDeleted = true
	accident.DeletedAt = time.Now()

	return c.repository.UpdateAccident(ctx, accident)
}

// DeleteClassification implements Cases.
func (c *cases) DeleteClassification(ctx context.Context, id string) error {
	objID, _ := primitive.ObjectIDFromHex(id)
	classification, err := c.repository.GetClassificationById(ctx, objID)
	if err != nil {
		return err
	}

	classification.IsDeleted = true
	classification.DeletedAt = time.Now()

	return c.repository.UpdateClassification(ctx, classification)
}

// DeleteGravity implements Cases.
func (c *cases) DeleteGravity(ctx context.Context, id string) error {
	objID, _ := primitive.ObjectIDFromHex(id)
	gravity, err := c.repository.GetGravityById(ctx, objID)
	if err != nil {
		return err
	}

	gravity.IsDeleted = true
	gravity.DeletedAt = time.Now()

	return c.repository.UpdateGravity(ctx, gravity)
}

// DeleteInjuredPart implements Cases.
func (c *cases) DeleteInjuredPart(ctx context.Context, id string) error {
	objID, _ := primitive.ObjectIDFromHex(id)
	injuredPart, err := c.repository.GetInjuredPartById(ctx, objID)
	if err != nil {
		return err
	}

	injuredPart.IsDeleted = true
	injuredPart.DeletedAt = time.Now()

	return c.repository.UpdateInjuredPart(ctx, injuredPart)
}

// GetAccidentById implements Cases.
func (c *cases) GetAccidentById(ctx context.Context, id string) (*Accident, error) {
	objID, _ := primitive.ObjectIDFromHex(id)
	return c.repository.GetAccidentById(ctx, objID)
}

// GetAccidents implements Cases.
func (c *cases) GetAccidents(ctx context.Context, query GetAccidentsQuery) ([]Accident, error) {
	return c.repository.GetAccidents(ctx, query)
}

// GetClassificationById implements Cases.
func (c *cases) GetClassificationById(ctx context.Context, id string) (*Classification, error) {
	objID, _ := primitive.ObjectIDFromHex(id)
	return c.repository.GetClassificationById(ctx, objID)
}

// GetClassifications implements Cases.
func (c *cases) GetClassifications(ctx context.Context) ([]Classification, error) {
	return c.repository.GetClassifications(ctx)
}

// GetGravities implements Cases.
func (c *cases) GetGravities(ctx context.Context) ([]Gravity, error) {
	return c.repository.GetGravities(ctx)
}

// GetGravityById implements Cases.
func (c *cases) GetGravityById(ctx context.Context, id string) (*Gravity, error) {
	objID, _ := primitive.ObjectIDFromHex(id)
	return c.repository.GetGravityById(ctx, objID)
}

// GetInjuredPartById implements Cases.
func (c *cases) GetInjuredPartById(ctx context.Context, id string) (*InjuredPart, error) {
	objID, _ := primitive.ObjectIDFromHex(id)
	return c.repository.GetInjuredPartById(ctx, objID)
}

// GetInjuredParts implements Cases.
func (c *cases) GetInjuredParts(ctx context.Context) ([]InjuredPart, error) {
	return c.repository.GetInjuredParts(ctx)
}

// UpdateAccident implements Cases.
func (c *cases) UpdateAccident(ctx context.Context, id string, accident *Accident) error {
	objID, _ := primitive.ObjectIDFromHex(id)
	result, err := c.repository.GetAccidentById(ctx, objID)

	if err != nil {
		return err
	}

	classification, err := c.repository.GetClassificationById(ctx, accident.Classification.ID)

	if err != nil {
		return errors.New("clasificación no encontrada")
	}

	result.Classification = *classification

	gravity, err := c.repository.GetGravityById(ctx, accident.Gravity.ID)

	if err != nil {
		return errors.New("gravedad no encontrada")
	}

	result.Gravity = *gravity

	injuredPart, err := c.repository.GetInjuredPartById(ctx, accident.InjuredPart.ID)

	if err != nil {
		return errors.New("parte lesionada no encontrada")
	}

	result.InjuredPart = *injuredPart

	project, err := c.repository.GetProjectById(ctx, accident.Project.ID)

	if err != nil {
		return errors.New("proyecto no encontrado")
	}

	result.Project = *project

	employee, err := c.repository.GetEmployeeById(ctx, accident.Employee.ID)

	if err != nil {
		return errors.New("empleado no encontrado")
	}

	result.Employee = *employee

	if accident.Description != result.Description {
		result.Description = accident.Description
	}

	if accident.ConstructionArea != result.ConstructionArea {
		result.ConstructionArea = accident.ConstructionArea
	}

	if accident.AccidentDate != result.AccidentDate {
		result.AccidentDate = accident.AccidentDate
	}

	accident.UpdatedAt = time.Now()

	return c.repository.UpdateAccident(ctx, accident)
}

// UpdateClassification implements Cases.
func (c *cases) UpdateClassification(ctx context.Context, id string, classification *Classification) error {
	objID, _ := primitive.ObjectIDFromHex(id)
	result, err := c.repository.GetClassificationById(ctx, objID)

	if err != nil {
		return err
	}

	if classification.Name != result.Name {
		result.Name = classification.Name
	}

	result.UpdatedAt = time.Now()

	return c.repository.UpdateClassification(ctx, result)
}

// UpdateGravity implements Cases.
func (c *cases) UpdateGravity(ctx context.Context, id string, gravity *Gravity) error {
	objID, _ := primitive.ObjectIDFromHex(id)
	result, err := c.repository.GetGravityById(ctx, objID)

	if err != nil {
		return err
	}

	if gravity.Name != result.Name {
		result.Name = gravity.Name
	}

	result.UpdatedAt = time.Now()

	return c.repository.UpdateGravity(ctx, result)
}

// UpdateInjuredPart implements Cases.
func (c *cases) UpdateInjuredPart(ctx context.Context, id string, injuredPart *InjuredPart) error {
	objID, _ := primitive.ObjectIDFromHex(id)
	result, err := c.repository.GetInjuredPartById(ctx, objID)

	if err != nil {
		return err
	}

	if injuredPart.Name != result.Name {
		result.Name = injuredPart.Name
	}

	result.UpdatedAt = time.Now()

	return c.repository.UpdateInjuredPart(ctx, result)
}
