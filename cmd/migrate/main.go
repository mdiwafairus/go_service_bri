package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
)

func main() {
	action := flag.String("action", "up", "Migration action: up, down, drop")
	flag.Parse()

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s connect_timeout=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_TIMEOUT"),
	)

	m, err := migrate.New("file://database/migrations", dsn)
	if err != nil {
		log.Fatal("Migration init failed:", err)
	}

	switch *action {
	case "up":
		if err := m.Up(); err != nil && err.Error() != "no change" {
			log.Fatal(err)
		}
		fmt.Println("Migration up success")
	case "down":
		if err := m.Steps(-1); err != nil {
			log.Fatal(err)
		}
		fmt.Println("Migration down success")
	case "drop":
		if err := m.Drop(); err != nil {
			log.Fatal(err)
		}
		fmt.Println("All tables dropped")
	default:
		fmt.Println("Unknown action. Use -action=up|down|drop")
	}
}
