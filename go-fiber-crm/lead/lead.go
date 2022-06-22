package lead

import (
	"github.com/RiteshPuvvada/Go-Grip/tree/main/go-fiber-crm/database"
	"github.com/gofiber/fiber"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/jinzhu/gorm"
)

type Lead struct{
	gorm.Model
	Name      string     `json:"name"`
	Company   string     `json:"company"`
	Email     string     `json:"email"`
	Phone     int        `json:"phone"`
}

func GetLeads(c *fiber.Ctx){
	db := database.DBConnection
	var leads []Lead
	db.Find(&leads)
	c.JSON(leads)
}

func GetLead(c *fiber.Ctx){
	id := c.Params("id")
	db := database.DBConnection
	var lead Lead
	db.Find(&lead, id)
	c.JSON(lead)

}

func NewLead(c *fiber.Ctx){
	db := database.DBConnection
	lead := new(Lead)
	if err:=c.BodyParser(lead); err != nil {
		c.Status(503).Send(err)
		return
	}
	db.Create(&lead)
	c.JSON(lead)
}

func DeleteLead(c *fiber.Ctx){
	id := c.Params("id")
	db := database.DBConnection
	var lead Lead
	db.First(&lead,id)
	if lead.Name == ""{
		c.Status(500).Send("No lead found with ID")
	}
	db.Delete(&lead)
	c.Send("Lead Successfully Deleted")
}