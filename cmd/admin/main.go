package main

import (
	"flag"
	"log"

	"github.com/cholaraja/revise/goweb/basicStructure/internal/platform/database"
	"github.com/cholaraja/revise/goweb/basicStructure/internal/schema"
)

func main() {
	flag.Parse()

	db, err := database.Open()
	if err != nil {
		log.Fatalf("error: Connecting to Database: %s", err)
	}
	defer db.Close()

	switch flag.Arg(0) {
	case "migrate":
		log.Print("Info: Migrating database")
		if err := schema.Migrate(db.DB); err != nil {
			log.Fatalf("error: Migrating database: %s", err)
		}
	case "seed":
		if err := schema.Seed(db.DB); err != nil {
			log.Fatalf("error: Seeding data: %s", err)
		}
	}

}
