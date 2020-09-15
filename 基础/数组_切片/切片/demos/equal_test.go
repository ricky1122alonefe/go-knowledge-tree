package main

import "testing"

type testSlice struct {
	ins1 []string
	ins2 []string
	out  bool
}

var tests = []testSlice{
	{ins1: []string{"a", "b", "c"}, ins2: []string{"a", "b", "c"}, out: true},
	{ins1: []string{"a", "b", "c"}, ins2: []string{"a", "b", "c", "d"}, out: false},
	{ins1: []string{"a", "b", "c"}, ins2: []string{"a", "b", "d"}, out: false},
}

func BenchmarkReflectSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, value := range tests {
			if value.out != ReflectSlice(value.ins1, value.ins2) {
				b.Error("test failed")
			}
		}
	}
}

func BenchmarkCompareSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, value := range tests {
			if value.out != CompareSlice(value.ins1, value.ins2) {
				b.Error("test failed")
			}
		}
	}
}
