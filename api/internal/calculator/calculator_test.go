package calculator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTotal(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  int
	}{
		{"success - happy path", []int{1, 2, 3}, 6},
		{"success - more nums", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 55},
		{"success - neg nums", []int{-1, -2, -3, -4, -5, -6, -7, -8, -9, -10}, -55},
		{"success - mixed nums", []int{1, 2, 3, -1}, 5},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			c := Calculator{tc.input}
			got := c.Total()
			assert.Equal(t, tc.want, got)
		})
	}
}
