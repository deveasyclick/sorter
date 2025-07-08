package strategies

import (
	"sort"

	"github.com/deveasyclick/sorter/types"
)

type SortBySalesViewRatio struct{}

func (s SortBySalesViewRatio) Sort(products []types.Product) []types.Product {
	sort.Slice(products, func(i, j int) bool {
		ratioI := float64(products[i].SalesCount) / float64(products[i].ViewsCount)
		ratioJ := float64(products[j].SalesCount) / float64(products[j].ViewsCount)
		return ratioI > ratioJ // descending order
	})
	return products
}
