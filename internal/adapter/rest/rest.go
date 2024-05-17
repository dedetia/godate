package rest

import (
	"fmt"
	"github.com/dedetia/godate/config"
	"github.com/dedetia/godate/internal/adapter/rest/handler"
	"github.com/dedetia/godate/internal/adapter/rest/middleware"
	"github.com/dedetia/godate/internal/core/port/registry"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Rest struct {
	config *config.MainConfig
	app    *fiber.App
}

func New(config *config.MainConfig) *Rest {
	app := fiber.New(fiber.Config{
		AppName:      config.Server.AppName,
		ReadTimeout:  time.Duration(config.Server.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(config.Server.WriteTimeout) * time.Second,
	})

	return &Rest{
		config: config,
		app:    app,
	}
}

func (r *Rest) RegisterHandler(serviceRegistry registry.ServiceRegistry) {
	h := handler.NewHandler(serviceRegistry)

	r.app.Use(cors.New())
	r.app.Use(recover.New(
		recover.Config{
			EnableStackTrace: true,
		}))
	r.app.Use(logger.New())

	r.app.Static("/photos", os.Getenv("PHOTO_DIR"))

	r.app.Post("/login", h.Login)
	r.app.Post("/signup", h.Signup)
	r.app.Get("/profiles", middleware.Auth, h.Profile)
	r.app.Post("/swipes", middleware.Auth, h.SwipeAction)
	r.app.Post("/purchase-premium", middleware.Auth, h.PurchasePremium)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		log.Println("Gracefully shutting down...")
		_ = r.app.Shutdown()
	}()
}

func (r *Rest) Run() {
	log.Println("Server is listening on port", r.config.Server.Port)
	err := r.app.Listen(fmt.Sprintf(":%d", r.config.Server.Port))
	if err != nil {
		log.Fatal(err)
	}
}
