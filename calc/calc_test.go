package calc_test

import (
	"testing"

	"github.com/hanahmily/tdd-sandbox/calc"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		name string
		a    int
		b    int
		want int
	}{
		{name: "positive values", a: 2, b: 3, want: 5},
		{name: "negative values", a: -2, b: -3, want: -5},
		{name: "mixed signs", a: -7, b: 4, want: -3},
		{name: "zero on the left", a: 0, b: 9, want: 9},
		{name: "zero on the right", a: -9, b: 0, want: -9},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calc.Add(tt.a, tt.b); got != tt.want {
				t.Fatalf("Add(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.want)
			}
		})
	}
}
