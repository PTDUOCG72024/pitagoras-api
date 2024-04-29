package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
	"github.com/joho/godotenv"
	"github.com/xbizzybone/pitagoras-api/employees"
	"github.com/xbizzybone/pitagoras-api/projects"
	"github.com/xbizzybone/pitagoras-api/users"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	_ "github.com/xbizzybone/pitagoras-api/docs"
)

var client *mongo.Client
var usersCollection *mongo.Collection
var projectsCollection *mongo.Collection
var employeesCollection *mongo.Collection
var nationalitiesCollection *mongo.Collection
var positionsCollection *mongo.Collection
var supervisorsCollection *mongo.Collection

var zapLogger *zap.Logger

func init() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	logConfig := zap.NewProductionConfig()
	logConfig.EncoderConfig.TimeKey = "timestamp"
	logConfig.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339)

	zapLogger, err = logConfig.Build()
	if err != nil {
		log.Fatal(err)
	}

	zapLogger = zapLogger.With(zap.String("service", "pitagoras-api"))
	zapLogger.Info("Logger initialized")

	clientOptions := options.Client().ApplyURI(os.Getenv("MONGO_CONNECTION_STRING"))
	client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	dbName := os.Getenv("MONGO_DB_NAME")
	usersCollection = client.Database(dbName).Collection(os.Getenv("MONGO_USERS_COLLECTION_NAME"))
	projectsCollection = client.Database(dbName).Collection(os.Getenv("MONGO_PROJECTS_COLLECTION_NAME"))
	employeesCollection = client.Database(dbName).Collection(os.Getenv("MONGO_EMPLOYEES_COLLECTION_NAME"))
	nationalitiesCollection = client.Database(dbName).Collection(os.Getenv("MONGO_NATIONALITIES_COLLECTION_NAME"))
	positionsCollection = client.Database(dbName).Collection(os.Getenv("MONGO_POSITIONS_COLLECTION_NAME"))
	supervisorsCollection = client.Database(dbName).Collection(os.Getenv("MONGO_SUPERVISORS_COLLECTION_NAME"))
	zapLogger.Info("MongoDB initialized")
}

// @title GoFiber Pitagoras API
// @version 1.0
// @description Golang GoFiber swagger auto generate step by step by swaggo
// @termsOfService http://swagger.io/terms/
// @contact.name xbizzybone
// @contact.email xbizzybone@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host pitagoras-api-production.up.railway.app
// @BasePath /
func main() {
	defer zapLogger.Sync()

	zapLogger.Info("Starting server")

	app := fiber.New()
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(cors.New())

	app.Get("/swagger/*", swagger.HandlerDefault) // default

	users.ApplyRoutes(app, zapLogger, usersCollection)
	projects.ApplyRoutes(app, zapLogger, projectsCollection)
	employees.ApplyRoutes(app, zapLogger, employeesCollection, nationalitiesCollection, positionsCollection, supervisorsCollection)

	app.Listen(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
