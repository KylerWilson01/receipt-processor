package main

import (
	"flag"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"github.com/KylerWilson01/receipt-processor.git/routes"
)

var prod = flag.Bool("prod", false, "Enable prefork in Production")

func main() {
	port := flag.String("port", "3000", "port to listen on")
	flag.Parse()

	// create app
	app := fiber.New(fiber.Config{
		Prefork: *prod,
	})

	// middleware
	app.Use(recover.New())
	app.Use(logger.New())

	// endpoints
	entrance := app.Group("/")
	routes.ReceiptRoute(entrance)

	// listen
	log.Fatal(app.Listen(":" + *port))
}
