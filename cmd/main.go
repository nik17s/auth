package main

import (
	"context"
	"github.com/joho/godotenv"
	entity "github.com/nik17s/auth"
	handler2 "github.com/nik17s/auth/pkg/handler"
	"github.com/nik17s/auth/pkg/repository/postgres"
	service2 "github.com/nik17s/auth/pkg/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"os"
	"os/signal"
)

func main() {
	logger := getLogger()

	if err := initConfig(); err != nil {
		logger.Fatalf("error intializing config: %s", err.Error())
	}
	logger.Infof("config initialized successfully")

	if err := godotenv.Load(); err != nil {
		logger.Fatalf("error loading env variables: %s", err.Error())
	}
	logger.Infof("env variables successfully loaded")

	postgresCfg := postgres.Config{
		Username:     viper.GetString("db.username"),
		Password:     os.Getenv("DB_PASSWORD"),
		Host:         viper.GetString("db.host"),
		Port:         viper.GetString("db.port"),
		DBName:       viper.GetString("db.name"),
		PoolMaxConns: 10,
	}

	postgresPool, err := postgres.NewConnectionPool(context.Background(), postgresCfg)
	if err != nil {
		logger.Fatalf("database connection error: %s", err.Error())
	}
	logger.Infof("database connection successfully made")

	repository := postgres.NewRepository(postgresPool)
	service := service2.NewService(repository)
	handler := handler2.NewHandler(service)

	server := new(entity.Server)

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		if err := server.Shutdown(context.Background()); err != nil {
			logrus.Infof("server shutdown error: %s", err.Error())
		}
	}()

	if err := server.Run(viper.GetString("port"), handler.GetRouter()); err != http.ErrServerClosed {
		log.Fatalf("server running error: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func getLogger() *logrus.Logger {
	logger := logrus.New()
	logger.SetFormatter(new(logrus.JSONFormatter))
	logger.SetReportCaller(true)

	return logger
}
