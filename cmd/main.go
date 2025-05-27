package main

import (
	"log"
	"musgit/internal/adapters/db"
	"musgit/internal/application/api"
	"musgit/internal/application/domain"
)

func main() {

	dbAdapter, dbErr := db.NewAdapter("musgit.db")
	if dbErr != nil {
		log.Fatalf("Failed to init db, err: %v", dbErr)
	}

	app := api.NewMusgitService(dbAdapter)
	newPiece := domain.NewPiece("Invention #1", domain.Easy)
	_, err := app.StartPractice(*newPiece)
	if err != nil {
		log.Fatalf("Failed to start practice, error: %v", err)
	}

}
