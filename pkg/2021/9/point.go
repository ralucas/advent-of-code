package day9

type Point struct {
	row   int
	col   int
	value int
}

func NewPoint(row, col, value int) Point {
	return Point{row, col, value}
}

type NodePoint struct {
	point    Point
	children []*NodePoint
}

func NewNodePoint(p Point) *NodePoint {
	return &NodePoint{point: p}
}

func (n *NodePoint) addChild(p Point) {
	np := NewNodePoint(p)

	for _, child := range n.children {
		if child.point == p {
			return
		}
	}

	n.children = append(n.children, np)
}

func (n *NodePoint) BuildBasin(grid *Grid) []Point {
	root := n
	basin := make([]Point, 0)

	tried := make(map[Point]bool)

	stack := NewStack(root)

	for !stack.Empty() {
		np, _ := stack.Pop()

		// if already tried
		if _, ok := tried[np.point]; ok {
			continue
		} else {
			tried[np.point] = true
			basin = append(basin, np.point)
		}

		sps := grid.surroundingPoints(np.point)

		for _, sp := range sps {
			if sp.value < 9 {
				if _, ok := tried[sp]; !ok {
					np.addChild(sp)
				}
			}
		}

		for _, child := range np.children {
			stack.Push(child)
		}

	}

	return basin
}
