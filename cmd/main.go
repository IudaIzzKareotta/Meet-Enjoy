package main

import (
	"context"
	MeetEnjoy2 "github.com/IudaIzzKareotta/Meet-Enjoy"
	"github.com/IudaIzzKareotta/Meet-Enjoy/pkg/handlers"
	repository2 "github.com/IudaIzzKareotta/Meet-Enjoy/pkg/repository"
	service2 "github.com/IudaIzzKareotta/Meet-Enjoy/pkg/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	if err := initConfig(); err != nil {
		logrus.Fatalf("Error initializing config: %s", err)
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("Error loading .env file: %s", err)
	}

	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetLevel(logrus.InfoLevel)

	db, err := repository2.NewPostgresDb(repository2.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		logrus.Fatalf("Error connecting to database: %s", err)
	}

	repos := repository2.NewRepository(db)
	services := service2.NewService(repos)
	handler := handlers.NewHandler(services)

	srv := new(MeetEnjoy2.Server)
	go func() {
		if err := srv.Run(viper.GetString("port"), handler.InitRoutes()); err != nil {
			logrus.Fatalf("Error starting server: %s", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	if err := srv.ShutDown(context.Background()); err != nil {
		logrus.Printf("error occured on server shutting down: %s", err.Error())
	}

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
