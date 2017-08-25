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

// FindWinningBids finds the winning bids for an auction
func FindWinningBids(bids bid.Bids, totalPercentage int) (bid.Bids, error) {
	var divisible int

	thresholds := flatThresholds(bids)
	sort.Sort(bids)

	if len(thresholds) == 1 {
		var n int
		for t := range thresholds {
			n = t
		}
		divisible = totalPercentage / n
	}

	var winners bid.Bids
	if divisible != 0 {
		for i, b := range bids {
			if i >= divisible {
				break
			}
			winners = append(winners, b)
		}
	} else {
		for t, v := range thresholds {
			for _, b := range bids {
				if b.Threshold == t && b.Value == v {
					winners = append(winners, b)
				}
			}
		}
	}

	return winners, nil
}

func flatThresholds(bids bid.Bids) map[int]float32 {
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

	return thresholds
}
