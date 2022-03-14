// db, _ = gorm.Open(“postgres”, “user:pass@tcp(127.0.0.1:3306)/database?charset=utf8&parseTime=True&loc=Local”)
package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// func InitDB() *gorm.DB {
// 	// Load environment variables

// 	// init postgres database
// 	dsn := os.Getenv("POSTGRES_DSN")
// 	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	return db
// }

func InitDB() *gorm.DB {

	// dsn := os.Getenv("POSTGRES_DSN")
	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	dsn := "host=localhost user=gorm password=gorm dbname=gorm port=5432 sslmode=disable TimeZone=Asia/Kolkata"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
	}

	// dbSQL, err := db.DB()
	// if err != nil {
	// 	defer dbSQL.Close()
	// }
	// defer dbSQL.Close()
	return db
}
