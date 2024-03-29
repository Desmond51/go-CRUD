package initializers

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
var DB *gorm.DB

func ConnectToDb(){
	var err error
	// postgres://gtkqcaeq:1OuxndbAt3s6mZtlq-SYUU2WPzbEqnDu@heffalump.db.elephantsql.com/gtkqcaeq
	dsn := os.Getenv("DB_URL")
DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

if err !=nil {
	log.Fatal("Failed to connect to database")
}
}