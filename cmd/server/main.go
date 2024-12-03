package main

import (
	"log"
)
func main() {
	db , err := database.NewDBconnection()
	if err != nil  {
		log.Fatalf("Database connection failed: %v", err)
	}

	if err := database.RunMigrations(db); err != nil {
		log.Fatalf("Database migration failed: %v",err)
	}
}