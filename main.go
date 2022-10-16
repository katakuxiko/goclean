package main

import (
	"log"
	"os"
	"github.com/joho/godotenv"

	"github.com/katakuxiko/clean_go/package/handler"
	"github.com/katakuxiko/clean_go/package/repository"
	"github.com/katakuxiko/clean_go/package/service"
	"github.com/katakuxiko/clean_go/structure"
)
func init() {
    if err := godotenv.Load(); err != nil {
        log.Print("No .env file found")
    } else{
		log.Print("Config is OK")
	}
}
func main(){
	port := os.Getenv("PORT")
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	
	srv := new(structure.Server)
	if err := srv.Run(port,handlers.InitRoutes()); err !=nil{
		log.Fatalf("error while runnig server:%s", err.Error())
	}
}

