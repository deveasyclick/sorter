package strategies

import (
	"testing"

	"github.com/deveasyclick/sorter/types"
)

func TestSortBySalesViewRatio(t *testing.T) {
	tests := []struct {
		name     string
		input    []types.Product
		expected []int // Expected order of IDs
	}{
		{
			name: "basic ratio sort",
			input: []types.Product{
				{ID: 1, SalesCount: 100, ViewsCount: 200}, // 0.5
				{ID: 2, SalesCount: 50, ViewsCount: 50},   // 1.0
				{ID: 3, SalesCount: 10, ViewsCount: 100},  // 0.1
			},
			expected: []int{2, 1, 3}, // highest ratio first
		},
		{
			name: "equal ratios",
			input: []types.Product{
				{ID: 4, SalesCount: 100, ViewsCount: 200}, // 0.5
				{ID: 5, SalesCount: 50, ViewsCount: 100},  // 0.5
			},
			expected: []int{4, 5},
		},
		{
			name: "zero views (should handle safely)",
			input: []types.Product{
				{ID: 6, SalesCount: 10, ViewsCount: 0},  // divide by zero — should treat as 0
				{ID: 7, SalesCount: 10, ViewsCount: 10}, // ratio = 1.0
			},
			expected: []int{7, 6},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sorted := SortBySalesViewRatio{}.Sort(tt.input)

			for i, p := range sorted {
				if p.ID != tt.expected[i] {
					t.Errorf("expected ID %d at position %d but got ID %d", tt.expected[i], i, p.ID)
				}
			}
		})
	}
}
