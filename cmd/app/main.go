package main

import (
	"github.com/gofiber/fiber/v2"
	"log/slog"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
	"todo/config"
	"todo/internal/db"
	"todo/internal/middleware"
	"todo/internal/router"
	"todo/internal/services"
)

func main() {
	cfg := config.Load()

	log := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	log.Info("app is starting...", "cfg", cfg)

	err := db.Connect(cfg.DatabasePath)
	if err != nil {
		panic(err)
	}
	log.Info("successfully connected to database")

	err = db.RunMigrations(db.DB, cfg.MigrationsPath)
	if err != nil {
		panic(err)
	}

	app := fiber.New()

	app.Use(middleware.LoggerMiddleware(log))

	router.SetupRoutes(app)

	var wg sync.WaitGroup
	wg.Add(1)
	go services.OverdueChecker(&wg, time.Hour*24)

	go func() {
		if err := app.Listen(":8080"); err != nil {
			panic(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Info("shutting down server...")
	err = app.Shutdown()
	if err != nil {
		panic(err)
	}
}
