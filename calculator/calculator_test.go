package calculator_test

import (
	"testing"

	"github.com/Julianrt/test/calculator"
)

func TestSum(t *testing.T) {
	cal := calculator.Calculator{}

	data := []struct {
		name string
		a    int
		b    int
		want int
	}{
		{"1", 2, 3, 5},
		{"2", -7, 10, 3},
		{"3", 47, 15, 62},
		{"4", -10, 0, -10},
		{"5", 10, -15, -5},
		{"6", 123, 321, 444},
		{"7", 12, 46, 58},
		{"8", 100, -1, 99},
		{"9", -15, 20, 5},
		{"10", -10, -10, -20},
	}

	for _, v := range data {
		t.Run(v.name, func(t *testing.T) {
			got := cal.Sum(v.a, v.b)
			if got != v.want {
				t.Errorf("ERROR: expected value %d and got %d", v.want, got)
			}
		})
	}
}

func TestSubtract(t *testing.T) {
	cal := calculator.Calculator{}

	data := []struct {
		name string
		a    int
		b    int
		want int
	}{
		{"1", 5, 10, -5},
		{"2", -7, 10, -17},
		{"3", 50, -15, 65},
		{"4", 27, 24, 3},
		{"5", 4, -15, 19},
		{"6", 30, 15, 15},
		{"7", 3, 1, 2},
		{"8", 490, 499, -9},
		{"9", -10, 20, -30},
		{"10", 30, 80, -50},
	}

	for _, v := range data {
		t.Run(v.name, func(t *testing.T) {
			got := cal.Subtract(v.a, v.b)
			if got != v.want {
				t.Errorf("ERROR: expected value %d and got %d", v.want, got)
			}
		})
	}
}

func TestMultiply(t *testing.T) {
	cal := calculator.Calculator{}

	data := []struct {
		name string
		a    int
		b    int
		want int
	}{
		{"1", 10, 10, 100},
		{"2", -15, -15, 225},
		{"3", 20, -5, -100},
		{"4", -7, 8, -56},
		{"5", 50, 0, 0},
		{"6", 0, 20, 0},
		{"7", 20, 20, 400},
		{"8", 2, 2, 4},
		{"9", 1, 1, 1},
		{"10", 5, -5, -25},
	}

	for _, v := range data {
		t.Run(v.name, func(t *testing.T) {
			got := cal.Multiply(v.a, v.b)
			if got != v.want {
				t.Errorf("ERROR: expected value %d and got %d", v.want, got)
			}
		})
	}
}

func TestDivide(t *testing.T) {
	cal := calculator.Calculator{}

	data := []struct {
		name string
		a    float64
		b    float64
		want float64
	}{
		{"1", 1, 0, 0},
		{"2", -15, -15, 1},
		{"3", 20, -5, -4},
		{"4", -4, 8, -0.5},
		{"5", 123, 4, 30.75},
		{"6", 85, 5, 17},
		{"7", -698, -8, 87.25},
		{"8", -15, 3, -5},
		{"9", 740, 8, 92.5},
		{"10", -32, 0.8, -40},
	}

	for _, v := range data {
		t.Run(v.name, func(t *testing.T) {
			got, err := cal.Divide(v.a, v.b)

			if v.b == 0 {
				if err == nil {
					t.Errorf("Error: expected an error when divisor is 0")
				}
				return
			}

			if got != v.want {
				t.Errorf("ERROR: expected value %.2f and got %.2f", v.want, got)
			}
		})
	}
}
