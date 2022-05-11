package main

import "testing"

func BenchmarkRandInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DNAStrand("AAAA")
	}
}
