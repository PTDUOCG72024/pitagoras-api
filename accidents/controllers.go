package accidents

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

// ActivateClassification godoc
// @Summary Activate a Classification
// @Description Activate a Classification
// @Tags classifications
// @Accept json
// @Produce json
// @Param id path string true "Classification ID"
// @Success 200 {string} string "Clasificación activada correctamente"
// @Failure 500 {string} string "Error activando la clasificación"
// @Router /classifications/activate/{id} [put]
func (c *controller) ActivateClassification(ctx *fiber.Ctx) error {
	classificationID := ctx.Params("id")

	if err := c.cases.ActivateClassification(ctx.Context(), classificationID); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error activando la clasificación",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Clasificación activada correctamente",
	})
}

// ActivateGravity godoc
// @Summary Activate a Gravity
// @Description Activate a Gravity
// @Tags gravities
// @Accept json
// @Produce json
// @Param id path string true "Gravity ID"
// @Success 200 {string} string "Gravedad activada correctamente"
// @Failure 500 {string} string "Error activando la gravedad"
// @Router /gravities/activate/{id} [put]
func (c *controller) ActivateGravity(ctx *fiber.Ctx) error {
	gravityID := ctx.Params("id")

	if err := c.cases.ActivateGravity(ctx.Context(), gravityID); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error activando la gravedad",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Gravedad activada correctamente",
	})
}

// ActivateInjuredPart godoc
// @Summary Activate a InjuredPart
// @Description Activate a InjuredPart
// @Tags injured-parts
// @Accept json
// @Produce json
// @Param id path string true "InjuredPart ID"
// @Success 200 {string} string "Parte lesionada activada correctamente"
// @Failure 500 {string} string "Error activando la parte lesionada"
// @Router /injured-parts/activate/{id} [put]
func (c *controller) ActivateInjuredPart(ctx *fiber.Ctx) error {
	injuredPartID := ctx.Params("id")

	if err := c.cases.ActivateInjuredPart(ctx.Context(), injuredPartID); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error activando la parte lesionada",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Parte lesionada activada correctamente",
	})
}

// CreateAccident godoc
// @Summary Create a new accident
// @Description Create a new accident
// @Tags accidents
// @Accept json
// @Produce json
// @Param body body CreateAccidentRequest true "Create Accident Request"
// @Success 201 {object} CreateAccidentResponse "Accidente creado correctamente"
// @Failure 400 {string} string "Error validando el cuerpo de la petición"
// @Failure 500 {string} string "Error creando el accidente"
// @Router /accidents/ [post]
func (c *controller) CreateAccident(ctx *fiber.Ctx) error {
	requestBody := new(CreateAccidentRequest)

	if err := utils.BodyParser(ctx, requestBody); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error validando el cuerpo de la petición :" + err.Error(),
		})
	}

	accident := new(Accident)
	accident.Description = requestBody.Description
	accident.ConstructionArea = requestBody.ConstructionArea
	accident.Classification.ID = requestBody.ClassificationID
	accident.Gravity.ID = requestBody.GravityID
	accident.Project.ID = requestBody.ProjectID
	accident.Employee.ID = requestBody.EmployeeID
	accident.InjuredPart.ID = requestBody.InjuredPartID
	accident.AccidentDate = requestBody.AccidentDate

	result, err := c.cases.CreateAccident(ctx.Context(), accident)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error creando el accidente",
		})
	}

	response := new(CreateAccidentResponse)
	response.ID = result.ID
	response.Description = result.Description
	response.ConstructionArea = result.ConstructionArea
	response.Classification = result.Classification
	response.Gravity = result.Gravity
	response.Project = result.Project
	response.Employee = result.Employee
	response.InjuredPart = result.InjuredPart
	response.AccidentDate = result.AccidentDate
	response.IsDeleted = result.IsDeleted
	response.CreatedAt = result.CreatedAt

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message":  "Accidente creado correctamente",
		"accident": response,
	})
}

// CreateClassification godoc
// @Summary Create a new Classification
// @Description Create a new Classification
// @Tags classifications
// @Accept json
// @Produce json
// @Param body body CreateClassificationRequest true "Create Classification Request"
// @Success 201 {object} CreateClassificationResponse "Clasificación creada correctamente"
// @Failure 400 {string} string "Error validando el cuerpo de la petición"
// @Failure 500 {string} string "Error creando la clasificación"
// @Router /classifications/ [post]
func (c *controller) CreateClassification(ctx *fiber.Ctx) error {
	requestBody := new(CreateClassificationRequest)

	if err := utils.BodyParser(ctx, requestBody); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error validando el cuerpo de la petición :" + err.Error(),
		})
	}

	classification := new(Classification)
	classification.Name = requestBody.Name

	result, err := c.cases.CreateClassification(ctx.Context(), classification)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error creando la clasificación",
		})
	}

	response := new(CreateClassificationResponse)
	response.ID = result.ID
	response.Name = result.Name

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message":        "Clasificación creada correctamente",
		"classification": response,
	})
}

// CreateGravity godoc
// @Summary Create a new Gravity
// @Description Create a new Gravity
// @Tags gravities
// @Accept json
// @Produce json
// @Param body body CreateGravityRequest true "Create Gravity Request"
// @Success 201 {object} CreateGravityResponse "Gravedad creada correctamente"
// @Failure 400 {string} string "Error validando el cuerpo de la petición"
// @Failure 500 {string} string "Error creando la gravedad"
// @Router /gravities/ [post]
func (c *controller) CreateGravity(ctx *fiber.Ctx) error {
	requestBody := new(CreateGravityRequest)

	if err := utils.BodyParser(ctx, requestBody); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error validando el cuerpo de la petición :" + err.Error(),
		})
	}

	gravity := new(Gravity)
	gravity.Name = requestBody.Name

	result, err := c.cases.CreateGravity(ctx.Context(), gravity)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error creando la gravedad",
		})
	}

	response := new(CreateGravityResponse)
	response.ID = result.ID
	response.Name = result.Name

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Gravedad creada correctamente",
		"gravity": response,
	})
}

// CreateInjuredPart godoc
// @Summary Create a new InjuredPart
// @Description Create a new InjuredPart
// @Tags injured-parts
// @Accept json
// @Produce json
// @Param body body CreateInjuredPartRequest true "Create InjuredPart Request"
// @Success 201 {object} CreateInjuredPartResponse "Parte lesionada creada correctamente"
// @Failure 400 {string} string "Error validando el cuerpo de la petición"
// @Failure 500 {string} string "Error creando la parte lesionada"
// @Router /injured-parts/ [post]
func (c *controller) CreateInjuredPart(ctx *fiber.Ctx) error {
	requestBody := new(CreateInjuredPartRequest)

	if err := utils.BodyParser(ctx, requestBody); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error validando el cuerpo de la petición :" + err.Error(),
		})
	}

	injuredPart := new(InjuredPart)
	injuredPart.Name = requestBody.Name

	result, err := c.cases.CreateInjuredPart(ctx.Context(), injuredPart)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error creando la parte lesionada",
		})
	}

	response := new(CreateInjuredPartResponse)
	response.ID = result.ID
	response.Name = result.Name

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message":      "Parte lesionada creada correctamente",
		"injured_part": response,
	})
}

// DeactivateClassification godoc
// @Summary Deactivate a Classification
// @Description Deactivate a Classification
// @Tags classifications
// @Accept json
// @Produce json
// @Param id path string true "Classification ID"
// @Success 200 {string} string "Clasificación desactivada correctamente"
// @Failure 500 {string} string "Error desactivando la clasificación"
// @Router /classifications/deactivate/{id} [put]
func (c *controller) DeactivateClassification(ctx *fiber.Ctx) error {
	classificationID := ctx.Params("id")

	if err := c.cases.DeactivateClassification(ctx.Context(), classificationID); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error desactivando la clasificación",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Clasificación desactivada correctamente",
	})
}

// DeactivateGravity godoc
// @Summary Deactivate a Gravity
// @Description Deactivate a Gravity
// @Tags gravities
// @Accept json
// @Produce json
// @Param id path string true "Gravity ID"
// @Success 200 {string} string "Gravedad desactivada correctamente"
// @Failure 500 {string} string "Error desactivando la gravedad"
// @Router /gravities/deactivate/{id} [put]
func (c *controller) DeactivateGravity(ctx *fiber.Ctx) error {
	gravityID := ctx.Params("id")

	if err := c.cases.DeactivateGravity(ctx.Context(), gravityID); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error desactivando la gravedad",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Gravedad desactivada correctamente",
	})
}

// DeactivateInjuredPart godoc
// @Summary Deactivate a InjuredPart
// @Description Deactivate a InjuredPart
// @Tags injured-parts
// @Accept json
// @Produce json
// @Param id path string true "InjuredPart ID"
// @Success 200 {string} string "Parte lesionada desactivada correctamente"
// @Failure 500 {string} string "Error desactivando la parte lesionada"
// @Router /injured-parts/deactivate/{id} [put]
func (c *controller) DeactivateInjuredPart(ctx *fiber.Ctx) error {
	injuredPartID := ctx.Params("id")

	if err := c.cases.DeactivateInjuredPart(ctx.Context(), injuredPartID); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error desactivando la parte lesionada",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Parte lesionada desactivada correctamente",
	})
}

// DeleteAccident godoc
// @Summary Delete an accident
// @Description Delete an accident
// @Tags accidents
// @Accept json
// @Produce json
// @Param id path string true "Accident ID"
// @Success 200 {string} string "Accidente eliminado correctamente"
// @Failure 500 {string} string "Error eliminando el accidente"
// @Router /accidents/{id} [delete]
func (c *controller) DeleteAccident(ctx *fiber.Ctx) error {
	accidentID := ctx.Params("id")

	if err := c.cases.DeleteAccident(ctx.Context(), accidentID); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error eliminando el accidente",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Accidente eliminado correctamente",
	})
}

// DeleteClassification godoc
// @Summary Delete a Classification
// @Description Delete a Classification
// @Tags classifications
// @Accept json
// @Produce json
// @Param id path string true "Classification ID"
// @Success 200 {string} string "Clasificación eliminada correctamente"
// @Failure 500 {string} string "Error eliminando la clasificación"
// @Router /classifications/{id} [delete]
func (c *controller) DeleteClassification(ctx *fiber.Ctx) error {
	classificationID := ctx.Params("id")

	if err := c.cases.DeleteClassification(ctx.Context(), classificationID); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error eliminando la clasificación",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Clasificación eliminada correctamente",
	})
}

// DeleteGravity godoc
// @Summary Delete a Gravity
// @Description Delete a Gravity
// @Tags gravities
// @Accept json
// @Produce json
// @Param id path string true "Gravity ID"
// @Success 200 {string} string "Gravedad eliminada correctamente"
// @Failure 500 {string} string "Error eliminando la gravedad"
// @Router /gravities/{id} [delete]
func (c *controller) DeleteGravity(ctx *fiber.Ctx) error {
	gravityID := ctx.Params("id")

	if err := c.cases.DeleteGravity(ctx.Context(), gravityID); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error eliminando la gravedad",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Gravedad eliminada correctamente",
	})
}

// DeleteInjuredPart godoc
// @Summary Delete a InjuredPart
// @Description Delete a InjuredPart
// @Tags injured-parts
// @Accept json
// @Produce json
// @Param id path string true "InjuredPart ID"
// @Success 200 {string} string "Parte lesionada eliminada correctamente"
// @Failure 500 {string} string "Error eliminando la parte lesionada"
// @Router /injured-parts/{id} [delete]
func (c *controller) DeleteInjuredPart(ctx *fiber.Ctx) error {
	injuredPartID := ctx.Params("id")

	if err := c.cases.DeleteInjuredPart(ctx.Context(), injuredPartID); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error eliminando la parte lesionada",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Parte lesionada eliminada correctamente",
	})
}

// GetAccidentById godoc
// @Summary Get an accident by id
// @Description Get an accident by id
// @Tags accidents
// @Accept json
// @Produce json
// @Param id path string true "Accident ID"
// @Success 200 {object} GetAccidentByIdResponse "Accidente obtenido correctamente"
// @Failure 500 {string} string "Error obteniendo el accidente"
// @Router /accidents/{id} [get]
func (c *controller) GetAccidentById(ctx *fiber.Ctx) error {
	accidentID := ctx.Params("id")

	result, err := c.cases.GetAccidentById(ctx.Context(), accidentID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error obteniendo el accidente",
		})
	}

	response := new(GetAccidentByIdResponse)
	response.ID = result.ID
	response.Description = result.Description
	response.ConstructionArea = result.ConstructionArea
	response.Classification = result.Classification
	response.Gravity = result.Gravity
	response.Project = result.Project
	response.Employee = result.Employee
	response.InjuredPart = result.InjuredPart
	response.AccidentDate = result.AccidentDate
	response.IsDeleted = result.IsDeleted
	response.CreatedAt = result.CreatedAt
	response.UpdatedAt = result.UpdatedAt
	response.DeletedAt = result.DeletedAt

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":  "Accidente obtenido correctamente",
		"accident": response,
	})
}

// GetAccidents godoc
// @Summary Get all accidents
// @Description Get all accidents
// @Tags accidents
// @Accept json
// @Produce json
// @Success 200 {object} GetAccidentByIdResponse "Accidentes obtenidos correctamente"
// @Failure 500 {string} string "Error obteniendo los accidentes"
// @Router /accidents/ [get]
func (c *controller) GetAccidents(ctx *fiber.Ctx) error {
	result, err := c.cases.GetAccidents(ctx.Context())
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error obteniendo los accidentes",
		})
	}

	response := make([]GetAccidentByIdResponse, 0)
	for _, accident := range result {
		response = append(response, GetAccidentByIdResponse{
			ID:               accident.ID,
			Description:      accident.Description,
			ConstructionArea: accident.ConstructionArea,
			Classification:   accident.Classification,
			Gravity:          accident.Gravity,
			Project:          accident.Project,
			Employee:         accident.Employee,
			InjuredPart:      accident.InjuredPart,
			AccidentDate:     accident.AccidentDate,
			IsDeleted:        accident.IsDeleted,
			CreatedAt:        accident.CreatedAt,
			UpdatedAt:        accident.UpdatedAt,
			DeletedAt:        accident.DeletedAt,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":   "Accidentes obtenidos correctamente",
		"accidents": response,
	})
}

// GetClassificationById godoc
// @Summary Get a Classification by id
// @Description Get a Classification by id
// @Tags classifications
// @Accept json
// @Produce json
// @Param id path string true "Classification ID"
// @Success 200 {object} GetClassificationByIdResponse "Clasificación obtenida correctamente"
// @Failure 500 {string} string "Error obteniendo la clasificación"
// @Router /classifications/{id} [get]
func (c *controller) GetClassificationById(ctx *fiber.Ctx) error {
	classificationID := ctx.Params("id")

	result, err := c.cases.GetClassificationById(ctx.Context(), classificationID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error obteniendo la clasificación",
		})
	}

	response := new(GetClassificationByIdResponse)
	response.ID = result.ID
	response.Name = result.Name
	response.IsActive = result.IsActive
	response.IsDeleted = result.IsDeleted
	response.CreatedAt = result.CreatedAt
	response.UpdatedAt = result.UpdatedAt
	response.DeletedAt = result.DeletedAt

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":        "Clasificación obtenida correctamente",
		"classification": response,
	})
}

// GetClassifications godoc
// @Summary Get all classifications
// @Description Get all classifications
// @Tags classifications
// @Accept json
// @Produce json
// @Success 200 {object} GetClassificationByIdResponse "Clasificaciones obtenidas correctamente"
// @Failure 500 {string} string "Error obteniendo las clasificaciones"
// @Router /classifications/ [get]
func (c *controller) GetClassifications(ctx *fiber.Ctx) error {
	result, err := c.cases.GetClassifications(ctx.Context())
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error obteniendo las clasificaciones",
		})
	}

	response := make([]GetClassificationByIdResponse, 0)
	for _, classification := range result {
		response = append(response, GetClassificationByIdResponse{
			ID:        classification.ID,
			Name:      classification.Name,
			IsActive:  classification.IsActive,
			IsDeleted: classification.IsDeleted,
			CreatedAt: classification.CreatedAt,
			UpdatedAt: classification.UpdatedAt,
			DeletedAt: classification.DeletedAt,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":         "Clasificaciones obtenidas correctamente",
		"classifications": response,
	})
}

// GetGravityById godoc
// @Summary Get a Gravity by id
// @Description Get a Gravity by id
// @Tags gravities
// @Accept json
// @Produce json
// @Param id path string true "Gravity ID"
// @Success 200 {object} GetGravityByIdResponse "Gravedad obtenida correctamente"
// @Failure 500 {string} string "Error obteniendo la gravedad"
// @Router /gravities/{id} [get]
func (c *controller) GetGravities(ctx *fiber.Ctx) error {
	result, err := c.cases.GetGravities(ctx.Context())
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error obteniendo las gravedades",
		})
	}

	response := make([]GetGravityByIdResponse, 0)
	for _, gravity := range result {
		response = append(response, GetGravityByIdResponse{
			ID:        gravity.ID,
			Name:      gravity.Name,
			IsActive:  gravity.IsActive,
			IsDeleted: gravity.IsDeleted,
			CreatedAt: gravity.CreatedAt,
			UpdatedAt: gravity.UpdatedAt,
			DeletedAt: gravity.DeletedAt,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":   "Gravedades obtenidas correctamente",
		"gravities": response,
	})
}

// GetInjuredPartById godoc
// @Summary Get a InjuredPart by id
// @Description Get a InjuredPart by id
// @Tags injured-parts
// @Accept json
// @Produce json
// @Param id path string true "InjuredPart ID"
// @Success 200 {object} GetInjuredPartByIdResponse "Parte lesionada obtenida correctamente"
// @Failure 500 {string} string "Error obteniendo la parte lesionada"
// @Router /injured-parts/{id} [get]
func (c *controller) GetGravityById(ctx *fiber.Ctx) error {
	gravityID := ctx.Params("id")

	result, err := c.cases.GetGravityById(ctx.Context(), gravityID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error obteniendo la gravedad",
		})
	}

	response := new(GetGravityByIdResponse)
	response.ID = result.ID
	response.Name = result.Name
	response.IsActive = result.IsActive
	response.IsDeleted = result.IsDeleted
	response.CreatedAt = result.CreatedAt
	response.UpdatedAt = result.UpdatedAt
	response.DeletedAt = result.DeletedAt

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Gravedad obtenida correctamente",
		"gravity": response,
	})
}

// GetInjuredPartById godoc
// @Summary Get a InjuredPart by id
// @Description Get a InjuredPart by id
// @Tags injured-parts
// @Accept json
// @Produce json
// @Param id path string true "InjuredPart ID"
// @Success 200 {object} GetInjuredPartByIdResponse "Parte lesionada obtenida correctamente"
// @Failure 500 {string} string "Error obteniendo la parte lesionada"
// @Router /injured-parts/{id} [get]
func (c *controller) GetInjuredPartById(ctx *fiber.Ctx) error {
	injuredPartID := ctx.Params("id")

	result, err := c.cases.GetInjuredPartById(ctx.Context(), injuredPartID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error obteniendo la parte lesionada",
		})
	}

	response := new(GetInjuredPartByIdResponse)
	response.ID = result.ID
	response.Name = result.Name
	response.IsActive = result.IsActive
	response.IsDeleted = result.IsDeleted
	response.CreatedAt = result.CreatedAt
	response.UpdatedAt = result.UpdatedAt
	response.DeletedAt = result.DeletedAt

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":      "Parte lesionada obtenida correctamente",
		"injured_part": response,
	})
}

// GetInjuredParts godoc
// @Summary Get all injured parts
// @Description Get all injured parts
// @Tags injured-parts
// @Accept json
// @Produce json
// @Success 200 {object} GetInjuredPartByIdResponse "Partes lesionadas obtenidas correctamente"
// @Failure 500 {string} string "Error obteniendo las partes lesionadas"
// @Router /injured-parts/ [get]
func (c *controller) GetInjuredParts(ctx *fiber.Ctx) error {
	result, err := c.cases.GetInjuredParts(ctx.Context())
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error obteniendo las partes lesionadas",
		})
	}

	response := make([]GetInjuredPartByIdResponse, 0)
	for _, injuredPart := range result {
		response = append(response, GetInjuredPartByIdResponse{
			ID:        injuredPart.ID,
			Name:      injuredPart.Name,
			IsActive:  injuredPart.IsActive,
			IsDeleted: injuredPart.IsDeleted,
			CreatedAt: injuredPart.CreatedAt,
			UpdatedAt: injuredPart.UpdatedAt,
			DeletedAt: injuredPart.DeletedAt,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":       "Partes lesionadas obtenidas correctamente",
		"injured_parts": response,
	})
}

// UpdateAccident godoc
// @Summary Update an accident
// @Description Update an accident
// @Tags accidents
// @Accept json
// @Produce json
// @Param id path string true "Accident ID"
// @Param body body UpdateAccidentRequest true "Update Accident Request"
// @Success 200 {string} string "Accidente actualizado correctamente"
// @Failure 400 {string} string "Error validando el cuerpo de la petición"
// @Failure 500 {string} string "Error actualizando el accidente"
// @Router /accidents/{id} [put]
func (c *controller) UpdateAccident(ctx *fiber.Ctx) error {
	accidentID := ctx.Params("id")
	requestBody := new(UpdateAccidentRequest)

	if err := utils.BodyParser(ctx, requestBody); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error validando el cuerpo de la petición :" + err.Error(),
		})
	}

	accident := new(Accident)
	accident.Description = requestBody.Description
	accident.ConstructionArea = requestBody.ConstructionArea
	accident.Classification.ID = requestBody.ClassificationID
	accident.Gravity.ID = requestBody.GravityID
	accident.Project.ID = requestBody.ProjectID
	accident.Employee.ID = requestBody.EmployeeID
	accident.InjuredPart.ID = requestBody.InjuredPartID
	accident.AccidentDate = requestBody.AccidentDate

	if err := c.cases.UpdateAccident(ctx.Context(), accidentID, accident); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error actualizando el accidente",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Accidente actualizado correctamente",
	})
}

// UpdateClassification godoc
// @Summary Update a Classification
// @Description Update a Classification
// @Tags classifications
// @Accept json
// @Produce json
// @Param id path string true "Classification ID"
// @Param body body UpdateClassificationRequest true "Update Classification Request"
// @Success 200 {string} string "Clasificación actualizada correctamente"
// @Failure 400 {string} string "Error validando el cuerpo de la petición"
// @Failure 500 {string} string "Error actualizando la clasificación"
// @Router /classifications/{id} [put]
func (c *controller) UpdateClassification(ctx *fiber.Ctx) error {
	classificationID := ctx.Params("id")
	requestBody := new(UpdateClassificationRequest)

	if err := utils.BodyParser(ctx, requestBody); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error validando el cuerpo de la petición :" + err.Error(),
		})
	}

	classification := new(Classification)
	classification.Name = requestBody.Name

	if err := c.cases.UpdateClassification(ctx.Context(), classificationID, classification); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error actualizando la clasificación",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Clasificación actualizada correctamente",
	})
}

// UpdateGravity godoc
// @Summary Update a Gravity
// @Description Update a Gravity
// @Tags gravities
// @Accept json
// @Produce json
// @Param id path string true "Gravity ID"
// @Param body body UpdateGravityRequest true "Update Gravity Request"
// @Success 200 {string} string "Gravedad actualizada correctamente"
// @Failure 400 {string} string "Error validando el cuerpo de la petición"
// @Failure 500 {string} string "Error actualizando la gravedad"
// @Router /gravities/{id} [put]
func (c *controller) UpdateGravity(ctx *fiber.Ctx) error {
	gravityID := ctx.Params("id")
	requestBody := new(UpdateGravityRequest)

	if err := utils.BodyParser(ctx, requestBody); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error validando el cuerpo de la petición :" + err.Error(),
		})
	}

	gravity := new(Gravity)
	gravity.Name = requestBody.Name

	if err := c.cases.UpdateGravity(ctx.Context(), gravityID, gravity); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error actualizando la gravedad",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Gravedad actualizada correctamente",
	})
}

// UpdateInjuredPart godoc
// @Summary Update a InjuredPart
// @Description Update a InjuredPart
// @Tags injured-parts
// @Accept json
// @Produce json
// @Param id path string true "InjuredPart ID"
// @Param body body UpdateInjuredPartRequest true "Update InjuredPart Request"
// @Success 200 {string} string "Parte lesionada actualizada correctamente"
// @Failure 400 {string} string "Error validando el cuerpo de la petición"
// @Failure 500 {string} string "Error actualizando la parte lesionada"
// @Router /injured-parts/{id} [put]
func (c *controller) UpdateInjuredPart(ctx *fiber.Ctx) error {
	injuredPartID := ctx.Params("id")
	requestBody := new(UpdateInjuredPartRequest)

	if err := utils.BodyParser(ctx, requestBody); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error validando el cuerpo de la petición :" + err.Error(),
		})
	}

	injuredPart := new(InjuredPart)
	injuredPart.Name = requestBody.Name

	if err := c.cases.UpdateInjuredPart(ctx.Context(), injuredPartID, injuredPart); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error actualizando la parte lesionada",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Parte lesionada actualizada correctamente",
	})
}
