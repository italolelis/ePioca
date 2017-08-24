package auction

import (
	"github.com/italolelis/epioca/service/pkg/domain/bid"
	"github.com/pkg/errors"
)

var (
	// ErrInvalidPercentage is used when an invalid percentage value is given
	ErrInvalidPercentage = errors.New("invalid percentage given")
)

func findWinningBid(bids []*bid.Bid, totalPercentage int) ([]*bid.Bid, error) {
	thresholds := make(map[int]float32)
	for _, b := range bids {
		minValue, ok := thresholds[b.Threshold]
		if !ok {
			thresholds[b.Threshold] = b.Value
			continue
		}

		if b.Value < minValue {
			thresholds[b.Threshold] = b.Value
		}
	}

	var total int
	for t := range thresholds {
		total += t
	}

	if total != totalPercentage {
		return nil, ErrInvalidPercentage
	}

	var winners []*bid.Bid
	for t, v := range thresholds {
		for _, b := range bids {
			if b.Threshold == t && b.Value == v {
				winners = append(winners, b)
			}
		}
	}

	return winners, nil
}
