package roundrobin

/**
 * Weighted Round Robin
 * http://kb.linuxvirtualserver.org/wiki/Weighted_Round-Robin_Scheduling
 */
type WRR struct {
	nodes []*RRNode

	n       int
	gcdW    int
	maxW    int
	currPos int
	currW   int
}

func (wrr *WRR) InitRR(nodes []*RRNode) {
	wrr.nodes = nodes
	wrr.n = len(nodes)
	wrr.maxW = wrr.getMaxWeight()
	wrr.gcdW = wrr.getWeightGcd()
	wrr.currPos = -1
}

func (wrr *WRR) Next() *RRNode {
	if wrr.n <= 0 {
		return nil
	}

	pos := wrr.currPos

	for {
		pos = (pos + 1) % wrr.n
		if pos == 0 {
			wrr.currW -= wrr.gcdW
			if wrr.currW <= 0 {
				wrr.currW = wrr.maxW
			}
		}

		if wrr.nodes[pos].Weight >= wrr.currW {
			wrr.currPos = pos
			return wrr.nodes[pos]
		}
	}
}

func (wrr *WRR) getMaxWeight() (maxWeight int) {
	for _, v := range wrr.nodes {
		if v.Weight > maxWeight {
			maxWeight = v.Weight
		}
	}

	return
}

func (wrr *WRR) getWeightGcd() int {
	ret := wrr.nodes[0].Weight

	for _, v := range wrr.nodes {
		ret = wrr.gcd(ret, v.Weight)
	}

	return ret
}

// Get greatest common divisor
func (wrr *WRR) gcd(a, b int) int {
	var c = 0

	for b > 0 {
		c = b
		b = a % b
		a = c
	}

	return a
}
