package main

import (
	hickmet "hickmet_whatapp_bot"
	"hickmet_whatapp_bot/internal/handlers"
	"hickmet_whatapp_bot/internal/repository"
	"hickmet_whatapp_bot/internal/service"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func main() {
	if err := godotenv.Load("./.env"); err != nil {
		log.Fatalf("some error with env initializiing: %s", err.Error())
	}

	if err := initConfig(); err != nil {
		log.Fatalf("some error with initializiing: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})

	if err != nil {
		log.Fatalf("error with data base: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	service := service.NewService(repos)
	handler := handlers.NewHandler(service)

	srv := new(hickmet.Server)
	if err := srv.Run(viper.GetString("port"), handler.InitRoutes()); err != nil {
		log.Fatalf("error with run server: %s", err.Error())
	}

}

func initConfig() error {
	viper.AddConfigPath("./config")
	viper.SetConfigName("db")
	return viper.ReadInConfig()
}
