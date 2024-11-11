package schema

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id        uint
	Nama      string
	Usia      int
	CreatedAt time.Time
	UpdatedAt time.Time
}

func Schema(db *gorm.DB) {
	db.AutoMigrate(User{})
}
