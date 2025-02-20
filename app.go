package main

import (
	"binvault/handlers"

	"flag"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

var (
	port = flag.String("port", ":3000", "Port to listen on")
	prod = flag.Bool("prod", false, "Enable prefork in Production")
)

func main() {
	flag.Parse()

	app := fiber.New(fiber.Config{
		Prefork: *prod,
	})

	app.Use(recover.New())
	app.Use(logger.New())

	v1 := app.Group("/api/v1")

	v1.Get("/buckets", handlers.BucketList)
	v1.Post("/buckets", handlers.BucketCreate)
	v1.Get("/buckets/:id", handlers.BucketGet)
	v1.Delete("/buckets/:id", handlers.BucketDelete)

	v1.Get("/buckets/:bucketId/files", handlers.FileList)
	v1.Post("/buckets/:bucketId/files", handlers.FileCreate)
	v1.Get("/buckets/:bucketId/files/:id", handlers.FileGet)
	v1.Delete("/buckets/:bucketId/files/:id", handlers.FileDelete)

	log.Fatal(app.Listen(*port))
}
