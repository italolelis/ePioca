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
        	*
		FROM
            bids
        WHERE
            auction_id = $1
    `

	err := r.db.Select(&bids, query, auctionID)

	return bids, err
}

// FindByUser func
func (r *BidRepo) FindByUser(userID uuid.UUID) ([]*bid.Bid, error) {
	var bids []*bid.Bid

	query := `
        SELECT
			*
        FROM
            bids
        WHERE
            "user.id" = $1
    `

	err := r.db.Select(&bids, query, userID)

	return bids, err
}

// FindLatestByAuctionUser func
func (r *BidRepo) FindLatestByAuctionUser(auctionID uuid.UUID, userID uuid.UUID) (*bid.Bid, error) {
	var bid bid.Bid

	query := `
        SELECT *
        FROM
            bids
        WHERE
			auction_id = $1 AND "user.id" = $2
		ORDER BY 
			created DESC
		LIMIT 1	
    `

	err := r.db.Get(&bid, query, auctionID, userID)

	if err != nil {
		return &bid, errors.Wrap(err, "Get latest auction user bid")
	}

	return &bid, nil
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

	query := `
		INSERT INTO
			bids
			VALUES (:id, :auction_id, :user.id, :user.name, :threshold, :value, :created)
	`

	_, err := r.db.NamedExec(query, &bid)

	if err != nil {
		return errors.Wrap(err, "Insert a bid")
	}

	return nil
}
