package main

import (
	"golangpet/internal/app"
	"log"
)

func main() {
	application := new(app.App)
	if err := application.Run(); err != nil {
		log.Fatal(err)
	}
}
