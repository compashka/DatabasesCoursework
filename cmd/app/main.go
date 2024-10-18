package main

import (
	"github.com/compashka/DatabasesCoursework/internal/app"
	"log"
)

func main() {
	err := app.Run()
	log.Fatal("App is stopped", "err", err)
}
