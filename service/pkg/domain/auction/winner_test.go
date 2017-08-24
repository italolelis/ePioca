package auction

import (
	"testing"

	"github.com/italolelis/epioca/service/pkg/domain/bid"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWinningBid(t *testing.T) {
	supplierA := uuid.NewV4()
	supplierB := uuid.NewV4()
	supplierC := uuid.NewV4()

	bids := []*bid.Bid{
		&bid.Bid{
			Threshold: 25,
			Value:     60,
			UserID:    supplierA,
		},
		&bid.Bid{
			Threshold: 25,
			Value:     50,
			UserID:    supplierB,
		},
		&bid.Bid{
			Threshold: 25,
			Value:     45,
			UserID:    supplierC,
		},
		&bid.Bid{
			Threshold: 25,
			Value:     67,
			UserID:    supplierC,
		},
		&bid.Bid{
			Threshold: 25,
			Value:     15,
			UserID:    supplierC,
		},
		&bid.Bid{
			Threshold: 25,
			Value:     90,
			UserID:    supplierC,
		},
		&bid.Bid{
			Threshold: 75,
			Value:     120,
			UserID:    supplierC,
		},
		&bid.Bid{
			Threshold: 75,
			Value:     100,
			UserID:    supplierA,
		},
		&bid.Bid{
			Threshold: 75,
			Value:     90,
			UserID:    supplierB,
		},
		&bid.Bid{
			Threshold: 75,
			Value:     200,
			UserID:    supplierB,
		},
		&bid.Bid{
			Threshold: 75,
			Value:     60,
			UserID:    supplierB,
		},
		&bid.Bid{
			Threshold: 75,
			Value:     61,
			UserID:    supplierB,
		},
	}

	winners, err := FindWinningBids(bids, 100)
	require.NoError(t, err)
	assert.Len(t, winners, 2)
}

func TestWinningBid2(t *testing.T) {
	supplierA := uuid.NewV4()
	supplierB := uuid.NewV4()

	bids := []*bid.Bid{
		&bid.Bid{
			Threshold: 25,
			Value:     60,
			UserID:    supplierA,
		},
		&bid.Bid{
			Threshold: 25,
			Value:     50,
			UserID:    supplierB,
		},
		&bid.Bid{
			Threshold: 75,
			Value:     60,
			UserID:    supplierA,
		},
		&bid.Bid{
			Threshold: 75,
			Value:     70,
			UserID:    supplierB,
		},
	}

	winners, err := FindWinningBids(bids, 100)
	require.NoError(t, err)
	assert.Len(t, winners, 2)
}

func TestWinningBidEvenThreshold(t *testing.T) {
	supplierA := uuid.NewV4()
	supplierB := uuid.NewV4()
	supplierC := uuid.NewV4()
	supplierD := uuid.NewV4()

	bids := []*bid.Bid{
		&bid.Bid{
			Threshold: 50,
			Value:     60,
			UserID:    supplierA,
		},
		&bid.Bid{
			Threshold: 50,
			Value:     50,
			UserID:    supplierB,
		},
		&bid.Bid{
			Threshold: 50,
			Value:     61,
			UserID:    supplierC,
		},
		&bid.Bid{
			Threshold: 50,
			Value:     70,
			UserID:    supplierD,
		},
	}

	winners, err := FindWinningBids(bids, 100)
	require.NoError(t, err)
	assert.Len(t, winners, 2)
}

func TestWinningBid4(t *testing.T) {
	supplierA := uuid.NewV4()
	supplierB := uuid.NewV4()

	bids := []*bid.Bid{
		&bid.Bid{
			Threshold: 25,
			Value:     60,
			UserID:    supplierA,
		},
		&bid.Bid{
			Threshold: 25,
			Value:     61,
			UserID:    supplierB,
		},
		&bid.Bid{
			Threshold: 25,
			Value:     62,
			UserID:    supplierA,
		},
		&bid.Bid{
			Threshold: 25,
			Value:     63,
			UserID:    supplierA,
		},
	}

	winners, err := FindWinningBids(bids, 100)
	require.NoError(t, err)
	assert.Len(t, winners, 4)
}
