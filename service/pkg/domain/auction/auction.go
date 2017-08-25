package auction

import (
	"time"

	"github.com/lib/pq"
	uuid "github.com/satori/go.uuid"
)

const (
	//Running auction status
	Running string = "running"
	//Scheduled auction status
	Scheduled string = "scheduled"
	//Completed auction status
	Completed string = "completed"
)

// IsValidStatus checks if the status is a valid one
func IsValidStatus(s string) bool {
	return s == Running || s == Scheduled || s == Completed
}

// Repository defines the basic methods for a repository
type Repository interface {
	Find(status string) ([]*Auction, error)
	FindByID(id uuid.UUID) (*Auction, error)
	Add(auction *Auction) error
	Remove(id uuid.UUID) error
}

// Ingredient represents a HelloFresh ingredient
type Ingredient struct {
	Name string `json:"name" db:"name"`
	SKU  string `json:"sku" db:"sku"`
}

// Auction represents an auction
type Auction struct {
	ID             uuid.UUID     `json:"id" db:"id"`
	Week           string        `json:"week" db:"week"`
	Country        string        `json:"country" db:"country"`
	DC             string        `json:"dc" db:"dc"`
	Ingredient     *Ingredient   `json:"ingredient" db:"ingredient"`
	Duration       time.Duration `json:"duration" db:"duration"`
	StartDate      time.Time     `json:"start_date" db:"start_date"`
	Qty            float32       `json:"qty" db:"qty"`
	Threshold      pq.Int64Array `json:"threshold" db:"threshold"`
	MaxPrice       float32       `json:"max_price" db:"max_price"`
	PriceIncrement float32       `json:"price_increment" db:"price_increment"`
	VStatus        string        `json:"status, omitempty" db:"-"`
	VTimeRemaining float64       `json:"time_remaining, omitempty" db:"-"`
}

func (a *Auction) TimeRemaining() time.Duration {
	return a.StartDate.Sub(time.Now()) + a.Duration*time.Second
}

// NewAuction creates a new instance of auction
func NewAuction(id uuid.UUID) *Auction {
	return &Auction{ID: id}
}

// Status verifies if an auction is scheduled
func (a *Auction) Status() string {
	if time.Now().After(a.StartDate.Add(a.Duration * time.Second)) {
		return Completed
	} else if time.Now().After(a.StartDate) {
		return Running
	}
	return Scheduled
}

// IsScheduled verifies if an auction is scheduled
func (a *Auction) IsScheduled() bool {
	return a.Status() == Scheduled
}

// IsCompleted verifies if an auction is Completed
func (a *Auction) IsCompleted() bool {
	return a.Status() == Completed
}

// IsThresholdValid verifies if the sum of all thresholds will be 100 percent
func (a *Auction) IsThresholdValid() bool {
	var total int64
	for _, t := range a.Threshold {
		total += t
	}

	return total == 100
}
