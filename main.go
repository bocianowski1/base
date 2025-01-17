package main

import (
	"log"

	"github.com/KristianElde/butte/handlers"
	"github.com/KristianElde/butte/models"
	"github.com/KristianElde/butte/repo"
	"github.com/KristianElde/butte/services"
	"github.com/KristianElde/butte/util"
	"github.com/KristianElde/butte/views"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("dev.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	if util.IsDev() {
		log.Println("Dropping tables in dev mode")
		db.Migrator().DropTable(&models.User{})
	}

	db.AutoMigrate(&models.User{})
	log.Println("Database migrated")

	app := fiber.New(fiber.Config{
		Views:             html.New("./views", ".html"),
		PassLocalsToViews: true,
	})

	// common middlewares
	app.Use(logger.New())

	// repos
	userRepo := repo.NewUserRepo(db)

	// services
	userSrv := services.NewUserService(userRepo)
	authSrv := services.NewAuthService()

	// handlers
	userHandler := handlers.NewUserHandler(userSrv)
	authHandler := handlers.NewAuthHandler(userSrv, authSrv)

	// routes
	app.Get("/", func(c *fiber.Ctx) error {
		return views.Render(c, "index", fiber.Map{"Title": "Home"})
	})

	app.Get("/auth/register", func(c *fiber.Ctx) error {
		return views.Render(c, "auth/register", fiber.Map{"Title": "Register"})
	})
	app.Post("/auth/register", userHandler.HandleCreate)

	app.Get("/auth/login", func(c *fiber.Ctx) error {
		return views.Render(c, "auth/login", fiber.Map{"Title": "Login"})
	})
	app.Post("/auth/login", authHandler.HandleLogin)

	app.Static("/static", "./static")
	log.Fatal(app.Listen(":3000"))
}
