package bid

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// Repository defines the basic methods for a repository
type Repository interface {
	FindByAuction(auctionID uuid.UUID) ([]*Bid, error)
	FindByUser(userID uuid.UUID) ([]*Bid, error)
	FindLowest(auctionID uuid.UUID) (*Bid, error)
	FindLatestByAuctionUser(auctionID uuid.UUID, userID uuid.UUID) (*Bid, error)
	Add(bid *Bid) error
}

type User struct {
	ID   uuid.UUID `json:"id" db:"id"`
	Name string    `json:"name" db:"name"`
}

// Bid represents a supplier bid to an auction
type Bid struct {
	ID        uuid.UUID `json:"id" db:"id"`
	AuctionID uuid.UUID `json:"-" db:"auction_id"`
	User      *User     `json:"user" db:"user"`
	Threshold int       `json:"threshold" db:"threshold"`
	Value     float32   `json:"value" db:"value"`
	Created   time.Time `json:"created" db:"created"`
}

// NewBid creates a new instance of Bid
func NewBid(id uuid.UUID) *Bid {
	return &Bid{ID: id}
}

type Bids []*Bid

func (slice Bids) Len() int {
	return len(slice)
}

func (slice Bids) Less(i, j int) bool {
	return slice[i].Value < slice[j].Value
}

func (slice Bids) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}
