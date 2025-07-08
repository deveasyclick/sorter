package sorters

import "github.com/deveasyclick/sorter/types"

type ProductSorter interface {
	Sort(data []types.Product) []types.Product
	SetSortStrategy(strategy types.SortStrategy[types.Product])
}

type productSorter struct {
	sortStrategy types.SortStrategy[types.Product]
}

func (p *productSorter) Sort(data []types.Product) []types.Product {
	if p.sortStrategy == nil {
		return data
	}
	return p.sortStrategy.Sort(data)
}

func (p *productSorter) SetSortStrategy(strategy types.SortStrategy[types.Product]) {
	p.sortStrategy = strategy
}

func NewProduct(sortStrategy types.SortStrategy[types.Product]) ProductSorter {
	return &productSorter{
		sortStrategy: sortStrategy,
	}
}
