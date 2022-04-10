package structPackage

import "fmt"

type Circle struct {
	x float64
	y float64
	r float64
}

func Struct() {
	var c Circle
	fmt.Printf("%+v\n", c)

	ci := new(Circle)
	fmt.Printf("%+v\n", ci)

	cir := Circle{
		x: 1,
		y: 2,
		r: 3,
	}

	fmt.Printf("%+v\n", cir)
	fmt.Println(cir)

	fmt.Println(cir.x)
	compare(cir)
}

func compare(circ Circle) {
	circle := circ

	if circle == circ {
		fmt.Println("equal")
	} else {
		fmt.Println("not equal")
	}
	fmt.Println(circle)
}
