package logger

import (
	"gopkg.in/natefinch/lumberjack.v2"
	"log/slog"
	"os"
	"path/filepath"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func SetupLogger(env string, logPath string) *slog.Logger {
	//var handler slog.Handler
	var log *slog.Logger
	switch env {
	case envLocal:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envDev, envProd:

		if err := os.MkdirAll(filepath.Dir(logPath), 0755); err != nil {
			panic("Не удалось создать директорию для логов: " + err.Error())
		}

		fileLogger := &lumberjack.Logger{
			Filename:   logPath,
			MaxSize:    10, // МБ
			MaxBackups: 5,
			MaxAge:     28, // Дней
			Compress:   true,
		}

		level := slog.LevelInfo
		if env == envDev {
			level = slog.LevelDebug
		}

		log = slog.New(slog.NewJSONHandler(fileLogger, &slog.HandlerOptions{Level: level}))
	}
	return log
}
