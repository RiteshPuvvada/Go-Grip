package main

import (
	
	"github.com/gofiber/fiber"
)

func Routes(app *fiber.App){
	app.Get(GetLeads)
	app.Get(GetLead)
	app.Post(NewLead)
	app.Delete(DeleteLead)
}

func initDatabase(){

}

func main(){
	app := fiber.New()
	Routes(app)
	app.Listen(3000)

}