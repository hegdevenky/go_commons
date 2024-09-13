package input_generator

import (
	"fmt"
	"golang.org/x/exp/constraints"
	"strconv"
	"strings"
)

// Array function parses string representation of an array into golang slice
//
//	arrayString - "[1,0,3,4]" -> []int{1,0,3,4}
//	arrayString - "1,0,3,4" -> []int{1,0,3,4}
func Array[T constraints.Ordered](arrayString string) ([]T, error) {
	arrayString = strings.TrimSpace(strings.Trim(strings.TrimSpace(arrayString), "[]"))
	if arrayString == "" {
		return make([]T, 0), nil
	}

	split := strings.Split(arrayString, ",")
	result := make([]T, len(split))

	var ret any // To wrap the value and then return
	var t T     // To determine the type

	for i, v := range split {
		switch any(t).(type) {
		case int:
			ret, _ = strconv.Atoi(v)
		case float32, float64:
			ret, _ = strconv.ParseFloat(v, 64)
		case bool:
			ret, _ = strconv.ParseBool(v)
		case string:
			ret = v
		default:
			return nil, fmt.Errorf("type %T is not suppored", t)
		}
		result[i] = ret.(T)
	}

	return result, nil
}

// Arrays function parses string representation of an array or array into golang slice of slice.
//
//	arrayStrings - "[1,0,3,4]","[6,3]" -> [][]int{{1,0,3,4},{6,3}}
//	arrayStrings - "1,0,3,4", "6,3" -> [][]int{{1,0,3,4},{6,3}}
func Arrays[T constraints.Ordered](arrayStrings ...string) ([][]T, error) {
	result := make([][]T, len(arrayStrings))
	for i, arrStr := range arrayStrings {
		if array, err := Array[T](arrStr); err != nil {
			return make([][]T, 0), err
		} else {
			result[i] = array
		}
	}
	return result, nil
}
