package day17

import (
	"log"
	"strings"

	"github.com/pkg/errors"
)

type Cube struct {
	x         int
	y         int
	z         int
	state     State
	neighbors []*Cube
}

func NewCube(x, y, z int, state State) (*Cube, error) {
	c := &Cube{x, y, z, state, make([]*Cube, 26)}

	if err := c.Register(); err != nil {
		return nil, errors.Wrap(err, "failed to register cube")
	}

	return c, nil
}

func InitializeCubes(rows, cols int) ([][][]*Cube, error) {
	cubes := make([][][]*Cube, rows)

	for x := 0; x < rows; x++ {
		cubes[x] = make([][]*Cube, cols)

		for y := 0; y < cols; y++ {
			cubes[x][y] = make([]*Cube, 3)

			for z := 0; z < 3; z++ {
				var err error

				cubes[x][y][z], err = NewCube(x, y, z, Inactive)

				if err != nil {
					return nil, err
				}
			}
		}
	}

	return cubes, nil
}

func (c *Cube) Register() error {
	return nil
}

func (c *Cube) GetState() State {
	return c.state
}

func (c *Cube) SetState(s State) {
	c.state = s
}

func (c *Cube) GetNeighbors() []*Cube {
	return c.neighbors
}

func (c *Cube) SetNeighbor(n *Cube) {
	for i, neighbor := range c.GetNeighbors() {
		if neighbor == n {
			c.neighbors[i] = n
		}
	}

	c.neighbors = append(c.neighbors, n)
}

func (c *Cube) RegisterNeighbors(cubes [][][]*Cube) error {
	for i := -1; i < 2; i++ {
		x := c.x + i

		for j := -1; j < 2; j++ {
			y := c.y + j

			for k := -1; k < 2; k++ {
				z := c.z + k

				if i != 0 && j != 0 && k != 0 {
					c.SetNeighbor(cubes[x][y][z])
				}
			}
		}
	}

	return nil
}

func CubeStateToString(zdim int, cubes [][][]*Cube) string {
	var sb strings.Builder
	for _, xc := range cubes {
		for _, yc := range xc {
			c := yc[zdim]
			if _, err := sb.WriteString(c.GetState().String()); err != nil {
				log.Fatal(err)
			}
			if _, err := sb.WriteString(","); err != nil {
				log.Fatal(err)
			}
		}
		sb.WriteString("\n")
	}

	return sb.String()
}
