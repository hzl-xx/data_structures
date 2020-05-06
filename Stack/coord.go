package Stack

import "fmt"

type Coord struct {
	x, y int
}

func (c *Coord) InitCoort(x, y int) {
	c.x = x
	c.y = y
}

func (c *Coord) PrintCoort() {
	fmt.Println("("+fmt.Sprintf("%d", c.x)+", "+fmt.Sprintf("%d", c.y)+")")
}
