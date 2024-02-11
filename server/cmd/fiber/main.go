package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	api "github.com/panupong-develop/go-simple/api"
	v1 "github.com/panupong-develop/go-simple/api/v1"
	"github.com/panupong-develop/go-simple/configs"
	"github.com/panupong-develop/go-simple/pkg"
	"github.com/panupong-develop/go-simple/pkg/database"
)

type fiberApp struct {
}

func main() {
	// Load Configurations
	err := configs.LoadConfig()
	if err != nil {
		log.Fatalf(err.Error())
	}

	// Get Configurations
	config := configs.GetConfig()
	fmt.Printf("app name: %q \n", config.App.Name)

	// Connect Database
	db, err := database.ConnectSQLDatabase(config.Database)
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer db.Close()
	fmt.Println("connected to db:", db.DriverName())

	// Create Fiber app to server a web server
	app := fiber.New(fiber.Config{
		AppName:      config.App.Name,
		BodyLimit:    config.App.BodyLimitBytes,
		ReadTimeout:  config.App.ReadTimeout,
		WriteTimeout: config.App.WriteTimeout,
		JSONEncoder:  json.Marshal,
		JSONDecoder:  json.Unmarshal,
	})
	// API Routers

	// http://baseurl:port/
	app.Get(api.Index, func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// http://baseurl:port/v1/...
	apiGroup := app.Group(api.Version1)
	apiGroup.Get(api.Ping, v1.PingHandler)

	// Handle gracefully shutdown
	// Register channel to immediately handle signals to perform shutting down
	c := make(chan os.Signal, 1)
	pkg.HandleShutdownGracefully(c, app.Shutdown)

	// Server starts listening
	log.Println("server is starting up...")
	app.Listen(config.App.GetWebAppAddress())
}

func RunMainFiber() {
	main()
}
