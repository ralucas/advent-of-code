package day10

import (
	"fmt"
	"log"
	"sync"

	"github.com/ralucas/advent-of-code/pkg/utils"
)

type Day struct {
	data []int
}

// TODO: Alter this for actual implementation
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

	_, counts := DistinctArrangements(upd)

	return counts
	//return -1
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

func CountDistinctArrangements(svi []int) int {
	slen := len(svi)

	counts := make([]int, slen-1)

	end := 3
	for i := 0; i < slen-1; i++ {
		if slen-end == i {
			end--
		}
		for k := 1; k <= end; k++ {
			diff := svi[i+k] - svi[i]
			if diff > 3 {
				break
			}
			counts[i]++
		}
	}

	out := 1
	for _, c := range counts {
		if c != 0 {
			out *= c
		}
	}

	return out
}

type Tree struct {
	Root   *Node
	AdjMap map[int][]*Node
	Edges  int
}

type Node struct {
	Val      int
	Children []*Node
}

func (t *Tree) FindAll(val int) []*Node {
	if t.Root.Val == val {
		return []*Node{t.Root}
	}

	//v := val - 3
	//if v < 0 {
	//	v = 0
	//}
	//
	//for {
	//	if _, ok := t.AdjMap[v]; ok {
	//		break
	//	} else {
	//		v++
	//	}
	//}

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

func CountPaths(svi []int) int {
	slen := len(svi)
	counts := make([]int, slen)

	end := 3
	for i, n := range svi {
		if slen-end == i {
			end--
		}
		if i > 0 {
			counts[i] = counts[i-1]
		}
		for k := 1; k <= end; k++ {
			diff := svi[i+k] - n
			if diff <= 3 {
				counts[i] += 1
			}
		}
	}

	return len(counts)
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

// Adapters can only connect to a source 1-3 jolts lower than its rating
// The charging outlet has an effective rating of 0 jolts, so the only
// adapters that could connect to it directly would need to have a joltage
// rating of 1, 2, or 3 jolts.
func DistinctArrangements(svi []int) (Tree, int) {
	slen := len(svi)

	root := Node{Val: 0, Children: make([]*Node, 0)}
	tree := Tree{Root: &root, AdjMap: make(map[int][]*Node)}

	end := 3
	paths := 0
	counts := make([]int, slen)
	counts[0] = 1

	for i := 0; i < slen-1; i++ {
		fmt.Println("on", i)
		if slen-end == i {
			end--
		}

		fmt.Println("prefindall")
		nodes := tree.FindAll(svi[i])
		fmt.Println("postfindall")
		//fmt.Println("postfindall", svi[i], nodes)
		//if len(nodes) > 0 {
		//	fmt.Println(nodes[0].Val)
		//}
		paths = len(nodes)
		//fmt.Println(paths)
		for k := 1; k <= end; k++ {
			diff := svi[i+k] - svi[i]
			if diff <= 3 {
				wg := sync.WaitGroup{}
				for _, n := range nodes {
					wg.Add(1)
					go func(node *Node, ii, kk int) {
						defer wg.Done()
						node.InsertChild(svi[ii+kk])
						//tree.AdjMap[svi[ii]] = append(tree.AdjMap[svi[ii]], nn)
					}(n, i, k)
				}
				wg.Wait()
			}
		}
	}

	//fmt.Println(tree.AdjMap)

	return tree, paths
}
