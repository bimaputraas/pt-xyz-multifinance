package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"xyz-multifinance/internal/api/controller"
	"xyz-multifinance/internal/api/middleware"
	"xyz-multifinance/internal/api/routes"
	"xyz-multifinance/internal/logic"
	"xyz-multifinance/internal/repository"
	"xyz-multifinance/pkg"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := pkg.NewGorm(os.Getenv("MYSQL_URI"))
	if err != nil {
		log.Fatal(err)
	}

	repo := repository.NewMYSQL(db)
	l, err := logic.New(repo)
	if err != nil {
		log.Fatal(err)
	}
	ctr, err := controller.New(l)
	if err != nil {
		log.Fatal(err)
	}

	midware := middleware.New(l)

	r := routes.New(midware, ctr)

	r.Run(":8080")
}
