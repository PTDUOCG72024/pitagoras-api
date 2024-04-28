package users

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
	group := app.Group("/auth", utils.GetNextMiddleWare)
	group.Post("/register", ctrl.Register)                // register new user
	group.Post("/login", ctrl.Login)                      // login user
	group.Get("/user/:id", ctrl.GetUserById)              // get user by id
	group.Get("/user/email/:email", ctrl.GetUserByEmail)  // get user by email
	group.Put("/user/activate/:id", ctrl.Activate)        // activate user
	group.Delete("/user/deactivate/:id", ctrl.Deactivate) // deactivate user
	group.Put("/user/:id", ctrl.Update)                   // update user
}
