package main

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/ramasapto/clean-architecture/adapter"
	"github.com/ramasapto/clean-architecture/config"
	"github.com/ramasapto/clean-architecture/controllers"
	"github.com/ramasapto/clean-architecture/repository"
	"github.com/ramasapto/clean-architecture/routes"
	"github.com/ramasapto/clean-architecture/usecases"
)

func init() {
	service := "clean-architecture"

	config.LoadConfig(service)
}

func main() {
	db := adapter.DBSQL()
	repo := repository.NewRepo(db)
	uc := usecases.NewUC(repo)
	ctrl := controllers.NewCtrl(uc)

	router := routes.NewRouter(ctrl)
	router.Router(os.Getenv("SERVER_PORT"))
}
