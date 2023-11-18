package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/chrisbirster/webgamedev/pkg/database"
	"github.com/chrisbirster/webgamedev/pkg/logger"
	"github.com/chrisbirster/webgamedev/pkg/renderer"
	"github.com/chrisbirster/webgamedev/pkg/routes"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type environment string

const (
	EnvLocal       environment = "local"
	EnvDevelopment environment = "development"
	EnvProduction  environment = "production"
)

type Config struct {
	Environment environment `yaml:"environment"`
	HTTP        struct {
		Hostname string `yaml:"hostname"`
		Port     int    `yaml:"port"`
		TLS      struct {
			Enabled     bool   `yaml:"enabled"`
			Certificate string `yaml:"certificate"`
			Key         string `yaml:"key"`
		} `yaml:"tls"`
	} `yaml:"http"`
}

func main() {
	// move to using a config file
	c := Config{}
	c.Environment = EnvLocal
	c.HTTP.Hostname = "localhost"
	c.HTTP.Port = 42069
	web := InitWeb()

	// database
	db, err := database.InitDB("file:ligma.db")
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}

	// routes
	routes.HandleRoutes(web, db)
	// Start the server
	go func() {
		srv := http.Server{
			Addr:         fmt.Sprintf("%s:%d", c.HTTP.Hostname, c.HTTP.Port),
			Handler:      web,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
			IdleTimeout:  120 * time.Second,
		}

		if c.HTTP.TLS.Enabled {
			certs, err := tls.LoadX509KeyPair(c.HTTP.TLS.Certificate, c.HTTP.TLS.Key)
			if err != nil {
				web.Logger.Fatalf("cannot load TLS certificate: %v", err)
			}

			srv.TLSConfig = &tls.Config{
				Certificates: []tls.Certificate{certs},
			}
		}

		if err := web.StartServer(&srv); err != http.ErrServerClosed {
			web.Logger.Fatalf("shutting down the server: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	signal.Notify(quit, os.Kill)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := db.Close(); err != nil {
		web.Logger.Fatal(err)
	}

	if err := web.Shutdown(ctx); err != nil {
		web.Logger.Fatal(err)
	}
}

func InitWeb() *echo.Echo {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(logger.NewLoggerWithConfig())

	e.Renderer = renderer.NewTemplateRenderer()

	//

	// static files
	e.Static("/css", "css")
	e.Static("/static", "static")
	e.Static("/scripts", "scripts")
	return e
}
