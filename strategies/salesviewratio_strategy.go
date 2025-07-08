package strategies

import (
	"sort"

	"github.com/deveasyclick/sorter/types"
)

type SortBySalesViewRatio struct{}

func (s SortBySalesViewRatio) Sort(products []types.Product) []types.Product {
	sort.Slice(products, func(i, j int) bool {
		return ratio(products[i]) > ratio(products[j]) // descending
	})

	return products
}

func ratio(p types.Product) float64 {
	if p.ViewsCount == 0 {
		return 0 // or a sentinel value like -1 if you want to treat it as lowest
	}
	return float64(p.SalesCount) / float64(p.ViewsCount)
}
