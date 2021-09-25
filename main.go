// The Computer Language Benchmarks Game
// http://benchmarksgame.alioth.debian.org/
//
// Go implementation of binary-trees, based on the reference implementation
// gcc #3, on Go #8 (which is based on Rust #4) as the following links, below:
// - https://benchmarksgame-team.pages.debian.net/benchmarksgame/program/binarytrees-gcc-3.html
// - https://benchmarksgame-team.pages.debian.net/benchmarksgame/program/binarytrees-go-8.html
// - https://benchmarksgame-team.pages.debian.net/benchmarksgame/program/binarytrees-rust-4.html
//
// Comments aim to be analogous as those in the reference implementation and are
// intentionally verbose, to help programmers unexperienced in GO to understand
// the implementation.
//
// The following alternative implementations were considered before submitting
// this code. All of them had worse readability and didn't yield better results
// on my local machine:
//
// 0. general:
// 0.1 using uint32, instead of int;
//
// 1. func Count:
// 1.1 using a local stack, instead of using a recursive implementation; the
//     performance degraded, even using a pre-allocated slice as stack and
//     manually handling its size;
// 1.2 assigning Left and Right to nil after counting nodes; the idea to remove
//     references to instances no longer needed was to make GC easier, but this
//     did not work as intended;
// 1.3 using a walker and channel, sending 1 on each node; although this looked
//     idiomatic to GO, the performance suffered a lot;
// 2. func NewTree:
// 2.1 allocating all tree nodes on a tree slice upfront and making references
//     to those instances, instead of allocating two sub-trees on each call;
//     this did not improve performance;
//
// Contributed by Gerardo Lima (https://github.com/gerardolima)
// Based on previous work from Adam Shaver, Isaac Gouy, Marcel Ibes Jeremy,
//  Zerfas, Jon Harrop, Alex Mizrahi, Bruno Coutinho, ...
//

package main

import (
	"flag"
	"strconv"
)

func main() {
	n := 0
	flag.Parse()
	if flag.NArg() > 0 {
		n, _ = strconv.Atoi(flag.Arg(0))
	}

	Run(n)
}
