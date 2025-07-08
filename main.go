package main

import (
	"log/slog"
	"os"

	"github.com/deveasyclick/sorter/sorters"
	"github.com/deveasyclick/sorter/strategies"
	"github.com/deveasyclick/sorter/types"
)

func main() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))
	products := []types.Product{
		{ID: 1, Name: "Alabaster Table", Price: 12.99, Created: "2019-01-04", SalesCount: 32, ViewsCount: 730},
		{ID: 2, Name: "Zebra Table", Price: 44.49, Created: "2012-01-04", SalesCount: 301, ViewsCount: 3279},
		{ID: 3, Name: "Coffee Table", Price: 10.00, Created: "2014-05-28", SalesCount: 1048, ViewsCount: 20123},
	}

	sorter := sorters.NewProduct(strategies.SortByPrice{})
	sortedProducts := sorter.Sort(products)
	slog.Info("Sorted by Price", "products", sortedProducts)

	sorter.SetSortStrategy(strategies.SortBySalesViewRatio{})
	sortedProducts = sorter.Sort(products)
	slog.Info("Sorted by Sales/View Ratio", "products", sortedProducts)
}
