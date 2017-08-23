package repo

import (
	"ePioca/service/pkg/domain/bid"

	"github.com/jmoiron/sqlx"
	uuid "github.com/satori/go.uuid"
)

const bidTable = "bids"

//BidRepo struct
type BidRepo struct {
	db *sqlx.DB
}

//NewBidRepo func
func NewBidRepo(db *sqlx.DB) *BidRepo {
	return &BidRepo{db}
}

//FindByAuction func
func (r *BidRepo) FindByAuction(auctionID uuid.UUID) ([]*bid.Bid, error) {
	var bids []*bid.Bid

	query := `
        SELECT
            b.auction_id, b.user_id, b.threshold, b.value
        FROM
            bids b
        WHERE
            b.auction_id = $1
    `

	err := r.db.Select(&bids, query, auctionID)

	return bids, err
}

// FindByUser func
func FindByUser(userID uuid.UUID) ([]*bid.Bid, error) {
	return nil, nil
}

// FindLowest func
func FindLowest(auctionID uuid.UUID) (*bid.Bid, error) {
	return nil, nil
}

// Add func
func Add(bid *bid.Bid) error {
	return nil
}

// Remove func
func Remove(id uuid.UUID) error {
	return nil
}
