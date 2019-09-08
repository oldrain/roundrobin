package roundrobin

type swNode struct {
	Node *RRNode

	EffectiveW int
	CurrW      int
}

/**
 * Smooth Weighted Round Robin
 * https://github.com/nginx/nginx/commit/52327e0627f49dbda1e8db695e63a4b0af4448b1
 */
type SW struct {
	nodes []*swNode

	n       int
	currPos int
	totalW  int
}

func (sw *SW) InitRR(nodes []*RRNode) {
	var swNodes []*swNode

	for _, node := range nodes {
		swN := new(swNode)
		swN.Node = node
		swN.EffectiveW = node.Weight
		swN.CurrW = node.Weight

		swNodes = append(swNodes, swN)
	}

	sw.n = len(nodes)
	sw.nodes = swNodes
	sw.currPos = 0
	sw.totalW = sw.calTotalWeight()
}

func (sw *SW) Next() *RRNode {
	if sw.n <= 0 {
		return nil
	}

	pos := sw.findMaxWeightPos()
	currNode := sw.nodes[pos]

	sw.downCurrNodeWeight(pos)
	sw.raiseAllNodeWeight()

	return currNode.Node
}

func (sw *SW) raiseAllNodeWeight() {
	for k, v := range sw.nodes {
		sw.nodes[k].CurrW += v.EffectiveW
	}
}

func (sw *SW) downCurrNodeWeight(currPos int) {
	sw.nodes[currPos].CurrW -= sw.totalW
}

func (sw *SW) findMaxWeightPos() int {
	var pos = 0
	var maxW = 0
	for k, v := range sw.nodes {
		if v.CurrW > maxW {
			maxW = v.CurrW
			pos = k
		}
	}

	return pos
}

func (sw *SW) calTotalWeight() int {
	var count = 0

	for _, swNode := range sw.nodes {
		count += swNode.EffectiveW
	}

	return count
}
