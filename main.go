package main

import (
	"github.com/Golang_Portfolio_Website/model"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main(){
	// Loading environment variable
	err := godotenv.Load()
	if err!=nil{
		log.Fatalln("Can't Load environment variable")
	}
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")

	// establish go fiber condiguration 
	engine := html.New("./views",".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// set up public/static content serving thru public folder
	app.Static("/public","./public")
	
	// set up logger to log all traffic
	app.Use(logger.New())

	// opening portfolio.json file
	file,err := os.Open("portfolio.json")
	if err!=nil{
		log.Fatalln("Can't open file 'portfolio.json'")
	}
	defer file.Close() // close when its no longer needed

	// turning json file to array of struct
	var portfolio []model.Portfolio
	if err = json.NewDecoder(file).Decode(&portfolio); err != nil{
		log.Fatalln("Can't turn json file to array of struct'")
	}

	// Serve and render index.html
	app.Get("/", func(c *fiber.Ctx) error{
		return c.Render("index",fiber.Map{
			"portfolio":portfolio,
		})
	})
	// serve 404 Not Found page
	app.Use(func(c *fiber.Ctx) error{
		return c.Status(404).JSON(fiber.Map{
			"Message":"404 Not Found",
		})
	})

	app.Listen(fmt.Sprintf("%v:%v",host,port))
}