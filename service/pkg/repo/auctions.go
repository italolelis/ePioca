package repo

import (
	"github.com/italolelis/epioca/service/pkg/domain/auction"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
)

const (
	auctionTableName = "auctions"
)

type auctionRepo struct {
	db *sqlx.DB
}

// NewSemaphore returns a implementation of Auction repository
func NewAuction(db *sqlx.DB) auction.Repository {
	return &auctionRepo{db}
}

// FindAll returns a slice of all auctions
func (r *auctionRepo) FindAll() ([]*auction.Auction, error) {
	var auctions []*auction.Auction

	if err := r.db.Select(&auctions, "SELECT * FROM auctions"); err != nil {
		return auctions, errors.Wrap(err, "r.db.Select All auctions")
	}

	return auctions, nil
}

// Find returns one auction by id
func (r *auctionRepo) Find(id uuid.UUID) (*auction.Auction, error) {
	var auction *auction.Auction

	if err := r.db.Select(&auction, "SELECT * FROM $1 WHERE id = $2", auctionTableName, id); err != nil {
		return auction, errors.Wrap(err, "Get One auction by id")
	}

	return auction, nil
}

// Add upserts auction into storage
func (r *auctionRepo) Add(auction *auction.Auction) error {
	if _, err := r.db.NamedExec(`INSERT INTO auctions 
	VALUES (:id, :status, :week, :country, :dc, :ingredient, :duration, :start_date, :end_date, :qty, :threshold, :max_price)`,
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
