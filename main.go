package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"

	"github.com/katakuxiko/clean_go/package/handler"
	"github.com/katakuxiko/clean_go/package/repository"
	"github.com/katakuxiko/clean_go/package/service"
	"github.com/katakuxiko/clean_go/structure"
)
func init() {
    if err := godotenv.Load(); err != nil {
        logrus.Print("No .env file found")
    } else{
		logrus.Print("Config is OK")
	}
}
func main(){
	logrus.SetFormatter(new(logrus.JSONFormatter))
	port := os.Getenv("PORT")
	dbUrl:= os.Getenv("DATABASE_URL")
	db, err := repository.NewPostgresDB(dbUrl)
	if err != nil {
		logrus.Fatalf("Could not connect to database: %v", err)
	}
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(structure.Server)
	go func() {
		if err := srv.Run(port,handlers.InitRoutes()); err !=nil{
		logrus.Fatalf("error while runnig server:%s", err.Error())
	}
	}()
	logrus.Print("App started")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<- quit
	logrus.Print("Shutting down")
	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutdown %s", err)
	}
	if err := db.Close(); err != nil {
		logrus.Errorf("error occured on db connection close %s", err)

	}
}

