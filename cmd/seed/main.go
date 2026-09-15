package main

import (
	"flag"
	"fmt"
	"log"

	"server/seeder"

	"github.com/joho/godotenv"
)

func main() {
	force := flag.Bool("force", false, "reset even if the database contains journals or registered users")
	skipUsers := flag.Bool("skip-users", false, "don't create the demo users from users.json (use in production)")
	flag.Parse()

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	hasUserData, err := seeder.HasUserData()
	if err != nil {
		log.Fatal(err)
	}
	if hasUserData && !*force {
		log.Fatal("refusing to seed: the database contains journals or registered users, which seeding deletes. Run with -force to wipe them anyway.")
	}

	if err := seeder.ResetAndSeedDatabase(seeder.Options{SkipUsers: *skipUsers}); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Seeding complete.")
}
