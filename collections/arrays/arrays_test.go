package arrays

import (
	"testing"
)

//func TestCopyOf(t *testing.T) {
//	type args[T any] struct {
//		elems []T
//	}
//	type testCase[T any] struct {
//		name string
//		args args[T]
//		want []T
//	}
//	tests := []testCase[ /* TODO: Insert concrete types here */ ]{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := CopyOf(tt.args.elems...); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("CopyOf() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

func TestFill(t *testing.T) {
	// fill integer
	array1 := make([]int, 5)
	Fill(array1, 123)
	testFilledSlice(t, array1, 123)

	// reassign the array
	array1 = Fill(array1, 8)
	testFilledSlice(t, array1, 8)

	filled := Fill(nil, 1)
	if filled != nil {
		t.Errorf("Fill() = %v, want %v\n", filled, nil)
	}

	filled1 := Fill([]string{}, "a")
	if filled1 == nil || len(filled1) != 0 {
		t.Errorf("Fill() = %v, want %v\n", filled1, []string{})
	}

}

func testFilledSlice[T comparable](t *testing.T, a []T, want T) {
	for _, val := range a {
		if val != want {
			t.Errorf("Fill() = %v, want %v\n", val, want)
		}
	}
}

//func TestOf(t *testing.T) {
//	type args[T any] struct {
//		elems []T
//	}
//	type testCase[T any] struct {
//		name string
//		args args[T]
//		want []T
//	}
//	tests := []testCase[ /* TODO: Insert concrete types here */ ]{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := Of(tt.args.elems...); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("Of() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
