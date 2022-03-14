// db, _ = gorm.Open(“postgres”, “user:pass@tcp(127.0.0.1:3306)/database?charset=utf8&parseTime=True&loc=Local”)
package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
// "postgres", 5432, "user", "mypassword", "user")
func InitDB() *gorm.DB {
	dbURL := "postgres://user:mypassword@postgres/user"

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	// db.AutoMigrate(&models.Book{})

	return db
}

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

// func InitDB() *gorm.DB {

// 	// dsn := os.Getenv("POSTGRES_DSN")
// 	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	dsn := "host=localhost user=gorm password=gorm dbname=gorm port=5432 sslmode=disable TimeZone=Asia/Kolkata"
// 	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	// dbSQL, err := db.DB()
// 	// if err != nil {
// 	// 	defer dbSQL.Close()
// 	// }
// 	// defer dbSQL.Close()
// 	return db
// }
