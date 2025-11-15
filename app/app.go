package app

import (
	"take-home-test/app/controllers"
	"take-home-test/app/repositories"
	"take-home-test/app/routes"
	usecase "take-home-test/app/usecases"
	"take-home-test/pkg/config"
	"take-home-test/pkg/database"

	_ "take-home-test/docs" // ✅ PASTIKAN INI ADA

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type Main struct {
	cfg        *config.Config
	database   Database
	repo       *repositories.Main
	usecase    *usecase.Main
	controller *controllers.Main
	router     *fiber.App
}

type Database struct {
	MySQL    *gorm.DB
	Postgres *gorm.DB
}

func New() *Main {
	return new(Main)
}

func (m *Main) Init() (err error) {
	// Load environment variables
	viper.SetConfigFile(".env")
	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	// Load configuration
	m.cfg = config.NewConfig()

	// Initialize Fiber app
	app := fiber.New(fiber.Config{
		AppName: "Sports Booking API - " + m.cfg.ServiceEnvironment,
	})

	// Middleware
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET, POST, PUT, DELETE, PATCH, OPTIONS",
	}))

	// Database connection - menggunakan config Postgres
	m.database.Postgres, err = database.GetConnection(m.cfg.Postgres().Read.ToArgs(database.Postgres, database.ReadConn, nil))
	if err != nil {
		return
	}

	// Initialize layers
	m.repo = repositories.Init(repositories.Options{
		Postgres: m.database.Postgres,
		Config:   m.cfg,
	})

	m.usecase = usecase.Init(usecase.Options{
		Repository: m.repo,
		Config:     m.cfg,
	})

	m.controller = controllers.Init(controllers.Options{
		UseCases: m.usecase,
		Config:   m.cfg,
	})

	m.router = app

	// Configure routes
	routes.ConfigureRouter(app, m.controller)
	return err
}

func (m *Main) Run() (err error) {
	defer m.close()

	// Start server
	err = m.router.Listen(":" + m.cfg.ServicePort)
	return
}

func (m *Main) close() {
	if m.database.MySQL != nil {
		if db, err := m.database.MySQL.DB(); err == nil {
			db.Close()
		}
	}

	if m.database.Postgres != nil {
		if db, err := m.database.Postgres.DB(); err == nil {
			db.Close()
		}
	}
}
