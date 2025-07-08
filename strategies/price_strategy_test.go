package strategies

import (
	"testing"

	"github.com/deveasyclick/sorter/types"
)

func TestSortByPriceCases(t *testing.T) {
	tests := []struct {
		name     string
		input    []types.Product
		expected []int
	}{
		{
			name: "first case",
			input: []types.Product{
				{ID: 1, Price: 50},
				{ID: 2, Price: 10},
				{ID: 3, Price: 30},
			},
			expected: []int{2, 3, 1},
		},
		{
			name: "second case",
			input: []types.Product{
				{ID: 1, Price: 1},
				{ID: 2, Price: 0},
				{ID: 3, Price: 0},
			},
			expected: []int{2, 3, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sorted := SortByPrice{}.Sort(tt.input)
			for i, p := range sorted {
				if p.ID != tt.expected[i] {
					t.Errorf("expected ID %d at position %d but got ID %d", tt.expected[i], i, p.ID)
				}
			}
		})
	}
}
