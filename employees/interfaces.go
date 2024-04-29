package employees

import (
	"context"

	"github.com/gofiber/fiber/v2"
)

type Controller interface {
	CreateEmployee(ctx *fiber.Ctx) error
	GetEmployeeById(ctx *fiber.Ctx) error
	GetEmployees(ctx *fiber.Ctx) error
	UpdateEmployee(ctx *fiber.Ctx) error
	DeleteEmployee(ctx *fiber.Ctx) error
	ActivateEmployee(ctx *fiber.Ctx) error
	DeactivateEmployee(ctx *fiber.Ctx) error

	CreateSupervisor(ctx *fiber.Ctx) error
	GetSupervisorById(ctx *fiber.Ctx) error
	GetSupervisors(ctx *fiber.Ctx) error
	UpdateSupervisor(ctx *fiber.Ctx) error
	DeleteSupervisor(ctx *fiber.Ctx) error
	ActivateSupervisor(ctx *fiber.Ctx) error
	DeactivateSupervisor(ctx *fiber.Ctx) error

	CreatePosition(ctx *fiber.Ctx) error
	GetPositionById(ctx *fiber.Ctx) error
	GetPositions(ctx *fiber.Ctx) error
	UpdatePosition(ctx *fiber.Ctx) error
	DeletePosition(ctx *fiber.Ctx) error
	ActivatePosition(ctx *fiber.Ctx) error
	DeactivatePosition(ctx *fiber.Ctx) error

	CreateNationality(ctx *fiber.Ctx) error
	GetNationalityById(ctx *fiber.Ctx) error
	GetNationalities(ctx *fiber.Ctx) error
	UpdateNationality(ctx *fiber.Ctx) error
	DeleteNationality(ctx *fiber.Ctx) error
	ActivateNationality(ctx *fiber.Ctx) error
	DeactivateNationality(ctx *fiber.Ctx) error
}

type Cases interface {
	CreateEmployee(ctx context.Context, employee *Employee) (*Employee, error)
	GetEmployeeById(ctx context.Context, id string) (*Employee, error)
	GetEmployees(ctx context.Context) ([]Employee, error)
	UpdateEmployee(ctx context.Context, id string, employee *Employee) error
	DeleteEmployee(ctx context.Context, id string) error
	ActivateEmployee(ctx context.Context, id string) error
	DeactivateEmployee(ctx context.Context, id string) error

	CreateSupervisor(ctx context.Context, supervisor *Supervisor) (*Supervisor, error)
	GetSupervisorById(ctx context.Context, id string) (*Supervisor, error)
	GetSupervisors(ctx context.Context) ([]Supervisor, error)
	UpdateSupervisor(ctx context.Context, id string, supervisor *Supervisor) error
	DeleteSupervisor(ctx context.Context, id string) error
	ActivateSupervisor(ctx context.Context, id string) error
	DeactivateSupervisor(ctx context.Context, id string) error

	CreatePosition(ctx context.Context, position *Position) (*Position, error)
	GetPositionById(ctx context.Context, id string) (*Position, error)
	GetPositions(ctx context.Context) ([]Position, error)
	UpdatePosition(ctx context.Context, id string, position *Position) error
	DeletePosition(ctx context.Context, id string) error
	ActivatePosition(ctx context.Context, id string) error
	DeactivatePosition(ctx context.Context, id string) error

	CreateNationality(ctx context.Context, nationality *Nationality) (*Nationality, error)
	GetNationalityById(ctx context.Context, id string) (*Nationality, error)
	GetNationalities(ctx context.Context) ([]Nationality, error)
	UpdateNationality(ctx context.Context, id string, nationality *Nationality) error
	DeleteNationality(ctx context.Context, id string) error
	ActivateNationality(ctx context.Context, id string) error
	DeactivateNationality(ctx context.Context, id string) error
}

type Repository interface {
	CreateEmployee(ctx context.Context, employee *Employee) (*Employee, error)
	UpdateEmployee(ctx context.Context, employee *Employee) error
	GetEmployeeById(ctx context.Context, id string) (*Employee, error)
	GetEmployeeByEmail(ctx context.Context, email string) (*Employee, error)
	GetEmployeeByIdentificationNumber(ctx context.Context, identificationNumber string) (*Employee, error)
	GetEmployees(ctx context.Context) ([]Employee, error)

	CreateSupervisor(ctx context.Context, supervisor *Supervisor) (*Supervisor, error)
	UpdateSupervisor(ctx context.Context, supervisor *Supervisor) error
	GetSupervisorById(ctx context.Context, id string) (*Supervisor, error)
	GetSupervisorByName(ctx context.Context, name string) (*Supervisor, error)
	GetSupervisors(ctx context.Context) ([]Supervisor, error)

	CreatePosition(ctx context.Context, position *Position) (*Position, error)
	UpdatePosition(ctx context.Context, position *Position) error
	GetPositionById(ctx context.Context, id string) (*Position, error)
	GetPositionByName(ctx context.Context, name string) (*Position, error)
	GetPositions(ctx context.Context) ([]Position, error)

	CreateNationality(ctx context.Context, nationality *Nationality) (*Nationality, error)
	UpdateNationality(ctx context.Context, nationality *Nationality) error
	GetNationalityById(ctx context.Context, id string) (*Nationality, error)
	GetNationalityByName(ctx context.Context, name string) (*Nationality, error)
	GetNationalities(ctx context.Context) ([]Nationality, error)
}
