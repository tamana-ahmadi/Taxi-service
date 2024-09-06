package db

import (
	"Taxi_service/configs"
	"fmt"
	"os"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var conn *gorm.DB

func ConnectDB() error {
	connstr := fmt.Sprintf(`host=%s port=%s user=%s  database=%s password=%s`,
		configs.AppSettings.PostgresParams.Host,
		configs.AppSettings.PostgresParams.Port,
		configs.AppSettings.PostgresParams.User,
		configs.AppSettings.PostgresParams.Database,
		os.Getenv("DB_PASSWORD"),
	)
	db, err := gorm.Open(postgres.Open(connstr), &gorm.Config{})
	if err != nil {
		return err
	}
	fmt.Println("Connected to db")
	conn = db
	return nil
}

func CloseDB(db *gorm.DB) error {
	// err := conn.
	// if err != nil {
	// 	return err
	// }
	return nil

}
func GetconnectDB() *gorm.DB {
	return conn
}
