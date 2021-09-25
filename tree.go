package main

import (
	"fmt"
)

type Tree struct {
	Left  *Tree
	Right *Tree
}

// Count the nodes in the given complete binary tree.
func (t *Tree) Count() int {
	// Only test the Left node (this binary tree is expected to be complete).
	if t.Left == nil {
		return 1
	}
	return 1 + t.Right.Count() + t.Left.Count()
}

// NewTree creates a complete binary tree of `depth` and returns it as a pointer.
func NewTree(depth int) *Tree {
	if depth > 0 {
		return &Tree{Left: NewTree(depth - 1), Right: NewTree(depth - 1)}
	} else {
		return &Tree{}
	}
}

type treeTestResult struct {
	index int
	msg   string
}

func Run(maxDepth int) {
	res := make(chan *treeTestResult)

	var longLivedTree *Tree
	lltRes := make(chan *Tree)

	// Set minDepth to 4 and maxDepth to the maximum of maxDepth and minDepth +2.
	const minDepth = 4
	if maxDepth < minDepth+2 {
		maxDepth = minDepth + 2
	}

	// Create an indexed string buffer for outputing the result in order.
	outCurr := 0
	outSize := 3 + (maxDepth-minDepth)/2
	outBuff := make([]string, outSize)

	// Create a long-lived binary tree of depth maxDepth. Its statistics will be handled later.
	go func(r chan<- *Tree) {
		longLivedTree = NewTree(maxDepth)
		r <- longLivedTree
		close(r)
	}(lltRes)

	var goroutines int

	var tasks []*Task
	


	pool := NewPool()

	// Create binary tree of depth maxDepth+1, compute its Count and set the
	// first position of the outputBuffer with its statistics.
	goroutines++
	go func(r chan<- *treeTestResult) {
		tree := NewTree(maxDepth + 1)
		msg := fmt.Sprintf("stretch tree of depth %d\t check: %d",
			maxDepth+1,
			tree.Count())

		r <- &treeTestResult{
			index: 0,
			msg:   msg,
		}
	}(res)

	// Create a lot of binary trees, of depths ranging from minDepth to maxDepth,
	// compute and tally up all their Count and record the statistics.
	for depth := minDepth; depth <= maxDepth; depth += 2 {
		goroutines++
		iterations := 1 << (maxDepth - depth + minDepth)
		outCurr++

		go func(depth, iterations, index int, r chan<- *treeTestResult) {
			acc := 0
			for i := 0; i < iterations; i++ {
				// Create a binary tree of depth and accumulate total counter with its node count.
				a := NewTree(depth)
				acc += a.Count()
			}
			msg := fmt.Sprintf("%d\t trees of depth %d\t check: %d",
				iterations,
				depth,
				acc)

			r <- &treeTestResult{
				index: index,
				msg: msg,
			}
		}(depth, iterations, outCurr, res)
	}

	for i := 0; i < goroutines; i++ {
		ir := <- res
		outBuff[ir.index] = ir.msg
	}

	// Compute the checksum of the long-lived binary tree that we created earlier and store its statistics.
	longLivedTree = <- lltRes
	msg := fmt.Sprintf("long lived tree of depth %d\t check: %d",
		maxDepth,
		longLivedTree.Count())
	outBuff[outSize-1] = msg

	// Print the statistics for all of the various tree depths.
	for _, m := range outBuff {
		fmt.Println(m)
	}
}

