package math

import "testing"

func TestAbs(t *testing.T) {
	type args[T interface{ constraints.Integer | constraints.Integer }] struct {
		val T
	}
	type testCase[T interface{ constraints.Integer | constraints.Integer }] struct {
		name string
		args args[T]
		want T
	}
	tests := []testCase[ /* TODO: Insert concrete types here */ ]{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Abs(tt.args.val); got != tt.want {
				t.Errorf("Abs() = %v, want %v", got, tt.want)
			}
		})
	}
}
