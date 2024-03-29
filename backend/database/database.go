package database

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"gorm.io/driver/postgres"

	"github.com/Vanderscycle/3D-Printer-Exchange/config"
	"github.com/Vanderscycle/3D-Printer-Exchange/model"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Database instance
type Dbinstance struct {
	Db *gorm.DB
}

var DB Dbinstance

// Connect function
func Connect() {
	config.LocalEnvFile()
	p := config.ReadEnvVariableWithDefault("DATABASE_PORT")
	// because our config function returns a string, we are parsing our str to int here
	port, err := strconv.ParseUint(p, 10, 32)
	if err != nil {
		fmt.Println("Error parsing str to int")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
		config.ReadEnvVariableWithDefault("DATABASE_HOST"),
		config.ReadEnvVariableWithDefault("DATABASE_USER"),
		config.ReadEnvVariableWithDefault("DATABASE_PASSWORD"),
		config.ReadEnvVariableWithDefault("DATABASE_NAME"),
		port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}

	log.Println("Connected")
	db.Logger = logger.Default.LogMode(logger.Info)

	log.Println("running migrations")
	db.AutoMigrate(&model.User{}, &model.Printer{})

	DB = Dbinstance{
		Db: db,
	}
}
