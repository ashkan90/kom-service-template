package main

import (
	"context"
	"github.com/labstack/echo/v4"
	"kom-service-template/adapters/book/delivery/http/middleware"
	prometheus_http "kom-service-template/adapters/prometheus/handler/http"
	"log"
	"os"

	"github.com/golang-common-packages/storage"
	"kom-service-template/config"
)

var (
	cfg    config.Config
	dbConn storage.INoSQLDocument
	ctx    = context.Background()
)

func init() {

	cfg = config.New()
	if cfg.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}

	dbConn = storage.New(ctx, storage.NOSQLDOCUMENT)(storage.MONGODB, &storage.Config{MongoDB: storage.MongoDB{
		User:     cfg.GetString("database.mongodb.user"),
		Password: cfg.GetString("database.mongodb.password"),
		Hosts:    cfg.GetStringSlice("database.mongodb.hosts"),
		Options:  cfg.GetStringSlice("database.mongodb.options"),
		DB:       cfg.GetString("database.mongodb.dbName"),
	}}).(storage.INoSQLDocument)
}

func main() {

	e := echo.New()
	middL := BookHttpMiddleware.New()
	e.Use(middL.CORS)

	//bookRepo := bookMongoRepository.New(dbConn)
	//bookUCase := bookUsecase.New(bookRepo, cfg.GetString("database.mongodb.dbName"), cfg.GetString("database.mongodb.collections.book"))
	//bookHttpDelivery.New(e, bookUCase)


	prometheus_http.New(e)

	if cfg.GetBool("debug") {
		e.Start(":" + cfg.GetString("server.port"))
	} else {
		e.Start(":" + os.Getenv("PORT"))
	}
}
