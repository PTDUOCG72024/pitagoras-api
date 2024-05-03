package accidents

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/xbizzybone/pitagoras-api/employees"
	"github.com/xbizzybone/pitagoras-api/projects"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Controller interface {
	CreateAccident(ctx *fiber.Ctx) error
	GetAccidentById(ctx *fiber.Ctx) error
	GetAccidents(ctx *fiber.Ctx) error
	UpdateAccident(ctx *fiber.Ctx) error
	DeleteAccident(ctx *fiber.Ctx) error

	CreateClassification(ctx *fiber.Ctx) error
	GetClassificationById(ctx *fiber.Ctx) error
	GetClassifications(ctx *fiber.Ctx) error
	UpdateClassification(ctx *fiber.Ctx) error
	DeleteClassification(ctx *fiber.Ctx) error
	ActivateClassification(ctx *fiber.Ctx) error
	DeactivateClassification(ctx *fiber.Ctx) error

	CreateGravity(ctx *fiber.Ctx) error
	GetGravityById(ctx *fiber.Ctx) error
	GetGravities(ctx *fiber.Ctx) error
	UpdateGravity(ctx *fiber.Ctx) error
	DeleteGravity(ctx *fiber.Ctx) error
	ActivateGravity(ctx *fiber.Ctx) error
	DeactivateGravity(ctx *fiber.Ctx) error

	CreateInjuredPart(ctx *fiber.Ctx) error
	GetInjuredPartById(ctx *fiber.Ctx) error
	GetInjuredParts(ctx *fiber.Ctx) error
	UpdateInjuredPart(ctx *fiber.Ctx) error
	DeleteInjuredPart(ctx *fiber.Ctx) error
	ActivateInjuredPart(ctx *fiber.Ctx) error
	DeactivateInjuredPart(ctx *fiber.Ctx) error
}

type Cases interface {
	CreateAccident(ctx context.Context, accident *Accident) (*Accident, error)
	GetAccidentById(ctx context.Context, id string) (*Accident, error)
	GetAccidents(ctx context.Context, query GetAccidentsQuery) ([]Accident, error)
	UpdateAccident(ctx context.Context, id string, accident *Accident) error
	DeleteAccident(ctx context.Context, id string) error

	CreateClassification(ctx context.Context, classification *Classification) (*Classification, error)
	GetClassificationById(ctx context.Context, id string) (*Classification, error)
	GetClassifications(ctx context.Context) ([]Classification, error)
	UpdateClassification(ctx context.Context, id string, classification *Classification) error
	DeleteClassification(ctx context.Context, id string) error
	ActivateClassification(ctx context.Context, id string) error
	DeactivateClassification(ctx context.Context, id string) error

	CreateGravity(ctx context.Context, gravity *Gravity) (*Gravity, error)
	GetGravityById(ctx context.Context, id string) (*Gravity, error)
	GetGravities(ctx context.Context) ([]Gravity, error)
	UpdateGravity(ctx context.Context, id string, gravity *Gravity) error
	DeleteGravity(ctx context.Context, id string) error
	ActivateGravity(ctx context.Context, id string) error
	DeactivateGravity(ctx context.Context, id string) error

	CreateInjuredPart(ctx context.Context, injuredPart *InjuredPart) (*InjuredPart, error)
	GetInjuredPartById(ctx context.Context, id string) (*InjuredPart, error)
	GetInjuredParts(ctx context.Context) ([]InjuredPart, error)
	UpdateInjuredPart(ctx context.Context, id string, injuredPart *InjuredPart) error
	DeleteInjuredPart(ctx context.Context, id string) error
	ActivateInjuredPart(ctx context.Context, id string) error
	DeactivateInjuredPart(ctx context.Context, id string) error
}

type Repository interface {
	CreateAccident(ctx context.Context, accident *Accident) (*Accident, error)
	GetAccidentById(ctx context.Context, id primitive.ObjectID) (*Accident, error)
	GetAccidents(ctx context.Context, query GetAccidentsQuery) ([]Accident, error)
	UpdateAccident(ctx context.Context, accident *Accident) error

	CreateClassification(ctx context.Context, classification *Classification) (*Classification, error)
	GetClassificationById(ctx context.Context, id primitive.ObjectID) (*Classification, error)
	GetClassificationByName(ctx context.Context, name string) (*Classification, error)
	GetClassifications(ctx context.Context) ([]Classification, error)
	UpdateClassification(ctx context.Context, classification *Classification) error

	CreateGravity(ctx context.Context, gravity *Gravity) (*Gravity, error)
	GetGravityById(ctx context.Context, id primitive.ObjectID) (*Gravity, error)
	GetGravityByName(ctx context.Context, name string) (*Gravity, error)
	GetGravities(ctx context.Context) ([]Gravity, error)
	UpdateGravity(ctx context.Context, gravity *Gravity) error

	CreateInjuredPart(ctx context.Context, injuredPart *InjuredPart) (*InjuredPart, error)
	GetInjuredPartById(ctx context.Context, id primitive.ObjectID) (*InjuredPart, error)
	GetInjuredPartByName(ctx context.Context, name string) (*InjuredPart, error)
	GetInjuredParts(ctx context.Context) ([]InjuredPart, error)
	UpdateInjuredPart(ctx context.Context, injuredPart *InjuredPart) error

	GetEmployeeById(ctx context.Context, id primitive.ObjectID) (*employees.Employee, error)
	GetProjectById(ctx context.Context, id primitive.ObjectID) (*projects.Project, error)
}
