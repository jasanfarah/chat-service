package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"jasanfarah/chat-service/internal/config"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"jasanfarah/chat-service/internal/api"
	"jasanfarah/chat-service/internal/api/controllers"
	"jasanfarah/chat-service/internal/api/validation"
	"jasanfarah/chat-service/internal/database"
	"jasanfarah/chat-service/internal/migrations"
	"jasanfarah/chat-service/internal/repositories"
	"jasanfarah/chat-service/internal/services"
	customMiddleware "jasanfarah/chat-service/internal/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"gorm.io/gorm"
)

func main() {
	appConfig := config.LoadConfig("local")

	db, err := database.NewDatabase(database.PostgresConfig{
		Host:     appConfig.Database.Host,
		User:     appConfig.Database.User,
		Password: appConfig.Database.Password,
		DBName:   appConfig.Database.DBName,
		Port:     appConfig.Database.Port,
		SSLMode:  appConfig.Database.SSLMode,
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}

	err = migrations.Migrate(db.Connection)
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
		return
	}

	app := fiber.New()
	app.Use(recover.New())
	app.Use(customMiddleware.ErrorHandler())

	customValidator := validator.New()
	validation.RegisterValidators(customValidator)

	initalizeModules(app, db.Connection, logger, customValidator)

	err = app.Listen(":6500")
	if err != nil {
		_, _ = gin.DefaultErrorWriter.Write([]byte(fmt.Sprintf("Failed starting application: %v", err.Error())))

	}
}

func initalizeModules(app *fiber.App, connection *gorm.DB, log *zap.Logger, customValidator *validator.Validate) {


	// Initialize repositories
	conversationRepository := repositories.NewConversationRepository(connection, log)
	messageRepository := repositories.NewMessageRepository(connection, log)


	// Initialize services
	conversationService := services.NewConversationService(services.ConversationServiceOptions{
		Repository: conversationRepository,
		Logger:                 log,
	})
	
	messageService := services.NewMessageService(services.MessageServiceOptions{
		Repository: messageRepository,
		Logger:                 log,
		ConversationRepostory: conversationRepository,
	})



	// Initialize controllers
	conversationController := controllers.NewConversationController(controllers.ConversationControllerOptions{
		ConversationService: conversationService,
		Logger:              log,
		Validator:           customValidator,
	})
	messageController := controllers.NewMessageController(controllers.MessageControllerOptions{
		MessageService: messageService,
		Logger:  log,
		Validation: customValidator,
	})




	api.InitializeAPIRoutes(app, log, conversationController, messageController, customValidator)

}

type flags struct {
	Seed           bool
	EnableDatabase bool
	EnableAuth     bool
}

func parseFlags() flags {
	fs := flag.NewFlagSet("chat-service", flag.ExitOnError)

	// Declare the -seed flag. This flag does not need a value; its presence activates the mode.
	seedPresent := false
	fs.BoolVar(&seedPresent, "seed", false, "Seeds")

	enableDatabase := false
	fs.BoolVar(&enableDatabase, "database", false, "Enables database startup")

	enableAuth := true
	fs.BoolVar(&enableAuth, "auth", true, "Enables authentication middleware")

	// Parse the flags
	err := fs.Parse(os.Args[1:])
	if err != nil {
		fmt.Printf("Error parsing flags: %s\n", err)
		os.Exit(1)
	}

	// Initialize flags struct
	var myFlags = flags{
		Seed:           seedPresent,
		EnableDatabase: enableDatabase,
		EnableAuth:     enableAuth,
	}

	return myFlags
}
