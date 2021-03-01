package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/ridwanakf/autobase-twitter/internal/app"
	"github.com/ridwanakf/autobase-twitter/internal/delivery/worker"
)

func main() {
	// load .env
	err := godotenv.Load()
	if err != nil {
		log.Println("error loading .env file, will read directly from ENV vars instead")
	}

	// init app
	AutobaseApp, err := app.NewAutobaseApp()
	if err != nil {
		log.Fatalf("error initiating app %+v", err)
	}
	defer func() {
		errs := AutobaseApp.Close()
		for e := range errs {
			log.Println(e)
		}
	}()

	worker.Start(AutobaseApp)
}
