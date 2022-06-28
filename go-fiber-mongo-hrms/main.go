package main

import (
	"github.com/gofiber/fiber/v2"
)

type MongoInstance struct{
	Client
	Db
}

var MI MongoInstance

const DbName = "fiber-hrms"
const mongoURI = "mongodb://localhost:27017" + DbName

type Employee struct{
	ID        string 
	Name      string 
	Salary    float64
	Age       float64
}

func Connect() error{

}

func main(){
	app := fiber.New()

	app.Get("/employee", func(c *fiber.Ctx) error {
		
	})
	app.Post("/employee")
	app.Put("/employee/:id")
	app.Delete("/employee/:id")
}