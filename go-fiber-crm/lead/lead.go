package lead

import "github.com/jinzhu/gorm"

type Lead struct{
	gorm.Model
	Name      string
	Company   string
	Email     string
	Phone     int
}