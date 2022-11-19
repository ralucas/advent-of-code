package day17

import (
	"log"

	fileutil "github.com/ralucas/advent-of-code/pkg/util/file"
)

type Day struct {
	Cubes [][][]*Cube
}

// TODO: Alter this for actual implementation
func (d *Day) PrepareData(filepath string) {
	if filepath == "" {
		log.Fatalf("Missing input file")
	}
	data := fileutil.ReadFileToArray(filepath, "\n")

	var err error
	d.Cubes, err = InitializeCubes(len(data), len(data[0]))
	if err != nil {
		log.Fatalf("failed to initialize cubes %v", err)
	}

	// Set the state for the 2D based on the input
	for i, xc := range d.Cubes {
		row := data[i]
		for j, yc := range xc {
			sState := string(row[j])
			cube := yc[0]
			cube.SetState(NewState(sState))
		}
	}

	for _, xc := range d.Cubes {
		for _, yc := range xc {
			for _, zcube := range yc {
				if zcube.RegisterNeighbors(d.Cubes) != nil {
					log.Fatalf("failed to register neighbors %+v\n", zcube)
				}
			}
		}
	}

}

func (d *Day) Part1() interface{} {
	return -1
}

func (d *Day) Part2() interface{} {
	return -1
}

// During a cycle, all cubes simultaneously change their state according to the following rules:
//  1. If a cube is active and exactly 2 or 3 of its neighbors are also active, the cube remains active.
//     Otherwise, the cube becomes inactive.
//  2. If a cube is inactive but exactly 3 of its neighbors are active, the cube becomes active.
//     Otherwise, the cube remains inactive.
func ApplyRules(cube *Cube, cubeState [][][]Cube) *Cube {
	activeNeighborCount := 0

	for _, c := range cube.GetNeighbors() {
		neighbor := cubeState[c.x][c.y][c.z]

		if neighbor.GetState() == Active {
			activeNeighborCount += 1
		}

		switch cube.GetState() {
		case Active:
			if activeNeighborCount == 2 || activeNeighborCount == 3 {
				cube.SetState(Inactive)
			} else {
				cube.SetState(Active)
			}

		case Inactive:
			if activeNeighborCount == 3 {
				cube.SetState(Active)
			} else {
				cube.SetState(Inactive)
			}
		}
	}

	return cube
}

func RunCycle(cubes [][][]*Cube) [][][]*Cube {
	cpCubes := make([][][]Cube, len(cubes))
	for x := range cubes {
		for y := range cubes[x] {
			for z := range cubes[x][y] {
				cpCubes[x][y][z] = *cubes[x][y][z]
			}
		}
	}

	for x := range cubes {
		for y := range cubes[x] {
			for z := range cubes[x][y] {
				ApplyRules(cubes[x][y][z], cpCubes)
			}
		}
	}

	return cubes
}
