package varPackage

import "fmt"

func Woof() {

	var n string = "Boom"
	fmt.Println("String variable: ", n)

	var i int = 100
	fmt.Printf("The word is %v and the Integer is %d \n", n, i)

	var f float64 = 0.5
	fmt.Printf("Float: %v \n", f)

	var b bool = false
	fmt.Printf("Bool: %v \n", b)
}
