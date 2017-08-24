package repo

import (
	"database/sql"

	"github.com/italolelis/epioca/service/pkg/domain/auction"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
)

var (
	//ErrInvalidStatusProvided is used when an invalid status is provided
	ErrInvalidStatusProvided = errors.New("you need to provide a valid status (running, scheduled or completed)")
)

const (
	auctionTableName = "auctions"
)

type auctionRepo struct {
	db *sqlx.DB
}

// NewAuction returns a implementation of Auction repository
func NewAuction(db *sqlx.DB) auction.Repository {
	return &auctionRepo{db}
}

// Find returns a slice of all auctions
func (r *auctionRepo) Find(status string) ([]*auction.Auction, error) {
	auctions := make([]*auction.Auction, 0)

	if len(status) != 0 {
		if !auction.IsValidStatus(status) {
			return nil, ErrInvalidStatusProvided
		}

		if err := r.db.Select(&auctions, "SELECT * FROM auctions WHERE status = $1", status); err != nil && err != sql.ErrNoRows {
			return auctions, errors.Wrap(err, "r.db.Select All auctions")
		}
	} else {
		if err := r.db.Select(&auctions, "SELECT * FROM auctions"); err != nil {
			return auctions, errors.Wrap(err, "r.db.Select All auctions")
		}
	}

	return auctions, nil
}

// FindByID returns one auction by id
func (r *auctionRepo) FindByID(id uuid.UUID) (*auction.Auction, error) {
	var auction auction.Auction

	if err := r.db.Get(&auction, "SELECT * FROM auctions WHERE id = $1", id); err != nil {
		return &auction, errors.Wrap(err, "Get One auction by id")
	}

	return &auction, nil
}

// Add upserts auction into storage
func (r *auctionRepo) Add(auction *auction.Auction) error {
	if _, err := r.db.NamedExec(`INSERT INTO auctions 
		VALUES (:id, :status, :week, :country, :dc, :ingredient, :duration, :start_date, :end_date, :qty, :threshold, :max_price)
			ON CONFLICT (id) DO
		UPDATE SET  (status, week, country, dc, ingredient, duration, start_date, end_date, qty, threshold, max_price) =  (:status, :week, :country, :dc, :ingredient, :duration, :start_date, :end_date, :qty, :threshold, :max_price)
	`,
		auction); err != nil {
		return errors.Wrap(err, "Insert one auction")
	}

	return nil
}

// Remove removes auction by id
func (r *auctionRepo) Remove(id uuid.UUID) error {
	if _, err := r.db.Exec("DELETE FROM $1 WHERE id = $2", auctionTableName, id); err != nil {
		return errors.Wrap(err, "Delete one auction")
	}

	return nil
}
