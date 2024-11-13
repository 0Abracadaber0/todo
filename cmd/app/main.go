package main

import (
	"github.com/gofiber/fiber/v2"
	"log/slog"
	"os"
	"os/signal"
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

	dbConn, err := db.Connect(cfg.DatabasePath)
	if err != nil {
		panic(err)
	}
	log.Info("successfully connected to database")
	defer dbConn.Close()

	err = db.RunMigrations(dbConn, cfg.MigrationsPath)
	if err != nil {
		panic(err)
	}

	app := fiber.New()
	taskService := services.NewTaskService(dbConn)
	go taskService.OverdueChecker(time.Minute)

	app.Use(middleware.LoggerMiddleware(log))
	app.Use(middleware.TaskServiceMiddleware(taskService))

	router.SetupRoutes(app)

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
