package main

import (
	"fmt"
	"log"

	"server/seeder"
)

func main() {
	// if err := godotenv.Load(); err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	// if err := seeder.RequireTestEnvironment(); err != nil {
	// 	log.Fatal(err)
	// }

	if err := seeder.ResetAndSeedDatabase(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Seeding complete.")
}
