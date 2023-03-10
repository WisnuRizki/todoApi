package database

import (
	"fmt"
	"os"

	"todoapi.wisnu.net/modules"

	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func createActivityTable(tx *gorm.DB) error {
	err := tx.AutoMigrate(&modules.Activity{})
	if err != nil {
		return err
	}
	return nil
}


func createTodoTable(tx *gorm.DB) error {
	err := tx.AutoMigrate(&modules.Todo{})
	if err != nil {
		return err
	}
	return nil
}


func ConnectDatabase() {

	err := godotenv.Load()
    if err != nil {
        fmt.Println("Error loading .env file")
        os.Exit(1)
    }
	
	dbHost := os.Getenv("MYSQL_HOST")
	dbPort := os.Getenv("MYSQL_PORT")
	if dbPort == "" {
        dbPort = "3036"
    }
	dbUser := os.Getenv("MYSQL_USER")
	dbPassword := os.Getenv("MYSQL_PASSWORD")
	dbName := os.Getenv("MYSQL_DBNAME")

	dbURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)

	db, err := gorm.Open(mysql.Open(dbURL), &gorm.Config{})
	if err != nil {
		fmt.Println("Error opening connection to database:", err)
		return
	}
	
	if err != nil {
		panic(err)
	}

	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: "activity-migrations",
			Migrate: func(tx *gorm.DB) error {
				if err := createActivityTable(tx); err != nil {
					return err
				}
				if err := createTodoTable(tx); err != nil {
					return err
				}
				return nil
			},
			
		},
	})

	if err := m.Migrate(); err != nil {
		panic(err)
	}

	DB = db
}