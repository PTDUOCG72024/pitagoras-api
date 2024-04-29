package employees

import "github.com/gofiber/fiber/v2"

type controller struct {
	cases Cases
}

func NewController(cases Cases) Controller {
	return &controller{cases}
}

// ActivateEmployee implements Controller.
func (c *controller) ActivateEmployee(ctx *fiber.Ctx) error {
	panic("unimplemented")
}

// ActivateNationality implements Controller.
func (c *controller) ActivateNationality(ctx *fiber.Ctx) error {
	panic("unimplemented")
}

// ActivatePosition implements Controller.
func (c *controller) ActivatePosition(ctx *fiber.Ctx) error {
	panic("unimplemented")
}

// ActivateSupervisor implements Controller.
func (c *controller) ActivateSupervisor(ctx *fiber.Ctx) error {
	panic("unimplemented")
}

// CreateEmployee implements Controller.
func (c *controller) CreateEmployee(ctx *fiber.Ctx) error {
	panic("unimplemented")
}

// CreateNationality implements Controller.
func (c *controller) CreateNationality(ctx *fiber.Ctx) error {
	panic("unimplemented")
}

// CreatePosition implements Controller.
func (c *controller) CreatePosition(ctx *fiber.Ctx) error {
	panic("unimplemented")
}

// CreateSupervisor implements Controller.
func (c *controller) CreateSupervisor(ctx *fiber.Ctx) error {
	panic("unimplemented")
}

// DeactivateEmployee implements Controller.
func (c *controller) DeactivateEmployee(ctx *fiber.Ctx) error {
	panic("unimplemented")
}

// DeactivateNationality implements Controller.
func (c *controller) DeactivateNationality(ctx *fiber.Ctx) error {
	panic("unimplemented")
}

// DeactivatePosition implements Controller.
func (c *controller) DeactivatePosition(ctx *fiber.Ctx) error {
	panic("unimplemented")
}

// DeactivateSupervisor implements Controller.
func (c *controller) DeactivateSupervisor(ctx *fiber.Ctx) error {
	panic("unimplemented")
}

// DeleteEmployee implements Controller.
func (c *controller) DeleteEmployee(ctx *fiber.Ctx) error {
	panic("unimplemented")
}

// DeleteNationality implements Controller.
func (c *controller) DeleteNationality(ctx *fiber.Ctx) error {
	panic("unimplemented")
}

// DeletePosition implements Controller.
func (c *controller) DeletePosition(ctx *fiber.Ctx) error {
	panic("unimplemented")
}

// DeleteSupervisor implements Controller.
func (c *controller) DeleteSupervisor(ctx *fiber.Ctx) error {
	panic("unimplemented")
}

// GetEmployeeById implements Controller.
func (c *controller) GetEmployeeById(ctx *fiber.Ctx) error {
	panic("unimplemented")
}

// GetEmployees implements Controller.
func (c *controller) GetEmployees(ctx *fiber.Ctx) error {
	panic("unimplemented")
}

// GetNationalities implements Controller.
func (c *controller) GetNationalities(ctx *fiber.Ctx) error {
	panic("unimplemented")
}

// GetNationalityById implements Controller.
func (c *controller) GetNationalityById(ctx *fiber.Ctx) error {
	panic("unimplemented")
}

// GetPositionById implements Controller.
func (c *controller) GetPositionById(ctx *fiber.Ctx) error {
	panic("unimplemented")
}

// GetPositions implements Controller.
func (c *controller) GetPositions(ctx *fiber.Ctx) error {
	panic("unimplemented")
}

// GetSupervisorById implements Controller.
func (c *controller) GetSupervisorById(ctx *fiber.Ctx) error {
	panic("unimplemented")
}

// GetSupervisors implements Controller.
func (c *controller) GetSupervisors(ctx *fiber.Ctx) error {
	panic("unimplemented")
}

// UpdateEmployee implements Controller.
func (c *controller) UpdateEmployee(ctx *fiber.Ctx) error {
	panic("unimplemented")
}

// UpdateNationality implements Controller.
func (c *controller) UpdateNationality(ctx *fiber.Ctx) error {
	panic("unimplemented")
}

// UpdatePosition implements Controller.
func (c *controller) UpdatePosition(ctx *fiber.Ctx) error {
	panic("unimplemented")
}

// UpdateSupervisor implements Controller.
func (c *controller) UpdateSupervisor(ctx *fiber.Ctx) error {
	panic("unimplemented")
}
