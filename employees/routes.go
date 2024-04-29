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

func bootstrap(logger *zap.Logger, mongoCollection *mongo.Collection) {
	repo = NewRepository(logger, mongoCollection)
	cs = NewCases(logger, repo)
	ctrl = NewController(cs)
}

func ApplyRoutes(app *fiber.App, logger *zap.Logger, mongoCollection *mongo.Collection) {
	bootstrap(logger, mongoCollection)
	group := app.Group("/employees", utils.GetNextMiddleWare)
	group.Post("/", ctrl.CreateEmployee)                  // Create a new employee
	group.Get("/:id", ctrl.GetEmployeeById)               // Get an employee by id
	group.Put("/:id", ctrl.UpdateEmployee)                // Update an employee
	group.Delete("/:id", ctrl.DeleteEmployee)             // Delete an employee
	group.Get("/", ctrl.GetEmployees)                     // Get all employees
	group.Put("/activate/:id", ctrl.ActivateEmployee)     // Activate an employee
	group.Put("/deactivate/:id", ctrl.DeactivateEmployee) // Deactivate an employee

	group.Post("/nationalities", ctrl.CreateNationality)                   // Create a new Nationality
	group.Get("/nationalities/:id", ctrl.GetNationalityById)               // Get a Nationality by id
	group.Get("/nationalities", ctrl.GetNationalities)                     // Get all Nationalities
	group.Put("/nationalities/:id", ctrl.UpdateNationality)                // Update a Nationality
	group.Delete("/nationalities/:id", ctrl.DeleteNationality)             // Delete a Nationality
	group.Put("/nationalities/activate/:id", ctrl.ActivateNationality)     // Activate a Nationality
	group.Put("/nationalities/deactivate/:id", ctrl.DeactivateNationality) // Deactivate a Nationality

	group.Post("/positions", ctrl.CreatePosition)                   // Create a new Position
	group.Get("/positions/:id", ctrl.GetPositionById)               // Get a Position by id
	group.Get("/positions", ctrl.GetPositions)                      // Get all Positions
	group.Put("/positions/:id", ctrl.UpdatePosition)                // Update a Position
	group.Delete("/positions/:id", ctrl.DeletePosition)             // Delete a Position
	group.Put("/positions/activate/:id", ctrl.ActivatePosition)     // Activate a Position
	group.Put("/positions/deactivate/:id", ctrl.DeactivatePosition) // Deactivate a Position

	group.Post("/supervisors", ctrl.CreateSupervisor)                   // Create a new Supervisor
	group.Get("/supervisors/:id", ctrl.GetSupervisorById)               // Get a Supervisor by id
	group.Get("/supervisors", ctrl.GetSupervisors)                      // Get all Supervisors
	group.Put("/supervisors/:id", ctrl.UpdateSupervisor)                // Update a Supervisor
	group.Delete("/supervisors/:id", ctrl.DeleteSupervisor)             // Delete a Supervisor
	group.Put("/supervisors/activate/:id", ctrl.ActivateSupervisor)     // Activate a Supervisor
	group.Put("/supervisors/deactivate/:id", ctrl.DeactivateSupervisor) // Deactivate a Supervisor
}
