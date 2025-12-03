package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"
)

const migrationsDir = "./internal/infra/db/migrations"

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: migrate [new|up] [name]")
	}

	switch os.Args[1] {
	case "new":
		if len(os.Args) < 3 {
			log.Fatal("Usage: migrate new <migration_name>")
		}
		name := os.Args[2]
		createMigration(name)

	default:
		log.Fatalf("Unknown command: %s", os.Args[1])
	}
}

func createMigration(name string) {
	if _, err := os.Stat(migrationsDir); os.IsNotExist(err) {
		err := os.Mkdir(migrationsDir, 0755)
		if err != nil {
			log.Fatalf("Cannot create migrations directory: %v", err)
		}
	}

	timestamp := time.Now().Format("20060102_150405")

	filename := fmt.Sprintf("%s_%s.sql", timestamp, name)
	path := filepath.Join(migrationsDir, filename)

	content := fmt.Sprintf("-- Migration: %s\n-- Created at: %s\n\n",
		name,
		time.Now().Format(time.RFC3339),
	)

	err := ioutil.WriteFile(path, []byte(content), 0644)
	if err != nil {
		log.Fatalf("Cannot write migration: %v", err)
	}

	fmt.Println("Created migration:", path)
}
