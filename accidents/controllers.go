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
