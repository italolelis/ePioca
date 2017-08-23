package migrations

import (
	"github.com/italolelis/epioca/service/pkg/repo"
	"github.com/jmoiron/sqlx"
	"github.com/mattes/migrate"
	_ "github.com/mattes/migrate/database/postgres"
	_ "github.com/mattes/migrate/source/file"
)

type MigrationService struct {
	semaphore      *repo.Semaphore
	dbDSN          string
	migrationsPath string
}

func NewMigrationService(semaphore *repo.Semaphore, dbDSN, migrationsPath string) *MigrationService {
	return &MigrationService{semaphore, dbDSN, migrationsPath}
}

// Up applies all migrations one by one
func (ms *MigrationService) Up() (error, bool) {
	if err := ms.semaphore.Lock(repo.SemaphoreMigrations); err != nil {
		return err, false
	}

	m, err := migrate.New(ms.migrationsPath, ms.dbDSN)
	if err != nil {
		return err, false
	}

	err = m.Up()
	if err != nil {
		return err, false
	}

	// unlock result must not affect migration run result, but may cause more errors
	if err := ms.semaphore.Unlock(repo.SemaphoreMigrations); err != nil {
		return err, false
	}

	return nil, true
}

// ResetDB removes all data from DB
func (ms *MigrationService) ResetDB(db *sqlx.DB) error {
	tables := []string{
		"schema_migrations",
		"auctions",
		"bids",
	}
	for i := range tables {
		_, err := db.Exec("DROP TABLE IF EXISTS " + tables[i])
		if nil != err {
			return err
		}
	}
	return nil
}
