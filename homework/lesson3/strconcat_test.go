package main

import (
	"testing"
)

const concatSteps = 10000

func BenchmarkStrConcat(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var str string

		for i := 0; i < concatSteps; i++ {
			str += "x"
		}
	}
}
