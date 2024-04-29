package employees

import (
	"github.com/gofiber/fiber/v2"
	"github.com/xbizzybone/pitagoras-api/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

var repo Repository
var ctrl Controller
var cs Cases

func bootstrap(logger *zap.Logger, employeesCollection *mongo.Collection, nationalitiesCollection *mongo.Collection, positionsCollection *mongo.Collection, supervisorsCollection *mongo.Collection) {
	repo = NewRepository(logger, employeesCollection, nationalitiesCollection, positionsCollection, supervisorsCollection)
	cs = NewCases(logger, repo)
	ctrl = NewController(cs)
}

func ApplyRoutes(app *fiber.App, logger *zap.Logger, employeesCollection *mongo.Collection, nationalitiesCollection *mongo.Collection, positionsCollection *mongo.Collection, supervisorsCollection *mongo.Collection) {
	bootstrap(logger, employeesCollection, nationalitiesCollection, positionsCollection, supervisorsCollection)
	group := app.Group("/employees", utils.GetNextMiddleWare)
	group.Post("/", ctrl.CreateEmployee)                  // Create a new employee
	group.Get("/:id", ctrl.GetEmployeeById)               // Get an employee by id
	group.Put("/:id", ctrl.UpdateEmployee)                // Update an employee
	group.Delete("/:id", ctrl.DeleteEmployee)             // Delete an employee
	group.Get("/", ctrl.GetEmployees)                     // Get all employees
	group.Put("/activate/:id", ctrl.ActivateEmployee)     // Activate an employee
	group.Put("/deactivate/:id", ctrl.DeactivateEmployee) // Deactivate an employee

	group = app.Group("/nationalities", utils.GetNextMiddleWare)
	group.Post("/", ctrl.CreateNationality)                  // Create a new Nationality
	group.Get("/:id", ctrl.GetNationalityById)               // Get a Nationality by id
	group.Get("/", ctrl.GetNationalities)                    // Get all Nationalities
	group.Put("/:id", ctrl.UpdateNationality)                // Update a Nationality
	group.Delete("/:id", ctrl.DeleteNationality)             // Delete a Nationality
	group.Put("/activate/:id", ctrl.ActivateNationality)     // Activate a Nationality
	group.Put("/deactivate/:id", ctrl.DeactivateNationality) // Deactivate a Nationality

	group = app.Group("/positions", utils.GetNextMiddleWare)
	group.Post("/", ctrl.CreatePosition)                  // Create a new Position
	group.Get("/:id", ctrl.GetPositionById)               // Get a Position by id
	group.Get("/", ctrl.GetPositions)                     // Get all Positions
	group.Put("/:id", ctrl.UpdatePosition)                // Update a Position
	group.Delete("/:id", ctrl.DeletePosition)             // Delete a Position
	group.Put("/activate/:id", ctrl.ActivatePosition)     // Activate a Position
	group.Put("/deactivate/:id", ctrl.DeactivatePosition) // Deactivate a Position

	group = app.Group("/supervisors", utils.GetNextMiddleWare)
	group.Post("/", ctrl.CreateSupervisor)                  // Create a new Supervisor
	group.Get("/:id", ctrl.GetSupervisorById)               // Get a Supervisor by id
	group.Get("/", ctrl.GetSupervisors)                     // Get all Supervisors
	group.Put("/:id", ctrl.UpdateSupervisor)                // Update a Supervisor
	group.Delete("/:id", ctrl.DeleteSupervisor)             // Delete a Supervisor
	group.Put("/activate/:id", ctrl.ActivateSupervisor)     // Activate a Supervisor
	group.Put("/deactivate/:id", ctrl.DeactivateSupervisor) // Deactivate a Supervisor
}
