package maps

import (
	"reflect"
	"testing"
)

func TestGetOrDefault(t *testing.T) {
	type args[K comparable, V any] struct {
		m            map[K]V
		key          K
		defaultValue V
	}
	type testCase[K comparable, V any] struct {
		name string
		args args[K, V]
		want V
	}
	tests := []testCase[string, int]{
		testCase[string, int]{
			name: "test 1",
			args: args[string, int]{
				m: map[string]int{
					"A": 65,
					"Z": 90,
					"a": 97,
					"z": 122,
				},
				key:          "m",
				defaultValue: 10,
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetOrDefault(tt.args.m, tt.args.key, tt.args.defaultValue); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetOrDefault() = %v, want %v", got, tt.want)
			}
		})
	}
}
