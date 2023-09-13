package main

import (
	"booksApi"
	"booksApi/pkg/handler"
	"booksApi/pkg/repository"
	"booksApi/pkg/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
	"os"
	"testing"
)

func TestMain_a2e85e6415(t *testing.T) {
	t.Log("Starting server...")
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		t.Errorf("error initializing database: %s", err.Error())
		return
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(booksApi.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		t.Errorf("error occurred while running http server: %s", err.Error())
		return
	}
	t.Log("Server running successfully")
}
