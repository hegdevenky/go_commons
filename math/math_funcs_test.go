package math

import (
	"testing"
)

func TestAbs(t *testing.T) {
	var tests = []struct {
		name  string
		input int
		want  int
	}{
		{"Abs of -2 should be 2", -2, 2},
		{"Abs of 2 should be 2", 2, 2},
		{"Abs of 0 should be 0", -0, 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ans := Abs(test.input)
			if ans != test.want {
				t.Errorf("got %d, want %d", ans, test.want)
			}
		})
	}

	var testFloats = []struct {
		name  string
		input float64
		want  float64
	}{
		{"Abs of -8.345 should be 8.345", -8.345, 8.345},
		{"Abs of 9 should be 9", 9, 9},
		{"Abs of 0 should be 0", 0, 0},
	}

	for _, test := range testFloats {
		t.Run(test.name, func(t *testing.T) {
			ans := Abs(test.input)
			if ans != test.want {
				t.Errorf("got %f, want %f", ans, test.want)
			}
		})
	}
}
