package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"
)

func main(){
	err := godotenv.Load()

	if err!=nil{
		log.Println("Can't Load environment variable")
	}

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	engine := html.New("./views",".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/","./public")

	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error{

		return c.Render("index",nil)
	})

	app.Use(func(c *fiber.Ctx) error{
		return c.Status(404).JSON(fiber.Map{
			"Message":"404 Not Found",
		})
	})

	app.Listen(fmt.Sprintf("%v:%v",host,port))
}