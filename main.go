package main

import (
	"golangpet/internal/app"
	"log"
)

func main() {
	application := &app.App{}
	if err := application.Run(); err != nil {
		log.Fatal(err)
	}
}
