package schema

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id        uint `gorm:primaryKey`
	Name      string
	Age       int
	CreatedAt time.Time
	UpdatedAt time.Time
}

func Migrations(db *gorm.DB) {
	db.AutoMigrate(&User{})
}
