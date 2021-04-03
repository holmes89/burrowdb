package burrowdb

type NodeCreator interface {
	Create(n *Node) error
}
