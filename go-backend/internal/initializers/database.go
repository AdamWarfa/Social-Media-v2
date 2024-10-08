package initializers

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	// Connect to the database
	var err error
	dsn := os.Getenv("DB_URL_PSQL_DOCKER")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
}
