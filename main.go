package main

import (
	"fmt"
	"log"
	"todo-app/config"
	"todo-app/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/spf13/viper"
)

func main() {
	config.LoadConfig()
	config.ConnectDB()
	app := fiber.New()
	frontend := viper.GetString("FRONTEND_URL")
	app.Use(cors.New(cors.Config{

		AllowOrigins: fmt.Sprintf("http://localhost:3000, http://127.0.0.1:8000, http://localhost:5173/ ,%v", frontend),
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	routes.RouteIndex(app)

	port := viper.GetString("PORT")
	err := app.Listen(fmt.Sprintf(":%v", port))
	if err != nil {
		log.Println(err, "Server Error")
	}

}
