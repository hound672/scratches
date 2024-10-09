// run:
// go test -bench . -benchmem
package main

import (
	"testing"
)

func BenchmarkRange(b *testing.B) {
	sl := make([]Obj, 64*1024+2)
	var res int
	for i := 0; i < b.N; i++ {
		res = sumRange(sl)
	}
	_ = res
}

func BenchmarkLoop(b *testing.B) {
	sl := make([]Obj, 64*1024+2)
	var res int
	for i := 0; i < b.N; i++ {
		res = sumLoop(sl)
	}
	_ = res

}