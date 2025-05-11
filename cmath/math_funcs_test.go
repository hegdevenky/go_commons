package cmath

import (
	"golang.org/x/exp/constraints"
	"iter"
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

func TestPercentageOf(t *testing.T) {
	type args[T interface {
		constraints.Integer | constraints.Float
	}] struct {
		p T
		x T
	}
	type testCase[T interface {
		constraints.Integer | constraints.Float
	}] struct {
		name string
		args args[T]
		want float64
	}

	// testing float64
	tests := []testCase[float64]{
		{"test1", args[float64]{3, 91}, 2.73},
		{"test2", args[float64]{5, 91}, 4.55},
		{"test3", args[float64]{0, 91}, 0},
		{"test4", args[float64]{7, 0}, 0},
		{"test5", args[float64]{0, 0}, 0},
		{"test6", args[float64]{2.5, 68.65}, 1.7162500000000003},
		{"test7", args[float64]{10, 3.6872}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PercentageOf(tt.args.p, tt.args.x); got != tt.want {
				t.Errorf("PercentageOf() = %v, want %v", got, tt.want)
			}
		})
	}

	// testing int
	testInts := []testCase[int]{
		{"test1", args[int]{3, 91}, 2.73},
		{"test2", args[int]{5, 91}, 4.55},
		{"test3", args[int]{0, 91}, 0},
		{"test4", args[int]{7, 0}, 0},
		{"test5", args[int]{0, 0}, 0},
	}
	for _, tt := range testInts {
		t.Run(tt.name, func(t *testing.T) {
			if got := PercentageOf(tt.args.p, tt.args.x); got != tt.want {
				t.Errorf("PercentageOf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInts(t *testing.T) {
	i := 2
	for next, stop := iter.Pull(Ints(i, 5)); i <= 5; i++ {
		if val, ok := next(); !ok {
			stop()
			break
		} else {
			t.Log("TestInts: (2,5):", val)
			if val != i {
				t.Errorf("got = %v, expected %v", val, i)
			}
		}
	}

	i = 7
	for next, stop := iter.Pull(Ints(i, 2)); i >= 2; i-- {
		if val, ok := next(); !ok {
			stop()
			break
		} else {
			t.Log("TestInts: (7,2):", val)
			if val != i {
				t.Errorf("got = %v, expected %v", val, i)
			}
		}
	}

	i = 3
	for next, stop := iter.Pull(Ints(i, i)); i >= 3; i++ {
		if val, ok := next(); !ok {
			stop()
			break
		} else {
			t.Log("TestInts: (3,3):", val)
			if val != i {
				t.Log("val:", val)
				t.Errorf("got = %v, expected %v", val, i)
			}
		}
	}

	i = -5
	for next, stop := iter.Pull(Ints(i, 5)); i <= 5; i++ {
		if val, ok := next(); !ok {
			stop()
			break
		} else {
			t.Log("TestInts: (-5,5):", val)
			if val != i {
				t.Errorf("got = %v, expected %v", val, i)
			}
		}
	}
}
