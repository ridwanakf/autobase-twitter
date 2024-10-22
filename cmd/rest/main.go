package main

import (
	"log"

	"github.com/ridwanakf/autobase-twitter/internal/app"
	"github.com/ridwanakf/autobase-twitter/internal/delivery/rest"
)

func main() {
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

	rest.Start(AutobaseApp)
}
