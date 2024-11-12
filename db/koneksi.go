package db

import (
	"belajar/schema"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Koneksi() {
	dsn := "host=localhost user=postgres password=123456 dbname=belajar port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	DB, _ = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	schema.Migrations(DB)
}
