package main

import (
	"github.com/atadzan/AdvertAPI"
	"github.com/atadzan/AdvertAPI/pkg/handler"
	"github.com/atadzan/AdvertAPI/pkg/repository"
	"github.com/atadzan/AdvertAPI/pkg/service"
	"github.com/spf13/viper"
	_ "github.com/lib/pq"
	"log"
)

func main(){
	if err := initConfig(); err != nil{
		log.Fatalf("error initializing configs: %s", err.Error())
	}
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(AdvertAPI.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func initConfig() error{
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
