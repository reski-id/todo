package utils

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	// err := godotenv.Load()
	// if err != nil {
	// 	fmt.Println("Error loading .env file")
	// 	os.Exit(1)
	// }

	// dbUser := os.Getenv("MYSQL_USER")
	// dbPassword := os.Getenv("MYSQL_PASSWORD")
	// dbHost := os.Getenv("MYSQL_HOST")
	// dbPort := os.Getenv("MYSQL_PORT")
	// dbName := os.Getenv("MYSQL_DBNAME")

	dbUser := "root"
	dbPassword := ""
	dbHost := "localhost"
	dbPort := "3306"
	dbName := "todoappdb"

	dsn := dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"

	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
