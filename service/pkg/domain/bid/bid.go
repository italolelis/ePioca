package bid

import (
	uuid "github.com/satori/go.uuid"
)

// Repository defines the basic methods for a repository
type Repository interface {
	FindByAuction(auctionID uuid.UUID) ([]*Bid, error)
	FindByUser(userID uuid.UUID) ([]*Bid, error)
	FindLowest(auctionID uuid.UUID) (*Bid, error)
	Add(bid *Bid) error
	Remove(id uuid.UUID) error
}

// Bid represents a supplier bid to an auction
type Bid struct {
	ID        uuid.UUID `json:"id" db:"id"`
	AuctionID uuid.UUID `json:"_" db:"auction_id"`
	UserID    uuid.UUID `json:"user_id" db:"user_id"`
	Threshold int       `json:"threshold" db:"threshold"`
	Value     float32   `json:"value" db:"value"`
}
