package repo

import (
	"github.com/italolelis/epioca/service/pkg/domain/bid"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
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
func (r *BidRepo) FindByUser(userID uuid.UUID) ([]*bid.Bid, error) {
	var bids []*bid.Bid

	query := `
        SELECT
            b.auction_id, b.user_id, b.threshold, b.value
        FROM
            bids b
        WHERE
            b.user_id = $1
    `

	err := r.db.Select(&bids, query, userID)

	return bids, err
}

// FindLowest func
func (r *BidRepo) FindLowest(auctionID uuid.UUID) (*bid.Bid, error) {
	var bid bid.Bid

	query := `
		SELECT 
			b.* 
		FROM bids b
		WHERE 
			b.auction_id = $1
		ORDER BY 
			b.value ASC
		LIMIT 1
    `

	err := r.db.Get(&bid, query, auctionID)

	if err != nil {
		return &bid, errors.Wrap(err, "Get lowest auction bid")
	}

	return &bid, nil
}

// Add func
func (r *BidRepo) Add(bid *bid.Bid) error {
	return nil
}

// Remove func
func (r *BidRepo) Remove(id uuid.UUID) error {
	return nil
}
