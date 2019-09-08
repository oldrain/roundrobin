package roundrobin

import (
	"math/rand"
	"time"
)

// Random Weight Round Robin
type RW struct {
	nodes []*RRNode

	totalW int
	r      *rand.Rand
}

func (rw *RW) InitRR(nodes []*RRNode) {
	rw.nodes = nodes
	rw.totalW = rw.calTotalWeight()
	rw.r = rand.New(rand.NewSource(time.Now().Unix()))
}

func (rw *RW) Next() *RRNode {
	randNum := rw.r.Intn(rw.totalW) + 1

	for _, node := range rw.nodes {
		randNum -= node.Weight
		if randNum <= 0 {
			return node
		}
	}

	return nil
}

func (rw *RW) calTotalWeight() int {
	var count = 0

	for _, node := range rw.nodes {
		count += node.Weight
	}

	return count
}
