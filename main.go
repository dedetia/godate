package main

import (
	"github.com/dedetia/godate/config"
	"github.com/dedetia/godate/internal/adapter/registry"
	"github.com/dedetia/godate/internal/adapter/rest"
	"github.com/dedetia/godate/pkg/auth"
	mongostore "github.com/dedetia/golib/storage/mongo"
	"log"
)

func main() {
	log.Println("Starting server...")

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("error load config: %v", err)
	}

	db := mongostore.New(&mongostore.Options{
		URI:     cfg.Database.Mongo.URI,
		DB:      cfg.Database.Mongo.DB,
		AppName: cfg.Server.AppName,
	})
	log.Println("Mongo connected...")

	err = auth.Configure(cfg.PrivateKey)
	if err != nil {
		log.Fatalf("error configure auth: %v", err)
	}

	repositoryRegistry := registry.NewRepositoryRegistry(db)
	serviceRegistry := registry.NewServiceRegistry(cfg, repositoryRegistry)

	server := rest.New(cfg)
	server.RegisterHandler(serviceRegistry)
	server.Run()
}
