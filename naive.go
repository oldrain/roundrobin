package roundrobin

// Naive Round Robin
type NRR struct {
	nodes []*RRNode

	n    int
	curr int
}

func (nrr *NRR) InitRR(nodes []*RRNode) {
	nrr.nodes = nodes
	nrr.n = len(nodes)
	nrr.curr = nrr.n - 1
}

func (nrr *NRR) Next() *RRNode {
	if nrr.n <= 0 {
		return nil
	}

	nrr.curr += 1
	next := nrr.curr % nrr.n

	return nrr.nodes[next]
}
