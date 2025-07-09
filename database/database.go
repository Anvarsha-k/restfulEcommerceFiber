package database

import (
	"log"
	"os"

	"github.com/Anvarsha-k/restfulEcommerceFiber/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func ConnectDb() {
	db, err := gorm.Open(sqlite.Open("api.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("failed to connect to db\n", err.Error())
		os.Exit(2)
	}
	log.Println("connected to the database successfully")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running Migrations")

	//TODO:	add migrations
	db.AutoMigrate(&models.User{},&models.Product{},&models.Orders{})
	Database = DbInstance{Db: db}

}
