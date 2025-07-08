package strategies

import (
	"sort"

	"github.com/deveasyclick/sorter/types"
)

type SortByPrice struct{}

func (s SortByPrice) Sort(products []types.Product) []types.Product {
	// Return early if there are no products or only one product
	if len(products) < 2 {
		return products
	}

	sort.Slice(products, func(i, j int) bool {
		return products[i].Price < products[j].Price
	})
	return products
}
