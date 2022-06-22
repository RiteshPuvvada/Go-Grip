package lead

import (
	"github.com/RiteshPuvvada/Go-Grip/tree/main/go-fiber-crm/database"
	"github.com/gofiber/fiber"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/jinzhu/gorm"
)

type Lead struct{
	gorm.Model
	Name      string
	Company   string
	Email     string
	Phone     int
}