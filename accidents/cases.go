package accidents

import (
	"context"

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
	panic("unimplemented")
}

// ActivateInjuredPart implements Cases.
func (c *cases) ActivateInjuredPart(ctx context.Context, id string) error {
	panic("unimplemented")
}

// CreateAccident implements Cases.
func (c *cases) CreateAccident(ctx context.Context, accident *Accident) (*Accident, error) {
	panic("unimplemented")
}

// CreateClassification implements Cases.
func (c *cases) CreateClassification(ctx context.Context, classification *Classification) (*Classification, error) {
	panic("unimplemented")
}

// CreateGravity implements Cases.
func (c *cases) CreateGravity(ctx context.Context, gravity *Gravity) (*Gravity, error) {
	panic("unimplemented")
}

// CreateInjuredPart implements Cases.
func (c *cases) CreateInjuredPart(ctx context.Context, injuredPart *InjuredPart) (*InjuredPart, error) {
	panic("unimplemented")
}

// DeactivateClassification implements Cases.
func (c *cases) DeactivateClassification(ctx context.Context, id string) error {
	panic("unimplemented")
}

// DeactivateGravity implements Cases.
func (c *cases) DeactivateGravity(ctx context.Context, id string) error {
	panic("unimplemented")
}

// DeactivateInjuredPart implements Cases.
func (c *cases) DeactivateInjuredPart(ctx context.Context, id string) error {
	panic("unimplemented")
}

// DeleteAccident implements Cases.
func (c *cases) DeleteAccident(ctx context.Context, id string) error {
	panic("unimplemented")
}

// DeleteClassification implements Cases.
func (c *cases) DeleteClassification(ctx context.Context, id string) error {
	panic("unimplemented")
}

// DeleteGravity implements Cases.
func (c *cases) DeleteGravity(ctx context.Context, id string) error {
	panic("unimplemented")
}

// DeleteInjuredPart implements Cases.
func (c *cases) DeleteInjuredPart(ctx context.Context, id string) error {
	panic("unimplemented")
}

// GetAccidentById implements Cases.
func (c *cases) GetAccidentById(ctx context.Context, id string) (*Accident, error) {
	panic("unimplemented")
}

// GetAccidents implements Cases.
func (c *cases) GetAccidents(ctx context.Context) ([]Accident, error) {
	panic("unimplemented")
}

// GetClassificationById implements Cases.
func (c *cases) GetClassificationById(ctx context.Context, id string) (*Classification, error) {
	panic("unimplemented")
}

// GetClassifications implements Cases.
func (c *cases) GetClassifications(ctx context.Context) ([]Classification, error) {
	panic("unimplemented")
}

// GetGravities implements Cases.
func (c *cases) GetGravities(ctx context.Context) ([]Gravity, error) {
	panic("unimplemented")
}

// GetGravityById implements Cases.
func (c *cases) GetGravityById(ctx context.Context, id string) (*Gravity, error) {
	panic("unimplemented")
}

// GetInjuredPartById implements Cases.
func (c *cases) GetInjuredPartById(ctx context.Context, id string) (*InjuredPart, error) {
	panic("unimplemented")
}

// GetInjuredParts implements Cases.
func (c *cases) GetInjuredParts(ctx context.Context) ([]InjuredPart, error) {
	panic("unimplemented")
}

// UpdateAccident implements Cases.
func (c *cases) UpdateAccident(ctx context.Context, id string, accident *Accident) error {
	panic("unimplemented")
}

// UpdateClassification implements Cases.
func (c *cases) UpdateClassification(ctx context.Context, id string, classification *Classification) error {
	panic("unimplemented")
}

// UpdateGravity implements Cases.
func (c *cases) UpdateGravity(ctx context.Context, id string, gravity *Gravity) error {
	panic("unimplemented")
}

// UpdateInjuredPart implements Cases.
func (c *cases) UpdateInjuredPart(ctx context.Context, id string, injuredPart *InjuredPart) error {
	panic("unimplemented")
}
