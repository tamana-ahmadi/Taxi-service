package main

import (
	"Taxi_service/configs"
	"Taxi_service/db"
	"Taxi_service/logger"
	"Taxi_service/package/controllers"
	"Taxi_service/server"
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
)

// @title  Fellow traveler API
// @version 1.0
// @description API Server for Fellow traveler Application

// @host localhost:8484
// @BasePath /

// @securityDefinitions.apikey AKA
// @in header
// @name Authorization

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(errors.New(fmt.Sprintf("error loading .env file. Error is %s", err)))
	}
	err = configs.ReadString()
	if err != nil {
		panic(err)
	}
	logger.Init()
	err = db.ConnectDB()
	if err != nil {
		panic(err)
	}
	err = db.Migrate()
	if err != nil {
		panic(err)
	}

	MainServer := new(server.Server)

	go func() {
		err = MainServer.Run(configs.AppSettings.AppParams.PortRun, controllers.InitRoutes())
		if err != nil {
			log.Fatal()
		}

	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	fmt.Printf("\nНачало завершения программ\n")

	sqlDB, err := db.GetconnectDB().DB()
	if err != nil {
		err = sqlDB.Close()
		if err != nil {
			log.Fatalf("Ошибка при закрытии соединения с БД: %s", err)
		}
	} else {
		log.Fatalf("Ошибка при получении *sql.DB из GORM: %s", err)
	}
	fmt.Println("Соединение с БД успешно закрыто")

	err = MainServer.Shutdown(context.Background())
	if err != nil {
		log.Fatalf("Ошибка при завершении работы сервера: %s", err)
	}
	fmt.Println("HTTP-сервис успешно выключен")
	fmt.Println("Конец завершения программы")

}
