package day15

type Node struct {
	row, col  int
	neighbors map[*Node]int
}

func NewNode(row, col int) *Node {
	return &Node{
		row:       row,
		col:       col,
		neighbors: make(map[*Node]int),
	}
}

func (n *Node) SetNeighbor(node *Node, weight int) {
	n.neighbors[node] = weight
}

func (n *Node) Neighbors() map[*Node]int {
	return n.neighbors
}

func (n *Node) Row() int {
	return n.row
}

func (n *Node) Col() int {
	return n.col
}
