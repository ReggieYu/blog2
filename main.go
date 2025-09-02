package main

import (
	"blog/config"
	"blog/database"
	"blog/router"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load() // Load .env file if it exists
	cfg := config.Load()

	if err := database.Init(cfg); err != nil {
		log.Fatalf("failed to initialize database: %v", err)
	}

	r := router.Setup(cfg)
	addr := ":" + cfg.Port
	log.Printf("server listening on %s mode = %s db = %s\n", addr, os.Getenv("GIN_MODE"), cfg.DBDriver)
	if err := r.Run(addr); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
