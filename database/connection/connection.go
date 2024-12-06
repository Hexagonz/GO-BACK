package database

import (
	"fmt"
	"log"

	"github.com/Hexagonz/back-end-go/database"
	"github.com/Hexagonz/back-end-go/env"
)

func init() {
	Conncetion()
}

func Conncetion() (interface{}) {
	db, err := database.SetupDatabase()
	db_name := env.DotEnvVariable("DB_NAME")
	if err != nil {
		log.Fatalf("failed to set up database: %v", err)
	}
	_ = db
	fmt.Printf("\nDatabase %s setup is complete!\n", db_name)
	return nil
}
