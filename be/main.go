package main

import (
	"log"
	"my-fiber-app/handlers"
	"my-fiber-app/middlewares"
	"my-fiber-app/schema"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/jmoiron/sqlx"
	_ "modernc.org/sqlite"
)

func main() {
	db, err := sqlx.Connect("sqlite", "database.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"Origin, Content-Type, Accept, Authorization"},
		AllowMethods: []string{"GET, POST, PUT, DELETE, OPTIONS"},
	}))
	schema.CreateTables(db)

	app.Get("/", handlers.HelloHandler)
	app.Post("/create-room", handlers.CreateRoomHandler(db))
	app.Post("/create-subjects", handlers.CreateSubjectHandler(db))
	app.Post("/create-sub-token", handlers.CreateSubToken(db))

	admin := app.Group("/admin", middlewares.IsAdminMiddleware)
	admin.Post("/create-sub-token", handlers.CreateSubToken(db))

	log.Fatal(app.Listen(":4200"))
}
