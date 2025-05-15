package app

import (
	"backnedTestGolang/internal/config"
	"backnedTestGolang/internal/database/postgres"
	"backnedTestGolang/internal/delivery/http"
	"backnedTestGolang/internal/logger"
	"backnedTestGolang/internal/repository"
	"backnedTestGolang/internal/services"
	"github.com/joho/godotenv"
	"log"
	"log/slog"
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

	db, err := postgres.NewPostgresDB(cfg.StorageConfig)
	if err != nil {
		log.Error("Error initializing database: ", err)
	}
	log.Info("Database initialized successfully")

	repCart := repository.NewCartRepo(db)
	repOrder := repository.NewOrderRepo(db)

	cartService := services.NewCartService(repCart, repOrder)
	orderService := services.NewOrderService(repOrder)

	cartHandler := http.NewCartHandler(cartService)
	orderHandler := http.NewOrderHandler(orderService)

	router := http.NewRouter(cartHandler, orderHandler)

	server := http.NewServer(&cfg.HttpConfig, router.Init(log))

	sqlDB, err := db.DB()
	if err != nil {
		log.Error("Error connecting to DB: ", err)
	}
	defer sqlDB.Close()

	if err = server.Run(); err != nil {
		log.Error("failed to init server", slog.Attr{Key: "error", Value: slog.StringValue(err.Error())})
	}
}
