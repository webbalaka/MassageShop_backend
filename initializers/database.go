package initializers

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	// dbURL := "postgres://root:root@localhost:5432/test_db"
	var err error
	dsn := "host=localhost user=root password=root dbname=postgres port=5433 sslmode=disable TimeZone=Asia/Shanghai"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil{
		log.Fatal("Fail to connect to database")
	}
}