package repo

import "github.com/jmoiron/sqlx"

const (
	tableName = "semaphore"

	SemaphoreMigrations = iota
)

// User represents an user repository
type Semaphore struct {
	db *sqlx.DB
}

func NewSemaphore(db *sqlx.DB) *Semaphore {
	return &Semaphore{db}
}

func (s *Semaphore) ensureSemaphoreTableExists() error {
	if _, err := s.db.Exec("CREATE TABLE IF NOT EXISTS " + tableName + " (id int not null primary key);"); err != nil {
		return err
	}
	return nil
}

func (s *Semaphore) ensureSemaphoreExists(id int) error {
	if err := s.ensureSemaphoreTableExists(); err != nil {
		return err
	}
	if _, err := s.db.Exec("INSERT INTO "+tableName+" (id) VALUES ($1) ON CONFLICT DO NOTHING", id); err != nil {
		return err
	}
	return nil
}

func (s *Semaphore) Lock(id int) error {
	if err := s.ensureSemaphoreExists(id); err != nil {
		return err
	}
	_, err := s.db.Exec("SELECT pg_advisory_lock(id) FROM "+tableName+" WHERE id = $1", id)
	return err
}

func (s *Semaphore) Unlock(id int) error {
	_, err := s.db.Exec("SELECT pg_advisory_unlock(id) FROM "+tableName+" WHERE id = $1", id)
	return err
}

func (s *Semaphore) TryLock(id int) (bool, error) {
	if err := s.ensureSemaphoreExists(id); err != nil {
		return false, err
	}
	var success bool
	err := s.db.QueryRow("SELECT pg_try_advisory_lock(id) FROM "+tableName+" WHERE id = $1", id).Scan(&success)
	return success, err
}
