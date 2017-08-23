package repo

<<<<<<< HEAD
=======
import (
	"github.com/italolelis/epioca/service/pkg/domain/auction"
	"github.com/jmoiron/sqlx"
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
	return nil, nil
}

// Find returns one auction by id
func (r *auctionRepo) Find(id uuid.UUID) (*auction.Auction, error) {
	return nil, nil
}

// Add upserts auction into storage
func (r *auctionRepo) Add(auction *auction.Auction) error {
	return nil
}

// Remove removes auction by id
func (r *auctionRepo) Remove(id uuid.UUID) error {
	return nil
}
>>>>>>> 90faa0f71b9bc20ac2f43c8b625005bcf15122ed
