package main

import "reflect"

// deepEqual 会做一些类型判断
func ReflectSlice(a, b []string) bool {
	return reflect.DeepEqual(a, b)
}

func CompareSlice(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	if (a == nil) != (b == nil) {
		return false
	}

	for key, value := range a {
		if value != b[key] {
			return false
		}
	}

	return true
}
