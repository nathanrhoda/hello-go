package methodSetPackage

import (
	"fmt"
)

type Circle struct {
	radius float64
	area   float64
}

func MethodSet() {
	c := Circle{radius: 5}
	c.calcArea()
	fmt.Printf("%+v\n", c)
}

func (c *Circle) calcArea() {
	c.area = 3.14 * c.radius * c.radius

}
