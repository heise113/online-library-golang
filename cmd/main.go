package main

import (
	_ "github.com/lib/pq"
	"online_lib_api"
	"online_lib_api/internal/handler"
	"online_lib_api/internal/service"
	"online_lib_api/internal/storage"
	"log"
)

func main() {
	db, err := storage.NewPostgresDB(storage.Config{
		Host: "localhost",
		Port: "5436",
		Username: "postgres",
		Password: "qwerty",
		DBName: "postgres",
		SSLMode: "disable",
	})
	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}

	storage := storage.NewStorage(db)
	services := service.NewService(storage)
	handlers := handler.NewHandler(services)

	srv := new(online_lib_api.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf(err.Error())
	}
}