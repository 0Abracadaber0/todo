package main

import (
	"github.com/gofiber/fiber/v2"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"todo/config"
	"todo/internal/db"
)

func main() {
	cfg := config.Load()

	log := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	log.Info("app is starting...", "cfg", cfg)

	dbConn, err := db.Connect(cfg.DatabasePath)
	if err != nil {
		panic(err)
	}
	defer dbConn.Close()

	app := fiber.New()

	go func() {
		if err := app.Listen(":8080"); err != nil {
			panic(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Info("Shutting down server...")
	app.Shutdown()
}
