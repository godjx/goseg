package goseg

type Comparable interface {
	CompareTo(interface{}) int
}

type Node struct {
	prev    *Node
	next    *Node
	content interface{}
}

type TermNode struct {
	prev    *TermNode
	next    *TermNode
	content *Term
}

func (node *Node) Prev() *Node {
	return node.prev
}

func (node *Node) Next() *Node {
	return node.next
}

func (node *Node) Content() interface{} {
	return node.content
}

func (tn *TermNode) Prev() *TermNode {
	return tn.prev
}

func (tn *TermNode) Next() *TermNode {
	return tn.next
}

func (tn *TermNode) CompareTo(o interface{}) int {
	other := o.(*TermNode)
	return tn.content.CompareTo(other.content)
}

func (tn *TermNode) Content() *Term {
	return tn.content
}
