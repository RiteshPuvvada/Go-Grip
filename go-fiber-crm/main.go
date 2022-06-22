package main

import (
	"fmt"
	"github.com/RiteshPuvvada/Go-Grip/tree/main/go-fiber-crm/lead"
	"github.com/RiteshPuvvada/Go-Grip/tree/main/go-fiber-crm/database"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

func Routes(app *fiber.App){
	app.Get("/api/v1/lead",lead.GetLeads)
	app.Get("/api/v1/lead/:id",lead.GetLead)
	app.Post("/api/v1/lead",lead.NewLead)
	app.Delete("/api/v1/lead/:id",lead.DeleteLead)
}

func initDatabase(){
	var err error
	database.DBConnection, err = gorm.Open("sqlite3","leads.db")
	if err != nil {
		panic("Failed to connect databse")
	} 
	fmt.Println("Connection opened to database")
	database.DBConnection.AutoMigrate(&lead.Lead{})
	fmt.Println("Database Migrated")
}

func main(){
	app := fiber.New()
	initDatabase()
	Routes(app)
	app.Listen(3000)
	defer database.DBConnection.Close()
}