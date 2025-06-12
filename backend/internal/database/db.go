package database

import (
	"fmt"
	"hospital-portal/internal/database/model"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func InitDB() error{

	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: No .env file found.",err)
		return err
	}

	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	name := os.Getenv("DB_NAME")

	
	Db, err = gorm.Open(postgres.New(postgres.Config{
 		DSN: fmt.Sprintf("port=%s user=%s password=%s dbname=%s sslmode=disable",port,user,pass,name),
  		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})

	if err!=nil {
		return fmt.Errorf("failed to connect to db, ERROR: %v",err)
	}
	err = Db.AutoMigrate(&model.User{},&model.Patient{})
	if err != nil {
		return err
	}
	return nil
}

func GetDB() *gorm.DB{
	return Db
}

func CloseDB(){
		conn, err := Db.DB()
	if err != nil {
		fmt.Println(err)
		return
	}
	conn.Close()
}
