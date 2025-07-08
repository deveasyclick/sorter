package strategies

import (
	"sort"

	"github.com/deveasyclick/sorter/types"
)

type SortBySalesViewRatio struct{}

func (s SortBySalesViewRatio) Sort(products []types.Product) []types.Product {
	// Return early if there are no products or only one product
	if len(products) < 2 {
		return products
	}

	sort.Slice(products, func(i, j int) bool {
		return ratio(products[i]) > ratio(products[j]) // descending
	})

	return products
}

func ratio(p types.Product) float64 {
	if p.ViewsCount <= 0 {
		return 0
	}
	return float64(p.SalesCount) / float64(p.ViewsCount)
}
