package auction

import (
	"testing"

	"github.com/italolelis/epioca/service/pkg/domain/bid"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWinningBid(t *testing.T) {
	supplierA := &bid.User{ID: uuid.NewV4(), Name: "Supplier A"}
	supplierB := &bid.User{ID: uuid.NewV4(), Name: "Supplier B"}
	supplierC := &bid.User{ID: uuid.NewV4(), Name: "Supplier C"}

	bids := []*bid.Bid{
		&bid.Bid{
			Threshold: 25,
			Value:     60,
			User:      supplierA,
		},
		&bid.Bid{
			Threshold: 25,
			Value:     50,
			User:      supplierB,
		},
		&bid.Bid{
			Threshold: 25,
			Value:     45,
			User:      supplierC,
		},
		&bid.Bid{
			Threshold: 25,
			Value:     67,
			User:      supplierC,
		},
		&bid.Bid{
			Threshold: 25,
			Value:     15,
			User:      supplierC,
		},
		&bid.Bid{
			Threshold: 25,
			Value:     90,
			User:      supplierC,
		},
		&bid.Bid{
			Threshold: 75,
			Value:     120,
			User:      supplierC,
		},
		&bid.Bid{
			Threshold: 75,
			Value:     100,
			User:      supplierA,
		},
		&bid.Bid{
			Threshold: 75,
			Value:     90,
			User:      supplierB,
		},
		&bid.Bid{
			Threshold: 75,
			Value:     200,
			User:      supplierB,
		},
		&bid.Bid{
			Threshold: 75,
			Value:     60,
			User:      supplierB,
		},
		&bid.Bid{
			Threshold: 75,
			Value:     61,
			User:      supplierB,
		},
	}

	winners, err := FindWinningBids(bids, 100)
	require.NoError(t, err)
	assert.Len(t, winners, 2)
}

func TestWinningBid2(t *testing.T) {
	supplierA := &bid.User{ID: uuid.NewV4(), Name: "Supplier A"}
	supplierB := &bid.User{ID: uuid.NewV4(), Name: "Supplier B"}

	bids := []*bid.Bid{
		&bid.Bid{
			Threshold: 25,
			Value:     60,
			User:      supplierA,
		},
		&bid.Bid{
			Threshold: 25,
			Value:     50,
			User:      supplierB,
		},
		&bid.Bid{
			Threshold: 75,
			Value:     60,
			User:      supplierA,
		},
		&bid.Bid{
			Threshold: 75,
			Value:     70,
			User:      supplierB,
		},
	}

	winners, err := FindWinningBids(bids, 100)
	require.NoError(t, err)
	assert.Len(t, winners, 2)
}

func TestWinningBidEvenThreshold(t *testing.T) {
	supplierA := &bid.User{ID: uuid.NewV4(), Name: "Supplier A"}
	supplierB := &bid.User{ID: uuid.NewV4(), Name: "Supplier B"}
	supplierC := &bid.User{ID: uuid.NewV4(), Name: "Supplier C"}
	supplierD := &bid.User{ID: uuid.NewV4(), Name: "Supplier D"}

	bids := []*bid.Bid{
		&bid.Bid{
			Threshold: 50,
			Value:     60,
			User:      supplierA,
		},
		&bid.Bid{
			Threshold: 50,
			Value:     50,
			User:      supplierB,
		},
		&bid.Bid{
			Threshold: 50,
			Value:     61,
			User:      supplierC,
		},
		&bid.Bid{
			Threshold: 50,
			Value:     70,
			User:      supplierD,
		},
	}

	winners, err := FindWinningBids(bids, 100)
	require.NoError(t, err)
	assert.Len(t, winners, 2)
}

func TestWinningBid4(t *testing.T) {
	supplierA := &bid.User{ID: uuid.NewV4(), Name: "Supplier A"}
	supplierB := &bid.User{ID: uuid.NewV4(), Name: "Supplier B"}

	bids := []*bid.Bid{
		&bid.Bid{
			Threshold: 25,
			Value:     60,
			User:      supplierA,
		},
		&bid.Bid{
			Threshold: 25,
			Value:     61,
			User:      supplierB,
		},
		&bid.Bid{
			Threshold: 25,
			Value:     62,
			User:      supplierA,
		},
		&bid.Bid{
			Threshold: 25,
			Value:     63,
			User:      supplierA,
		},
	}

	winners, err := FindWinningBids(bids, 100)
	require.NoError(t, err)
	assert.Len(t, winners, 4)
}

func TestWinningBidWithFloatValues(t *testing.T) {
	supplierA := &bid.User{ID: uuid.NewV4(), Name: "Supplier A"}
	supplierB := &bid.User{ID: uuid.NewV4(), Name: "Supplier B"}

	bids := []*bid.Bid{
		&bid.Bid{
			Threshold: 25,
			Value:     2.40,
			User:      supplierA,
		},
		&bid.Bid{
			Threshold: 25,
			Value:     2.50,
			User:      supplierB,
		},
		&bid.Bid{
			Threshold: 45,
			Value:     1.23,
			User:      supplierA,
		},
		&bid.Bid{
			Threshold: 45,
			Value:     1.19,
			User:      supplierB,
		},
		&bid.Bid{
			Threshold: 30,
			Value:     0.95,
			User:      supplierA,
		},
		&bid.Bid{
			Threshold: 30,
			Value:     1.15,
			User:      supplierB,
		},
	}

	winners, err := FindWinningBids(bids, 100)
	require.NoError(t, err)
	assert.Len(t, winners, 3)
}
