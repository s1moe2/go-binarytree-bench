package main

import "testing"
//
//func TestRun10(t *testing.T) {
//	// expected output from
//	// https://benchmarksgame-team.pages.debian.net/benchmarksgame/download/binarytrees-output.txt
//
//	Run(10)
//	// Output:
//	// stretch tree of depth 11	 check: 4095
//	// 1024	 trees of depth 4	 check: 31744
//	// 256	 trees of depth 6	 check: 32512
//	// 64	 trees of depth 8	 check: 32704
//	// 16	 trees of depth 10	 check: 32752
//	// long lived tree of depth 10	 check: 2047
//}

func TestRun21(t *testing.T) {
	Run(21)
}

func BenchmarkRun(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Run(21)
	}
}
