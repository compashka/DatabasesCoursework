package main

import (
	"github.com/youngpopeugene/DatabasesCoursework/internal/app"
	"log"
)

func main() {
	err := app.Run()
	log.Fatal("App is stopped", "err", err)
}
