package types

type SortStrategy[T any] interface {
	Sort(data []T) []T
}

type Product struct {
	ID         int
	Name       string
	Price      float64
	Created    string
	SalesCount int
	ViewsCount int
}
