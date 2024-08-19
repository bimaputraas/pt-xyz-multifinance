package main

import (
	"log"
	"xyz-multifinance/internal/api/controller"
	"xyz-multifinance/internal/api/middleware"
	"xyz-multifinance/internal/api/routes"
	"xyz-multifinance/internal/config"
	"xyz-multifinance/internal/logic"
	"xyz-multifinance/internal/repository"
	"xyz-multifinance/pkg"
)

func main() {
	conf, err := config.Init()
	if err != nil {
		log.Fatal(err)
	}

	db, err := pkg.NewGorm(conf.MySQLURI)
	if err != nil {
		log.Fatal(err)
	}
	repo := repository.NewMYSQL(db)
	l, err := logic.New(repo, conf)
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
