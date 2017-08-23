package main

import (
	"os"

	"github.com/italolelis/epioca/service/pkg/migrations"
	"github.com/italolelis/epioca/service/pkg/repo"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.WarnLevel)

	log.SetFormatter(&log.JSONFormatter{FieldMap: log.FieldMap{
		log.FieldKeyTime: "@timestamp",
		log.FieldKeyMsg:  "message",
	}})
}

func RunServer(cmd *cobra.Command, args []string) {

	log.WithField("path", globalConfig.MigrationsPath).Info("Running migrations")
	semaphoreRepo := repo.NewSemaphore(db)
	migrationService := migrations.NewMigrationService(semaphoreRepo, globalConfig.DatabaseDSN, globalConfig.MigrationsPath)
	errList, ok := migrationService.Up()

}
