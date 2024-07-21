package main

import (
	"testing"
)

func BenchmarkDoIf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		doIf()
	}
}

func BenchmarkDoCase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		doCase()
	}
}

