package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/iandanarko/concert/config"
	"github.com/iandanarko/concert/internal/router"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	cfg := buildConfig()
	db := buildDB(cfg)
	router.BuildRoutes(e, cfg, db)
	runServer(e, cfg.Port)
	waitShutDown(e)
}

func buildConfig() config.Config {
	cfg, err := config.New(".env")
	if err != nil {
		log.Fatalf("Error load config: %v", err)
	}
	return cfg
}

func runServer(e *echo.Echo, port string) {
	go func() {
		if err := e.Start(fmt.Sprintf(":%s", port)); err != nil {
			log.Fatalf("Error Start Server: %v", err)
		}
	}()
}

func waitShutDown(e *echo.Echo) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		log.Fatalf("Error Shutdown: %v", err)
	}
}

func buildDB(cfg config.Config) *sql.DB {
	sqlCfg := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		cfg.DB.Username,
		cfg.DB.Password,
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.Name,
	)
	log.Println(sqlCfg)

	db, err := sql.Open(config.DbDriver, sqlCfg)
	if err != nil {
		log.Fatalf("Error connect db: %v", err)
	}

	return db
}
