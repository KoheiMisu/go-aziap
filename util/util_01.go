package util

import (
	"reflect"
)

func Max(target []int) int {
	var result int

	for _, val := range target {
		if result < val {
			result = val
		}
	}

	return result
}

func Min(target []int) int {
	var result int

	for _, val := range target {
		if result > val {
			result = val
		}
	}

	return result
}

func InArray(target interface{}, arr []interface{}) (exists bool) {
	exists = false

	switch reflect.TypeOf(arr).Kind() {
		case reflect.Slice:
			s := reflect.ValueOf(arr)
			len := s.Len()

			for i := 0; i < len; i++ {
				if reflect.DeepEqual(target, s.Index(i).Interface()) == true {
					exists = true
					return
				}
			}
	}

	return
}

func ArrayFill(start, end int, value interface{}) []interface{} {
	var result []interface{}

	for i := start; i < end; i++ {
		result = append(result, value)
	}

	return result
}
