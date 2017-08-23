package auction

import (
	"time"

	"github.com/lib/pq"
	uuid "github.com/satori/go.uuid"
)

const (
	//Running status
	Running string = "running"
	//Scheduled status
	Scheduled string = "scheduled"
	//Completed status
	Completed string = "completed"
)

//Status represents the auction status
type Status string

// Repository defines the basic methods for a repository
type Repository interface {
	Find(status string) ([]*Auction, error)
	FindByID(id uuid.UUID) (*Auction, error)
	Add(auction *Auction) error
	Remove(id uuid.UUID) error
}

// Auction represents an auction
type Auction struct {
	ID         uuid.UUID     `json:"id" db:"id"`
	Status     string        `json:"status" db:"status"`
	Week       string        `json:"week" db:"week"`
	Country    string        `json:"country" db:"country"`
	DC         string        `json:"dc" db:"dc"`
	Ingredient string        `json:"ingredient" db:"ingredient"`
	Duration   time.Duration `json:"duration" db:"duration"`
	StartDate  time.Time     `json:"start_date" db:"start_date"`
	EndDate    time.Time     `json:"end_date" db:"end_date"`
	Qty        float32       `json:"qty" db:"qty"`
	Threshold  pq.Int64Array `json:"threshold" db:"threshold"`
	MaxPrice   float32       `json:"max_price" db:"max_price"`
}

// NewAuction creates a new instance of auction
func NewAuction(id uuid.UUID) *Auction {
	return &Auction{ID: id}
}

// IsScheduled verifies if an auction is scheduled
func (a *Auction) IsScheduled() bool {
	return a.Status == Scheduled
}
