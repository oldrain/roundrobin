package test

import (
	"roundrobin"
	"fmt"
	"testing"
)

func TestRW_Next(t *testing.T) {
	rr := new(roundrobin.RW)
	rr.InitRR(testNodes)

	results := make(map[string]int)

	for i := 0; i < 1000; i++ {
		node := rr.Next()
		results[node.Item.(string)]++
	}

	if !checkResultRange(results["a"], 400, 600) ||
		!checkResultRange(results["b"], 200, 400) ||
		!checkResultRange(results["c"], 100, 300) {
		t.Error("Algorithm is wrong", fmt.Sprintf("[a: %d, b: %d, c: %d]",
			results["a"], results["b"], results["c"]))
	}
}
