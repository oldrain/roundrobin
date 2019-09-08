package test

import (
	"roundrobin"
	"fmt"
	"testing"
)

func TestWRR_Next(t *testing.T) {
	rr := new(roundrobin.WRR)
	rr.InitRR(testNodes)

	results := make(map[string]int)

	for i := 0; i < 1000; i++ {
		node := rr.Next()
		results[node.Item.(string)]++
	}

	if results["a"] != 500 ||
		results["b"] != 300 ||
		results["c"] != 200 {
		t.Error("Algorithm is wrong", fmt.Sprintf("[a: %d, b: %d, c: %d]",
			results["a"], results["b"], results["c"]))
	}
}