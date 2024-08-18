package main

import (
	"fmt"

	"github.com/dannyoka/go-mongo/src/config"
	"github.com/dannyoka/go-mongo/src/db"
	"github.com/dannyoka/go-mongo/src/repositories"
	"github.com/dannyoka/go-mongo/src/services"
)

type Person struct {
	Name  string
	Value float64
}

func main() {
	config := config.InitConfig()
	db := db.InitDB(config, "test")
	repositoriesContext := repositories.InitRepositories(db)
	services.InitServices(repositoriesContext)
	fmt.Println("Hello World")
	fmt.Println("Hello World again")
}
