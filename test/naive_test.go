package test

import (
	"roundrobin"
	"fmt"
	"testing"
)

func TestNRR_Next(t *testing.T) {
	rr := new(roundrobin.NRR)
	rr.InitRR(testNodes)

	results := make(map[string]int)

	for i := 0; i < 999; i++ {
		node := rr.Next()
		results[node.Item.(string)]++
	}

	if results["a"] != 333 ||
		results["b"] != 333 ||
		results["c"] != 333 {
		t.Error("Algorithm is wrong", fmt.Sprintf("[a: %d, b: %d, c: %d]",
			results["a"], results["b"], results["c"]))
	}
}
