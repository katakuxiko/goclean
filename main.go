package main

import (
	"os"

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
	if err := srv.Run(port,handlers.InitRoutes()); err !=nil{
		logrus.Fatalf("error while runnig server:%s", err.Error())
	}
}

