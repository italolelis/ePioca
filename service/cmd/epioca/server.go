package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/italolelis/epioca/service/pkg/config"
	"github.com/italolelis/epioca/service/pkg/handlers"
	"github.com/italolelis/epioca/service/pkg/migrations"
	"github.com/italolelis/epioca/service/pkg/repo"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	globalConfig *config.Specification
)

func init() {
	c, err := config.Load()
	if nil != err {
		log.WithError(err).Panic("Could not parse the environment configurations")
	}

	globalConfig = c
}

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	level, err := log.ParseLevel(globalConfig.LogLevel)
	FailOnError(err, "Could not parse the log level")

	log.SetLevel(level)

	log.SetFormatter(&log.JSONFormatter{FieldMap: log.FieldMap{
		log.FieldKeyTime: "@timestamp",
		log.FieldKeyMsg:  "message",
	}})
}

func RunServer(cmd *cobra.Command, args []string) {
	log.WithField("dsn", globalConfig.Database.DSN).Info("Trying to connect to DB")
	db, err := sqlx.Connect("postgres", globalConfig.Database.DSN)
	failOnError(err, "Failed to connect to DB")
	defer db.Close()

	log.WithField("path", globalConfig.Database.MigrationsPath).Info("Running migrations")
	semaphoreRepo := repo.NewSemaphore(db)
	migrationService := migrations.NewMigrationService(semaphoreRepo, globalConfig.Database.DSN, globalConfig.Database.MigrationsPath)
	err, _ = migrationService.Up()
	failOnError(err, "Could not run migrations")

	r := initRouter()
	initAuctionRoutes(r, db)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", globalConfig.Port), r))
}

func initRouter() chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to ePioca!!"))
	})

	return r
}

func initAuctionRoutes(r chi.Router, db *sqlx.DB) {

	handler := handlers.NewAuction()
	r.Route("/auctions", func(r chi.Router) {
		r.Get("/", handler.Index)
		r.Post("/", handler.Create)
		r.Get("/{id}", handler.Show)
		r.Put("/{id}", handler.Update)
		r.Delete("/{id}", handler.Remove)
	})
}

func failOnError(err error, msg string) {
	if err != nil {
		log.WithError(err).Panic(msg)
	}
}
