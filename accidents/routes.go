package accidents

import (
	"github.com/gofiber/fiber/v2"
	"github.com/xbizzybone/pitagoras-api/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

var repo Repository
var ctrl Controller
var cs Cases

func bootstrap(logger *zap.Logger, accidentsCollection *mongo.Collection, classificationsCollection *mongo.Collection, gravitiesCollection *mongo.Collection, injuredPartsCollection *mongo.Collection, projectsCollection *mongo.Collection, employeesCollection *mongo.Collection) {
	repo = NewRepository(logger, accidentsCollection, classificationsCollection, gravitiesCollection, injuredPartsCollection, projectsCollection, employeesCollection)
	cs = NewCases(logger, repo)
	ctrl = NewController(cs)
}

func ApplyRoutes(app *fiber.App, logger *zap.Logger, accidentsCollection *mongo.Collection, classificationsCollection *mongo.Collection, gravitiesCollection *mongo.Collection, injuredPartsCollection *mongo.Collection, projectsCollection *mongo.Collection, employeesCollection *mongo.Collection) {
	bootstrap(logger, accidentsCollection, classificationsCollection, gravitiesCollection, injuredPartsCollection, projectsCollection, employeesCollection)
	group := app.Group("/accidents", utils.GetNextMiddleWare)
	group.Post("/", ctrl.CreateAccident)      // Create a new accident
	group.Get("/:id", ctrl.GetAccidentById)   // Get an accident by id
	group.Put("/:id", ctrl.UpdateAccident)    // Update an accident
	group.Delete("/:id", ctrl.DeleteAccident) // Delete an accident
	group.Get("/", ctrl.GetAccidents)         // Get all accidents

	group.Put("/updates/update-all-employee", ctrl.UpdateAllAccidentsEmployee) // Update all accidents employee

	group = app.Group("/classifications", utils.GetNextMiddleWare)
	group.Post("/", ctrl.CreateClassification)                  // Create a new Classification
	group.Get("/:id", ctrl.GetClassificationById)               // Get a Classification by id
	group.Get("/", ctrl.GetClassifications)                     // Get all Classifications
	group.Put("/:id", ctrl.UpdateClassification)                // Update a Classification
	group.Delete("/:id", ctrl.DeleteClassification)             // Delete a Classification
	group.Put("/activate/:id", ctrl.ActivateClassification)     // Activate a Classification
	group.Put("/deactivate/:id", ctrl.DeactivateClassification) // Deactivate a Classification

	group = app.Group("/gravities", utils.GetNextMiddleWare)
	group.Post("/", ctrl.CreateGravity)                  // Create a new Gravity
	group.Get("/:id", ctrl.GetGravityById)               // Get a Gravity by id
	group.Get("/", ctrl.GetGravities)                    // Get all Gravities
	group.Put("/:id", ctrl.UpdateGravity)                // Update a Gravity
	group.Delete("/:id", ctrl.DeleteGravity)             // Delete a Gravity
	group.Put("/activate/:id", ctrl.ActivateGravity)     // Activate a Gravity
	group.Put("/deactivate/:id", ctrl.DeactivateGravity) // Deactivate a Gravity

	group = app.Group("/injured-parts", utils.GetNextMiddleWare)
	group.Post("/", ctrl.CreateInjuredPart)                  // Create a new InjuredPart
	group.Get("/:id", ctrl.GetInjuredPartById)               // Get a InjuredPart by id
	group.Get("/", ctrl.GetInjuredParts)                     // Get all InjuredParts
	group.Put("/:id", ctrl.UpdateInjuredPart)                // Update a InjuredPart
	group.Delete("/:id", ctrl.DeleteInjuredPart)             // Delete a InjuredPart
	group.Put("/activate/:id", ctrl.ActivateInjuredPart)     // Activate a InjuredPart
	group.Put("/deactivate/:id", ctrl.DeactivateInjuredPart) // Deactivate a InjuredPart
}
