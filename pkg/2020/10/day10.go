package day10

import (
	"fmt"
	"log"

	"github.com/ralucas/advent-of-code/pkg/utils"
)

type Day struct {
	data []int
}

func (d *Day) PrepareData(filepath string) {
	if filepath == "" {
		log.Fatalf("Missing input file")
	}
	data := utils.ReadFileToArray(filepath, "\n")

	d.data = utils.MapToInt(data)

	return
}

func (d *Day) Part1() interface{} {
	svi := utils.QSort(d.data)
	upd := insertOutletAndDevice(svi)

	counts := CountDiffs(upd)

	return counts[1] * counts[3]
}

func (d *Day) Part2() interface{} {
	svi := utils.QSort(d.data)
	upd := insertOutletAndDevice(svi)

	counts := CountDistinctArrangements(upd)

	return counts
}

func getDevice(svi []int) int {
	return 3 + svi[len(svi)-1]
}

func insertOutletAndDevice(svi []int) []int {
	output := []int{0}
	output = append(output, svi...)
	output = append(output, getDevice(svi))

	return output
}

func CountDiffs(svi []int) map[int]int {
	counts := make(map[int]int)

	slen := len(svi)

	for i := 1; i < slen; i++ {
		diff := svi[i] - svi[i-1]
		counts[diff] += 1
	}

	return counts
}

// CountDistinctArrangements is the speedy and correct implementation
//
// Adapters can only connect to a source 1-3 jolts lower than its rating
// The charging outlet has an effective rating of 0 jolts, so the only
// adapters that could connect to it directly would need to have a joltage
// rating of 1, 2, or 3 jolts.
//
// Runtime O(n)
func CountDistinctArrangements(svi []int) int {
	slen := len(svi)

	pointers := make([]int, slen-1)
	branches := make([]int, slen)

	branches[0] = 1

	end := 3

	for i := 0; i < slen-1; i++ {
		if slen-end == i {
			end--
		}
		for k := 1; k <= end; k++ {
			diff := svi[i+k] - svi[i]
			if diff <= 3 {
				pointers[i]++
			}
		}
		// increment correctly
		// i.e. given this number, how many paths are created from it?
		// get the numbers multiplier, i.e. how many previous branches point to this number
		// then multiply branch[i] by multiplier and add

		// here take the current # of pointers
		// the current number has and add it to
		// any future number that creates a new branch
		p := pointers[i]
		for j := i + 1; j <= p+i; j++ {
			branches[j] += branches[i]
		}
	}

	return branches[len(branches)-1]
}

// Below here is a tree implementation
// Building the tree will correctly count the number
// of paths, however, it is extremely slow on datasets
// much larger than 50 numbers
// Was an interesting study in creation and runtime issues
type Tree struct {
	Root *Node
}

type Node struct {
	Val      int
	Children []*Node
}

func (t *Tree) FindAll(val int) []*Node {
	if t.Root.Val == val {
		return []*Node{t.Root}
	}

	queue := make([]*Node, len(t.Root.Children))
	copy(queue, t.Root.Children)

	found := make([]*Node, 0)

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node.Val == val {
			found = append(found, node)
		}
		queue = append(queue, node.Children...)
	}

	return found
}

func (n *Node) InsertChild(val int) *Node {
	newNode := Node{Val: val, Children: make([]*Node, 0)}
	n.Children = append(n.Children, &newNode)

	return &newNode
}

func (t *Tree) Print() {
	fmt.Printf("%d ", t.Root.Val)
	queue := make([]*Node, len(t.Root.Children))
	copy(queue, t.Root.Children)
	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]
		fmt.Printf("%d ", n.Val)
		queue = append(queue, n.Children...)
	}
	fmt.Print("\n")
}

func BuildTree(svi []int) (Tree, int) {
	slen := len(svi)

	root := Node{Val: 0, Children: make([]*Node, 0)}
	tree := Tree{Root: &root}

	end := 3
	paths := 0

	for i := 0; i < slen-1; i++ {
		if slen-end == i {
			end--
		}

		nodes := tree.FindAll(svi[i])
		paths = len(nodes)
		for k := 1; k <= end; k++ {
			diff := svi[i+k] - svi[i]
			if diff <= 3 {
				for _, n := range nodes {
					n.InsertChild(svi[i+k])
				}
			}
		}
	}

	return tree, paths
}
