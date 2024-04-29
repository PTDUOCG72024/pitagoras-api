package employees

import (
	"github.com/gofiber/fiber/v2"
	"github.com/xbizzybone/pitagoras-api/utils"
)

type controller struct {
	cases Cases
}

func NewController(cases Cases) Controller {
	return &controller{cases}
}

// ActivateEmployee godoc
// @Summary Activate a new employee
// @Description Activate a new employee
// @Tags employees
// @Accept json
// @Produce json
// @Param id path string true "Employee ID"
// @Success 200 {string} string "Empleado activado correctamente"
// @Failure 500 {string} string "Error activando el empleado"
// @Router /employees/activate/{id} [put]
func (c *controller) ActivateEmployee(ctx *fiber.Ctx) error {
	employeeID := ctx.Params("id")

	if err := c.cases.ActivateEmployee(ctx.UserContext(), employeeID); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error activando el empleado",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Empleado activado correctamente",
	})
}

// ActivateNationality godoc
// @Summary Activate a new nationality
// @Description Activate a new nationality
// @Tags employees
// @Accept json
// @Produce json
// @Param id path string true "Nationality ID"
// @Success 200 {string} string "Nacionalidad activada correctamente"
// @Failure 500 {string} string "Error activando la nacionalidad"
// @Router /employees/nationalities/activate/{id} [put]
func (c *controller) ActivateNationality(ctx *fiber.Ctx) error {
	nationalityID := ctx.Params("id")

	if err := c.cases.ActivateNationality(ctx.UserContext(), nationalityID); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error activando la nacionalidad",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Nacionalidad activada correctamente",
	})
}

// ActivatePosition godoc
// @Summary Activate a new position
// @Description Activate a new position
// @Tags employees
// @Accept json
// @Produce json
// @Param id path string true "Position ID"
// @Success 200 {string} string "Cargo activado correctamente"
// @Failure 500 {string} string "Error activando el cargo"
// @Router /employees/positions/activate/{id} [put]
func (c *controller) ActivatePosition(ctx *fiber.Ctx) error {
	positionID := ctx.Params("id")

	if err := c.cases.ActivatePosition(ctx.UserContext(), positionID); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error activando el cargo",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Cargo activado correctamente",
	})
}

// ActivateSupervisor godoc
// @Summary Activate a new supervisor
// @Description Activate a new supervisor
// @Tags employees
// @Accept json
// @Produce json
// @Param id path string true "Supervisor ID"
// @Success 200 {string} string "Supervisor activado correctamente"
// @Failure 500 {string} string "Error activando el supervisor"
// @Router /employees/supervisors/activate/{id} [put]
func (c *controller) ActivateSupervisor(ctx *fiber.Ctx) error {
	supervisorID := ctx.Params("id")

	if err := c.cases.ActivateSupervisor(ctx.UserContext(), supervisorID); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error activando el supervisor",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Supervisor activado correctamente",
	})
}

// CreateEmployee godoc
// @Summary Create a new employee
// @Description Create a new employee
// @Tags employees
// @Accept json
// @Produce json
// @Param employee body CreateEmployeeRequest true "Employee object that needs to be created"
// @Success 200 {object} CreateEmployeeResponse
// @Failure 400 {string} string "Error validando el cuerpo de la petición"
// @Failure 500 {string} string "Error creando el empleado"
// @Router /employees [post]
func (c *controller) CreateEmployee(ctx *fiber.Ctx) error {
	requestBody := new(CreateEmployeeRequest)

	if err := utils.BodyParser(ctx, requestBody); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error validando el cuerpo de la petición :" + err.Error(),
		})
	}

	employee := new(Employee)
	employee.Name = requestBody.Name
	employee.IdentificationType = requestBody.IdentificationType
	employee.IdentificationNumber = requestBody.IdentificationNumber
	employee.Email = requestBody.Email
	employee.InLaw = requestBody.InLaw
	employee.ContractType = requestBody.ContractType
	employee.StartDate = requestBody.StartDate
	employee.EndDate = requestBody.EndDate
	employee.Supervisor.ID = requestBody.SupervisorID
	employee.Position.ID = requestBody.PositionID
	employee.Nationality.ID = requestBody.NationalityID

	result, err := c.cases.CreateEmployee(ctx.UserContext(), employee)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error creando el empleado",
		})
	}

	response := new(CreateEmployeeResponse)
	response.ID = result.ID
	response.Name = result.Name
	response.IdentificationType = result.IdentificationType
	response.IdentificationNumber = result.IdentificationNumber
	response.Email = result.Email
	response.InLaw = result.InLaw
	response.ContractType = result.ContractType
	response.StartDate = result.StartDate
	response.EndDate = result.EndDate
	response.IsActive = result.IsActive
	response.IsDeleted = result.IsDeleted
	response.Supervisor = result.Supervisor
	response.Position = result.Position
	response.Nationality = result.Nationality
	response.CreatedAt = result.CreatedAt

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message":  "Empleado creado correctamente",
		"employee": response,
	})
}

// CreateNationality godoc
// @Summary Create a new nationality
// @Description Create a new nationality
// @Tags employees
// @Accept json
// @Produce json
// @Param nationality body CreateNationalityRequest true "Nationality object that needs to be created"
// @Success 200 {object} CreateNationalityResponse
// @Failure 400 {string} string "Error validando el cuerpo de la petición"
// @Failure 500 {string} string "Error creando la nacionalidad"
// @Router /employees/nationalities [post]
func (c *controller) CreateNationality(ctx *fiber.Ctx) error {
	requestBody := new(CreateNationalityRequest)

	if err := utils.BodyParser(ctx, requestBody); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error validando el cuerpo de la petición :" + err.Error(),
		})
	}

	nationality := new(Nationality)
	nationality.Name = requestBody.Name
	nationality.IsActive = true
	nationality.IsDeleted = false

	result, err := c.cases.CreateNationality(ctx.UserContext(), nationality)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error creando la nacionalidad",
		})
	}

	response := new(CreateNationalityResponse)
	response.ID = result.ID
	response.Name = result.Name
	response.IsActive = result.IsActive
	response.IsDeleted = result.IsDeleted
	response.CreatedAt = result.CreatedAt

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message":     "Nacionalidad creada correctamente",
		"nationality": response,
	})
}

// CreatePosition godoc
// @Summary Create a new position
// @Description Create a new position
// @Tags employees
// @Accept json
// @Produce json
// @Param position body CreatePositionRequest true "Position object that needs to be created"
// @Success 200 {object} CreatePositionResponse
// @Failure 400 {string} string "Error validando el cuerpo de la petición"
// @Failure 500 {string} string "Error creando el cargo"
// @Router /employees/positions [post]
func (c *controller) CreatePosition(ctx *fiber.Ctx) error {
	requestBody := new(CreatePositionRequest)

	if err := utils.BodyParser(ctx, requestBody); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error validando el cuerpo de la petición :" + err.Error(),
		})
	}

	position := new(Position)
	position.Name = requestBody.Name
	position.IsActive = true
	position.IsDeleted = false

	result, err := c.cases.CreatePosition(ctx.UserContext(), position)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error creando el cargo",
		})
	}

	response := new(CreatePositionResponse)
	response.ID = result.ID
	response.Name = result.Name
	response.IsActive = result.IsActive
	response.IsDeleted = result.IsDeleted
	response.CreatedAt = result.CreatedAt

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message":  "Cargo creado correctamente",
		"position": response,
	})
}

// CreateSupervisor godoc
// @Summary Create a new supervisor
// @Description Create a new supervisor
// @Tags employees
// @Accept json
// @Produce json
// @Param supervisor body CreateSupervisorRequest true "Supervisor object that needs to be created"
// @Success 200 {object} CreateSupervisorResponse
// @Failure 400 {string} string "Error validando el cuerpo de la petición"
// @Failure 500 {string} string "Error creando el supervisor"
// @Router /employees/supervisors [post]
func (c *controller) CreateSupervisor(ctx *fiber.Ctx) error {
	requestBody := new(CreateSupervisorRequest)

	if err := utils.BodyParser(ctx, requestBody); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error validando el cuerpo de la petición :" + err.Error(),
		})
	}

	supervisor := new(Supervisor)
	supervisor.Name = requestBody.Name
	supervisor.IsActive = true
	supervisor.IsDeleted = false

	result, err := c.cases.CreateSupervisor(ctx.UserContext(), supervisor)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error creando el supervisor",
		})
	}

	response := new(CreateSupervisorResponse)
	response.ID = result.ID
	response.Name = result.Name
	response.IsActive = result.IsActive
	response.IsDeleted = result.IsDeleted
	response.CreatedAt = result.CreatedAt

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message":    "Supervisor creado correctamente",
		"supervisor": response,
	})
}

// DeactivateEmployee godoc
// @Summary Deactivate a employee
// @Description Deactivate a employee
// @Tags employees
// @Accept json
// @Produce json
// @Param id path string true "Employee ID"
// @Success 200 {string} string "Empleado desactivado correctamente"
// @Failure 500 {string} string "Error desactivando el empleado"
// @Router /employees/deactivate/{id} [put]
func (c *controller) DeactivateEmployee(ctx *fiber.Ctx) error {
	employeeID := ctx.Params("id")

	if err := c.cases.DeactivateEmployee(ctx.UserContext(), employeeID); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error desactivando el empleado",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Empleado desactivado correctamente",
	})
}

// DeactivateNationality godoc
// @Summary Deactivate a nationality
// @Description Deactivate a nationality
// @Tags employees
// @Accept json
// @Produce json
// @Param id path string true "Nationality ID"
// @Success 200 {string} string "Nacionalidad desactivada correctamente"
// @Failure 500 {string} string "Error desactivando la nacionalidad"
// @Router /employees/nationalities/deactivate/{id} [put]
func (c *controller) DeactivateNationality(ctx *fiber.Ctx) error {
	nationalityID := ctx.Params("id")

	if err := c.cases.DeactivateNationality(ctx.UserContext(), nationalityID); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error desactivando la nacionalidad",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Nacionalidad desactivada correctamente",
	})
}

// DeactivatePosition godoc
// @Summary Deactivate a position
// @Description Deactivate a position
// @Tags employees
// @Accept json
// @Produce json
// @Param id path string true "Position ID"
// @Success 200 {string} string "Cargo desactivado correctamente"
// @Failure 500 {string} string "Error desactivando el cargo"
// @Router /employees/positions/deactivate/{id} [put]
func (c *controller) DeactivatePosition(ctx *fiber.Ctx) error {
	positionID := ctx.Params("id")

	if err := c.cases.DeactivatePosition(ctx.UserContext(), positionID); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error desactivando el cargo",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Cargo desactivado correctamente",
	})
}

// DeactivateSupervisor godoc
// @Summary Deactivate a supervisor
// @Description Deactivate a supervisor
// @Tags employees
// @Accept json
// @Produce json
// @Param id path string true "Supervisor ID"
// @Success 200 {string} string "Supervisor desactivado correctamente"
// @Failure 500 {string} string "Error desactivando el supervisor"
// @Router /employees/supervisors/deactivate/{id} [put]
func (c *controller) DeactivateSupervisor(ctx *fiber.Ctx) error {
	supervisorID := ctx.Params("id")

	if err := c.cases.DeactivateSupervisor(ctx.UserContext(), supervisorID); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error desactivando el supervisor",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Supervisor desactivado correctamente",
	})
}

// DeleteEmployee godoc
// @Summary Delete a employee
// @Description Delete a employee
// @Tags employees
// @Accept json
// @Produce json
// @Param id path string true "Employee ID"
// @Success 200 {string} string "Empleado eliminado correctamente"
// @Failure 500 {string} string "Error eliminando el empleado"
// @Router /employees/{id} [delete]
func (c *controller) DeleteEmployee(ctx *fiber.Ctx) error {
	employeeID := ctx.Params("id")

	if err := c.cases.DeleteEmployee(ctx.UserContext(), employeeID); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error eliminando el empleado",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Empleado eliminado correctamente",
	})
}

// DeleteNationality godoc
// @Summary Delete a nationality
// @Description Delete a nationality
// @Tags employees
// @Accept json
// @Produce json
// @Param id path string true "Nationality ID"
// @Success 200 {string} string "Nacionalidad eliminada correctamente"
// @Failure 500 {string} string "Error eliminando la nacionalidad"
// @Router /employees/nationalities/{id} [delete]
func (c *controller) DeleteNationality(ctx *fiber.Ctx) error {
	nationalityID := ctx.Params("id")

	if err := c.cases.DeleteNationality(ctx.UserContext(), nationalityID); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error eliminando la nacionalidad",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Nacionalidad eliminada correctamente",
	})
}

// DeletePosition godoc
// @Summary Delete a position
// @Description Delete a position
// @Tags employees
// @Accept json
// @Produce json
// @Param id path string true "Position ID"
// @Success 200 {string} string "Cargo eliminado correctamente"
// @Failure 500 {string} string "Error eliminando el cargo"
// @Router /employees/positions/{id} [delete]
func (c *controller) DeletePosition(ctx *fiber.Ctx) error {
	positionID := ctx.Params("id")

	if err := c.cases.DeletePosition(ctx.UserContext(), positionID); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error eliminando el cargo",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Cargo eliminado correctamente",
	})
}

// DeleteSupervisor godoc
// @Summary Delete a supervisor
// @Description Delete a supervisor
// @Tags employees
// @Accept json
// @Produce json
// @Param id path string true "Supervisor ID"
// @Success 200 {string} string "Supervisor eliminado correctamente"
// @Failure 500 {string} string "Error eliminando el supervisor"
// @Router /employees/supervisors/{id} [delete]
func (c *controller) DeleteSupervisor(ctx *fiber.Ctx) error {
	supervisorID := ctx.Params("id")

	if err := c.cases.DeleteSupervisor(ctx.UserContext(), supervisorID); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error eliminando el supervisor",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Supervisor eliminado correctamente",
	})
}

// GetEmployeeById godoc
// @Summary Get a employee by id
// @Description Get a employee by id
// @Tags employees
// @Accept json
// @Produce json
// @Param id path string true "Employee ID"
// @Success 200 {object} GetEmployeeResponse
// @Failure 500 {string} string "Error obteniendo el empleado"
// @Router /employees/{id} [get]
func (c *controller) GetEmployeeById(ctx *fiber.Ctx) error {
	employeeID := ctx.Params("id")

	employee, err := c.cases.GetEmployeeById(ctx.UserContext(), employeeID)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error obteniendo el empleado",
		})
	}

	employeeResponse := new(GetEmployeeResponse)
	employeeResponse.ID = employee.ID
	employeeResponse.Name = employee.Name
	employeeResponse.IdentificationType = employee.IdentificationType
	employeeResponse.IdentificationNumber = employee.IdentificationNumber
	employeeResponse.Email = employee.Email
	employeeResponse.InLaw = employee.InLaw
	employeeResponse.ContractType = employee.ContractType
	employeeResponse.StartDate = employee.StartDate
	employeeResponse.EndDate = employee.EndDate
	employeeResponse.IsActive = employee.IsActive
	employeeResponse.IsDeleted = employee.IsDeleted
	employeeResponse.Supervisor = employee.Supervisor
	employeeResponse.Position = employee.Position
	employeeResponse.Nationality = employee.Nationality

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":  "Empleado obtenido correctamente",
		"employee": employeeResponse,
	})
}

// GetEmployees godoc
// @Summary Get all employees
// @Description Get all employees
// @Tags employees
// @Accept json
// @Produce json
// @Success 200 {object} []GetEmployeeResponse
// @Failure 500 {string} string "Error obteniendo los empleados"
// @Router /employees [get]
func (c *controller) GetEmployees(ctx *fiber.Ctx) error {
	employees, err := c.cases.GetEmployees(ctx.UserContext())

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error obteniendo los empleados",
		})
	}

	employeesResponse := make([]GetEmployeeResponse, 0)
	for _, employee := range employees {
		employeeResponse := new(GetEmployeeResponse)
		employeeResponse.ID = employee.ID
		employeeResponse.Name = employee.Name
		employeeResponse.IdentificationType = employee.IdentificationType
		employeeResponse.IdentificationNumber = employee.IdentificationNumber
		employeeResponse.Email = employee.Email
		employeeResponse.InLaw = employee.InLaw
		employeeResponse.ContractType = employee.ContractType
		employeeResponse.StartDate = employee.StartDate
		employeeResponse.EndDate = employee.EndDate
		employeeResponse.IsActive = employee.IsActive
		employeeResponse.IsDeleted = employee.IsDeleted
		employeeResponse.Supervisor = employee.Supervisor
		employeeResponse.Position = employee.Position
		employeeResponse.Nationality = employee.Nationality

		employeesResponse = append(employeesResponse, *employeeResponse)
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":   "Empleados obtenidos correctamente",
		"employees": employeesResponse,
	})
}

// GetNationalities godoc
// @Summary Get all nationalities
// @Description Get all nationalities
// @Tags employees
// @Accept json
// @Produce json
// @Success 200 {object} []GetNationalityResponse
// @Failure 500 {string} string "Error obteniendo las nacionalidades"
// @Router /employees/nationalities [get]
func (c *controller) GetNationalities(ctx *fiber.Ctx) error {
	nationalities, err := c.cases.GetNationalities(ctx.UserContext())

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error obteniendo las nacionalidades",
		})
	}

	nationalitiesResponse := make([]GetNationalityResponse, 0)
	for _, nationality := range nationalities {
		nationalityResponse := new(GetNationalityResponse)
		nationalityResponse.ID = nationality.ID
		nationalityResponse.Name = nationality.Name
		nationalityResponse.IsActive = nationality.IsActive
		nationalityResponse.IsDeleted = nationality.IsDeleted
		nationalityResponse.CreatedAt = nationality.CreatedAt
		nationalityResponse.UpdatedAt = nationality.UpdatedAt
		nationalityResponse.DeleteAt = nationality.DeleteAt

		nationalitiesResponse = append(nationalitiesResponse, *nationalityResponse)
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":       "Nacionalidades obtenidas correctamente",
		"nationalities": nationalitiesResponse,
	})
}

// GetNationalityById godoc
// @Summary Get nationality by id
// @Description Get nationality by id
// @Tags employees
// @Accept json
// @Produce json
// @Param id path string true "Nationality ID"
// @Success 200 {object} GetNationalityResponse
// @Failure 500 {string} string "Error obteniendo la nacionalidad"
// @Router /employees/nationalities/{id} [get]
func (c *controller) GetNationalityById(ctx *fiber.Ctx) error {
	nationalityID := ctx.Params("id")

	nationality, err := c.cases.GetNationalityById(ctx.UserContext(), nationalityID)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error obteniendo la nacionalidad",
		})
	}

	nationalityResponse := new(GetNationalityResponse)
	nationalityResponse.ID = nationality.ID
	nationalityResponse.Name = nationality.Name
	nationalityResponse.IsActive = nationality.IsActive
	nationalityResponse.IsDeleted = nationality.IsDeleted
	nationalityResponse.CreatedAt = nationality.CreatedAt
	nationalityResponse.UpdatedAt = nationality.UpdatedAt
	nationalityResponse.DeleteAt = nationality.DeleteAt

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":     "Nacionalidad obtenida correctamente",
		"nationality": nationalityResponse,
	})
}

// GetPositionById godoc
// @Summary Get position by id
// @Description Get position by id
// @Tags employees
// @Accept json
// @Produce json
// @Param id path string true "Position ID"
// @Success 200 {object} GetPositionResponse
// @Failure 500 {string} string "Error obteniendo el cargo"
// @Router /employees/positions/{id} [get]
func (c *controller) GetPositionById(ctx *fiber.Ctx) error {
	positionID := ctx.Params("id")

	position, err := c.cases.GetPositionById(ctx.UserContext(), positionID)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error obteniendo el cargo",
		})
	}

	positionResponse := new(GetPositionResponse)
	positionResponse.ID = position.ID
	positionResponse.Name = position.Name
	positionResponse.IsActive = position.IsActive
	positionResponse.IsDeleted = position.IsDeleted
	positionResponse.CreatedAt = position.CreatedAt
	positionResponse.UpdatedAt = position.UpdatedAt
	positionResponse.DeleteAt = position.DeleteAt

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":  "Cargo obtenido correctamente",
		"position": positionResponse,
	})
}

// GetPositions godoc
// @Summary Get all positions
// @Description Get all positions
// @Tags employees
// @Accept json
// @Produce json
// @Success 200 {object} []GetPositionResponse
// @Failure 500 {string} string "Error obteniendo los cargos"
// @Router /employees/positions [get]
func (c *controller) GetPositions(ctx *fiber.Ctx) error {
	positions, err := c.cases.GetPositions(ctx.UserContext())

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error obteniendo los cargos",
		})
	}

	positionsResponse := make([]GetPositionResponse, 0)
	for _, position := range positions {
		positionResponse := new(GetPositionResponse)
		positionResponse.ID = position.ID
		positionResponse.Name = position.Name
		positionResponse.IsActive = position.IsActive
		positionResponse.IsDeleted = position.IsDeleted
		positionResponse.CreatedAt = position.CreatedAt
		positionResponse.UpdatedAt = position.UpdatedAt
		positionResponse.DeleteAt = position.DeleteAt

		positionsResponse = append(positionsResponse, *positionResponse)
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":   "Cargos obtenidos correctamente",
		"positions": positionsResponse,
	})
}

// GetSupervisorById godoc
// @Summary Get supervisor by id
// @Description Get supervisor by id
// @Tags employees
// @Accept json
// @Produce json
// @Param id path string true "Supervisor ID"
// @Success 200 {object} GetSupervisorResponse
// @Failure 500 {string} string "Error obteniendo el supervisor"
// @Router /employees/supervisors/{id} [get]
func (c *controller) GetSupervisorById(ctx *fiber.Ctx) error {
	supervisorID := ctx.Params("id")

	supervisor, err := c.cases.GetSupervisorById(ctx.UserContext(), supervisorID)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error obteniendo el supervisor",
		})
	}

	supervisorResponse := new(GetSupervisorResponse)
	supervisorResponse.ID = supervisor.ID
	supervisorResponse.Name = supervisor.Name
	supervisorResponse.IsActive = supervisor.IsActive
	supervisorResponse.IsDeleted = supervisor.IsDeleted
	supervisorResponse.CreatedAt = supervisor.CreatedAt
	supervisorResponse.UpdatedAt = supervisor.UpdatedAt
	supervisorResponse.DeleteAt = supervisor.DeleteAt

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":    "Supervisor obtenido correctamente",
		"supervisor": supervisorResponse,
	})
}

// GetSupervisors godoc
// @Summary Get all supervisors
// @Description Get all supervisors
// @Tags employees
// @Accept json
// @Produce json
// @Success 200 {object} []GetSupervisorResponse
// @Failure 500 {string} string "Error obteniendo los supervisores"
// @Router /employees/supervisors [get]
func (c *controller) GetSupervisors(ctx *fiber.Ctx) error {
	supervisors, err := c.cases.GetSupervisors(ctx.UserContext())

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error obteniendo los supervisores",
		})
	}

	supervisorsResponse := make([]GetSupervisorResponse, 0)
	for _, supervisor := range supervisors {
		supervisorResponse := new(GetSupervisorResponse)
		supervisorResponse.ID = supervisor.ID
		supervisorResponse.Name = supervisor.Name
		supervisorResponse.IsActive = supervisor.IsActive
		supervisorResponse.IsDeleted = supervisor.IsDeleted
		supervisorResponse.CreatedAt = supervisor.CreatedAt
		supervisorResponse.UpdatedAt = supervisor.UpdatedAt
		supervisorResponse.DeleteAt = supervisor.DeleteAt

		supervisorsResponse = append(supervisorsResponse, *supervisorResponse)
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":     "Supervisores obtenidos correctamente",
		"supervisors": supervisorsResponse,
	})
}

// UpdateEmployee godoc
// @Summary Update a employee
// @Description Update a employee
// @Tags employees
// @Accept json
// @Produce json
// @Param id path string true "Employee ID"
// @Param employee body UpdateEmployeeRequest true "Employee object that needs to be updated"
// @Success 200 {string} string "Empleado actualizado correctamente"
// @Failure 400 {string} string "Error validando el cuerpo de la petición"
// @Failure 500 {string} string "Error actualizando el empleado"
// @Router /employees/{id} [put]
func (c *controller) UpdateEmployee(ctx *fiber.Ctx) error {
	employeeID := ctx.Params("id")
	requestBody := new(UpdateEmployeeRequest)

	if err := utils.BodyParser(ctx, requestBody); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error validando el cuerpo de la petición :" + err.Error(),
		})
	}

	employee := new(Employee)
	employee.Name = requestBody.Name
	employee.IdentificationType = requestBody.IdentificationType
	employee.IdentificationNumber = requestBody.IdentificationNumber
	employee.Email = requestBody.Email
	employee.InLaw = requestBody.InLaw
	employee.ContractType = requestBody.ContractType
	employee.StartDate = requestBody.StartDate
	employee.EndDate = requestBody.EndDate
	employee.Supervisor.ID = requestBody.SupervisorID
	employee.Position.ID = requestBody.PositionID
	employee.Nationality.ID = requestBody.NationalityID

	err := c.cases.UpdateEmployee(ctx.UserContext(), employeeID, employee)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error actualizando el empleado",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Empleado actualizado correctamente",
	})
}

// UpdateNationality godoc
// @Summary Update a nationality
// @Description Update a nationality
// @Tags employees
// @Accept json
// @Produce json
// @Param id path string true "Nationality ID"
// @Param nationality body UpdateNationalityRequest true "Nationality object that needs to be updated"
// @Success 200 {string} string "Nacionalidad actualizada correctamente"
// @Failure 400 {string} string "Error validando el cuerpo de la petición"
// @Failure 500 {string} string "Error actualizando la nacionalidad"
// @Router /employees/nationalities/{id} [put]
func (c *controller) UpdateNationality(ctx *fiber.Ctx) error {
	nationalityID := ctx.Params("id")

	requestBody := new(UpdateNationalityRequest)

	if err := utils.BodyParser(ctx, requestBody); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error validando el cuerpo de la petición :" + err.Error(),
		})
	}

	nationality := new(Nationality)
	nationality.Name = requestBody.Name

	err := c.cases.UpdateNationality(ctx.UserContext(), nationalityID, nationality)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error actualizando la nacionalidad",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Nacionalidad actualizada correctamente",
	})
}

// UpdatePosition godoc
// @Summary Update a position
// @Description Update a position
// @Tags employees
// @Accept json
// @Produce json
// @Param id path string true "Position ID"
// @Param position body UpdatePositionRequest true "Position object that needs to be updated"
// @Success 200 {string} string "Cargo actualizado correctamente"
// @Failure 400 {string} string "Error validando el cuerpo de la petición"
// @Failure 500 {string} string "Error actualizando el cargo"
// @Router /employees/positions/{id} [put]
func (c *controller) UpdatePosition(ctx *fiber.Ctx) error {
	positionID := ctx.Params("id")

	requestBody := new(UpdatePositionRequest)

	if err := utils.BodyParser(ctx, requestBody); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error validando el cuerpo de la petición :" + err.Error(),
		})
	}

	position := new(Position)
	position.Name = requestBody.Name

	err := c.cases.UpdatePosition(ctx.UserContext(), positionID, position)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error actualizando el cargo",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Cargo actualizado correctamente",
	})
}

// UpdateSupervisor godoc
// @Summary Update a supervisor
// @Description Update a supervisor
// @Tags employees
// @Accept json
// @Produce json
// @Param id path string true "Supervisor ID"
// @Param supervisor body UpdateSupervisorRequest true "Supervisor object that needs to be updated"
// @Success 200 {string} string "Supervisor actualizado correctamente"
// @Failure 400 {string} string "Error validando el cuerpo de la petición"
// @Failure 500 {string} string "Error actualizando el supervisor"
// @Router /employees/supervisors/{id} [put]
func (c *controller) UpdateSupervisor(ctx *fiber.Ctx) error {
	supervisorID := ctx.Params("id")

	requestBody := new(UpdateSupervisorRequest)

	if err := utils.BodyParser(ctx, requestBody); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error validando el cuerpo de la petición :" + err.Error(),
		})
	}

	supervisor := new(Supervisor)
	supervisor.Name = requestBody.Name

	err := c.cases.UpdateSupervisor(ctx.UserContext(), supervisorID, supervisor)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error actualizando el supervisor",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Supervisor actualizado correctamente",
	})
}
