package roundrobin

// node
type RRNode struct {
	Item   interface{}
	Weight int
}

type RR interface {
	// Init
	InitRR(nodes []*RRNode)

	// Get next
	Next() *RRNode
}
