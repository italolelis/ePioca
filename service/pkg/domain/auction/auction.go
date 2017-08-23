package auction

import uuid "github.com/satori/go.uuid"
import "time"

const (
	//Running status
	Running Status = "running"
	//Scheduled status
	Scheduled Status = "scheduled"
	//Completed status
	Completed Status = "completed"
)

//Status represents the auction status
type Status string

// Repository defines the basic methods for a repository
type Repository interface {
	FindAll() ([]*Auction, error)
	Find(id uuid.UUID) (*Auction, error)
	Add(auction *Auction) error
	Remove(id uuid.UUID) error
}

// Auction represents an auction
type Auction struct {
	ID         uuid.UUID     `json:"id" db:"id"`
	Status     Status        `json:"status" db:"status"`
	Week       string        `json:"week" db:"week"`
	Country    string        `json:"country" db:"country"`
	DC         string        `json:"dc" db:"dc"`
	Ingredient string        `json:"ingredient" db:"ingredient"`
	Duration   time.Duration `json:"duration" db:"duration"`
	StartDate  time.Time     `json:"start_date" db:"start_date"`
	EndDate    time.Time     `json:"end_date" db:"end_date"`
	Qty        float32       `json:"qty" db:"qty"`
	Threshold  []int         `json:"threshold" db:"threshold"`
	MaxPrice   float32       `json:"max_price" db:"max_price"`
}
