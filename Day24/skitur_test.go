package main

import "testing"

func BenchmarkSkitur(b *testing.B) {
	for i := 0; i < b.N; i++ {
		main()
	}
}
