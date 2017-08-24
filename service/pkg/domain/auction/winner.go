package auction

import (
	"sort"

	"github.com/italolelis/epioca/service/pkg/domain/bid"
	"github.com/pkg/errors"
)

var (
	// ErrInvalidPercentage is used when an invalid percentage value is given
	ErrInvalidPercentage = errors.New("invalid percentage given")
)

func FindWinningBids(bids bid.Bids, totalPercentage int) ([]*bid.Bid, error) {
	var divisible int

	thresholds := make(map[int]int)
	for _, b := range bids {
		minValue, ok := thresholds[b.Threshold]
		if !ok {
			thresholds[b.Threshold] = int(b.Value)
			continue
		}

		if int(b.Value) < minValue {
			thresholds[b.Threshold] = int(b.Value)
		}
	}

	if len(thresholds) == 1 {
		var t int
		for k := range thresholds {
			t = k
		}
		divisible = totalPercentage / t
	}

	var winners []*bid.Bid
	if divisible != 0 {
		sort.Sort(bids)
		for i, b := range bids {
			if i >= divisible {
				break
			}
			winners = append(winners, b)
		}
	} else {
		for t, v := range thresholds {
			for _, b := range bids {
				if b.Threshold == t && int(b.Value) == v {
					winners = append(winners, b)
				}
			}
		}
	}

	return winners, nil
}
