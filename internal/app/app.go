package app

import (
	"backnedTestGolang/internal/config"
	"backnedTestGolang/internal/logger"
	"github.com/joho/godotenv"
	"log"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}
func Run() {
	cfg, err := config.Init()
	if err != nil {
		log.Fatal(err)
	}
	log := logger.SetupLogger(cfg.Env, cfg.LogPath)
	log.Info("Logger started successfully")
}
