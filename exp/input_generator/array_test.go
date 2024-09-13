package input_generator

import (
	"golang.org/x/exp/constraints"
	"reflect"
	"testing"
)

func TestArray(t *testing.T) {
	type args struct {
		arrayString string
	}
	type testCase[T constraints.Ordered] struct {
		name    string
		args    args
		want    []T
		wantErr bool
	}
	tests := []testCase[int]{
		// TODO: Add test cases.
		testCase[int]{
			name:    "test 1",
			args:    args{"[1,0,3,4]"},
			want:    []int{1, 0, 3, 4},
			wantErr: false,
		},
		testCase[int]{
			name:    "test 2",
			args:    args{"1,0,3,4"},
			want:    []int{1, 0, 3, 4},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Array[int](tt.args.arrayString)
			if (err != nil) != tt.wantErr {
				t.Errorf("Array() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Array() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArrays(t *testing.T) {
	type args struct {
		arrayStrings []string
	}
	type testCase[T constraints.Ordered] struct {
		name    string
		args    args
		want    [][]T
		wantErr bool
	}
	tests := []testCase[int]{
		testCase[int]{
			name:    "test 1",
			args:    args{[]string{"[1,0,3,4]", "[6,3]"}},
			want:    [][]int{{1, 0, 3, 4}, {6, 3}},
			wantErr: false,
		},
		testCase[int]{
			name:    "test 1",
			args:    args{[]string{"1,0,3,4", "6,3"}},
			want:    [][]int{{1, 0, 3, 4}, {6, 3}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Arrays[int](tt.args.arrayStrings...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Arrays() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Arrays() got = %v, want %v", got, tt.want)
			}
		})
	}
}
