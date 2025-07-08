package strategies

import (
	"sort"

	"github.com/deveasyclick/sorter/types"
)

type SortByPrice struct{}

func (s SortByPrice) Sort(products []types.Product) []types.Product {
	sort.Slice(products, func(i, j int) bool {
		return products[i].Price < products[j].Price
	})
	return products
}
