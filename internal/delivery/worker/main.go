package worker

import (
	"log"

	"github.com/ridwanakf/autobase-twitter/internal/app"
)

func Start(app *app.AutobaseApp) {
	for {
		log.Println("Test Worker")
	}
}