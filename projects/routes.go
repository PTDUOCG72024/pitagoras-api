package projects

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
	group := app.Group("/projects", utils.GetNextMiddleWare)
	group.Post("/", ctrl.CreateProject)                  // Create a new project
	group.Get("/:id", ctrl.GetProjectById)               // Get a project by id
	group.Put("/:id", ctrl.UpdateProject)                // Update a project
	group.Delete("/:id", ctrl.DeleteProject)             // Delete a project
	group.Get("/", ctrl.GetProjects)                     // Get all projects
	group.Put("/activate/:id", ctrl.ActivateProject)     // Activate a project
	group.Put("/deactivate/:id", ctrl.DeactivateProject) // Deactivate a project
}
