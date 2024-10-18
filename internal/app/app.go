package app

import (
	"github.com/compashka/DatabasesCoursework/internal/config"
	"github.com/compashka/DatabasesCoursework/internal/database"
	"github.com/compashka/DatabasesCoursework/internal/rest"
	"log"
)

func Run() error {
	cfg := config.NewConfig()
	db, err := database.InitDB(cfg.Postgres)
	if err != nil {
		log.Fatal("Failed to init database", "err", err)
	}
	router := rest.SetupRouter(db)
	err = router.Run()
	if err != nil {
		log.Fatal("Failed to run router", "err", err)
	}
	return nil
}
