package test

import (
	"roundrobin"
	"math/rand"
	"time"
)

var testNodes = []*roundrobin.RRNode{
	{
		Item:   "a",
		Weight: 5,
	},
	{
		Item:   "b",
		Weight: 3,
	},
	{
		Item:   "c",
		Weight: 2,
	},
}

var exampleNodes = []*roundrobin.RRNode{
	{
		Item:   "a",
		Weight: 5,
	},
	{
		Item:   "b",
		Weight: 1,
	},
	{
		Item:   "c",
		Weight: 1,
	},
}

var benchmarkNodes []*roundrobin.RRNode

func initBenchmark() {

	r := rand.New(rand.NewSource(time.Now().Unix()))

	for i := 0; i < 16; i++ {
		node := &roundrobin.RRNode{
			Item:   i,
			Weight: r.Intn(8),
		}
		benchmarkNodes = append(benchmarkNodes, node)
	}
}

func checkResultRange(v, min, max int) bool {
	return v >= min && v <= max
}
