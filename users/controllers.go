package users

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

// Register godoc
//
//	@Summary		register new user
//	@Description 	register new user
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			body body UserCreateRequest true "UserCreateRequest"
//	@Success		200 {object} UserCreateResponse
//	@Failure		400 {string} string "Error validando el cuerpo de la petición"
//	@Failure		409 {string} string "El usuario ya existe"
//	@Failure		500 {string} string "Error encriptando la contraseña"
//	@Failure		500 {string} string "Error creando el usuario"
//	@Router			/auth/register [post]
func (c *controller) Register(ctx *fiber.Ctx) error {
	requestBody := new(UserCreateRequest)

	if err := utils.BodyParser(ctx, requestBody); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error validando el cuerpo de la petición :" + err.Error(),
		})
	}

	user := new(User)
	user.Email = requestBody.Email
	user.Password = requestBody.Password
	user.Name = requestBody.Name

	createUser, err := c.cases.Register(ctx.UserContext(), user)
	if err != nil {
		if err.Error() == "user already exists" {
			return ctx.Status(fiber.StatusConflict).JSON(fiber.Map{
				"message": "El usuario ya existe",
			})
		}

		if err.Error() == "error while hashing password" {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Error encriptando la contraseña",
			})
		}

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error creando el usuario",
		})
	}

	userResponse := new(UserCreateResponse)
	userResponse.Name = createUser.Name
	userResponse.Email = createUser.Email
	userResponse.ID = createUser.ID

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Usuario creado correctamente",
		"user":    userResponse,
	})
}

// Login godoc
//
//	@Summary		login user
//	@Description 	login user
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			body body UserRequest true "UserRequest"
//	@Success		200 {object} UserLoginResponse
//	@Failure		400 {string} string "Error validando el cuerpo de la petición"
//	@Failure		404 {string} string "El usuario no existe"
//	@Failure		401 {string} string "Contraseña incorrecta"
//	@Failure		500 {string} string "Error iniciando sesión"
//	@Failure		401 {string} string "Cuenta de usuario deshabilitada"
//	@Router			/auth/login [post]
func (c *controller) Login(ctx *fiber.Ctx) error {
	requestBody := new(UserRequest)

	if err := utils.BodyParser(ctx, requestBody); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error validando el cuerpo de la petición :" + err.Error(),
		})
	}

	user := new(User)
	user.Email = requestBody.Email
	user.Password = requestBody.Password

	userResult, err := c.cases.Login(ctx.UserContext(), user)
	if err != nil {
		if err.Error() == "user not found" {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "El usuario no existe",
			})
		}

		if err.Error() == "invalid password" {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Contraseña incorrecta",
			})
		}

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error iniciando sesión",
		})
	}

	if userResult.IsDeleted {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Cuenta de usuario deshabilitada",
		})
	}

	response := new(UserLoginResponse)
	response.Email = userResult.Email
	response.Name = userResult.Name
	response.ID = userResult.ID
	response.LastLogin = userResult.LastLogin
	response.CreatedAt = userResult.CreatedAt
	response.UpdatedAt = userResult.UpdatedAt

	return ctx.Status(fiber.StatusOK).JSON(response)
}

// GetUserById godoc
//
//	@Summary		get user by id
//	@Description 	get user by id
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			id path string true "User ID"
//	@Success		200 {object} UserResponse
//	@Failure		400 {string} string "Usuario no existe"
//	@Failure		404 {string} string "El usuario no existe"
//	@Failure		500 {string} string "Error obteniendo el usuario"
//	@Router			/auth/user/{id} [get]
func (c *controller) GetUserById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	userResult, err := c.cases.GetUserById(ctx.UserContext(), id)
	if err != nil {
		if err.Error() == "invalid id" {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Usuario no existe",
			})
		}

		if err.Error() == "user not found" {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "El usuario no existe",
			})
		}

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error obteniendo el usuario",
		})
	}

	response := new(UserResponse)
	response.Email = userResult.Email
	response.Name = userResult.Name
	response.ID = userResult.ID
	response.IsDeleted = userResult.IsDeleted
	response.LastLogin = userResult.LastLogin
	response.CreatedAt = userResult.CreatedAt
	response.UpdatedAt = userResult.UpdatedAt
	response.DeleteAt = userResult.DeleteAt

	return ctx.Status(fiber.StatusOK).JSON(response)
}

// GetUserByEmail godoc
//
//	@Summary		get user by email
//	@Description 	get user by email
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			email path string true "User Email"
//	@Success		200 {object} UserResponse
//	@Failure		404 {string} string "El usuario no existe"
//	@Failure		500 {string} string "Error obteniendo el usuario"
//	@Router			/auth/user/email/{email} [get]
func (c *controller) GetUserByEmail(ctx *fiber.Ctx) error {
	email := ctx.Params("email")

	userResult, err := c.cases.GetUserByEmail(ctx.UserContext(), email)
	if err != nil {
		if err.Error() == "user not found" {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "El usuario no existe",
			})
		}

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error obteniendo el usuario",
		})
	}

	response := new(UserResponse)
	response.Email = userResult.Email
	response.Name = userResult.Name
	response.ID = userResult.ID
	response.IsDeleted = userResult.IsDeleted
	response.LastLogin = userResult.LastLogin
	response.CreatedAt = userResult.CreatedAt
	response.UpdatedAt = userResult.UpdatedAt
	response.DeleteAt = userResult.DeleteAt

	return ctx.Status(fiber.StatusOK).JSON(response)
}

// Deactivate godoc
//
//	@Summary		deactivate user
//	@Description 	deactivate user
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			id path string true "User ID"
//	@Success		200 {string} string "Usuario desactivado correctamente"
//	@Failure		400 {string} string "Usuario no existe"
//	@Failure		404 {string} string "El usuario no existe"
//	@Failure		500 {string} string "Error eliminando el usuario"
//	@Router			/auth/user/deactivate/{id} [delete]
func (c *controller) Deactivate(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	err := c.cases.Deactivate(ctx.UserContext(), id)
	if err != nil {
		if err.Error() == "invalid id" {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Usuario no existe",
			})
		}

		if err.Error() == "user not found" {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "El usuario no existe",
			})
		}

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error eliminando el usuario",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Usuario desactivado correctamente",
	})
}

// Activate godoc
//
//	@Summary		activate user
//	@Description 	activate user
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			id path string true "User ID"
//	@Success		200 {string} string "Usuario activado correctamente"
//	@Failure		400 {string} string "Usuario no existe"
//	@Failure		404 {string} string "El usuario no existe"
//	@Failure		500 {string} string "Error activando el usuario"
//	@Router			/auth/user/activate/{id} [put]
func (c *controller) Activate(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	err := c.cases.Activate(ctx.UserContext(), id)
	if err != nil {
		if err.Error() == "invalid id" {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Usuario no existe",
			})
		}

		if err.Error() == "user not found" {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "El usuario no existe",
			})
		}

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error activando el usuario",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Usuario activado correctamente",
	})
}

// Update godoc
//
//	@Summary		update user
//	@Description 	update user
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			id path string true "User ID"
//	@Param			body body UserUpdateRequest true "UserUpdateRequest"
//	@Success		200 {string} string "Usuario actualizado correctamente"
//	@Failure		400 {string} string "Usuario no existe"
//	@Failure		404 {string} string "El usuario no existe"
//	@Failure		500 {string} string "Error actualizando el usuario"
//	@Router			/auth/user/{id} [put]
func (c *controller) Update(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	requestBody := new(UserUpdateRequest)
	err := ctx.BodyParser(requestBody)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error parseando el body",
		})
	}

	if requestBody.Email != "" && !utils.IsValidEmail(requestBody.Email) {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Email invalido",
		})
	}

	user := new(User)
	user.Email = requestBody.Email
	user.Name = requestBody.Name
	user.Password = requestBody.Password

	err = c.cases.Update(ctx.UserContext(), id, user)
	if err != nil {
		if err.Error() == "invalid id" {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Usuario no existe",
			})
		}

		if err.Error() == "user not found" {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "El usuario no existe",
			})
		}

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error actualizando el usuario",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Usuario actualizado correctamente",
	})
}
