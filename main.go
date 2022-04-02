package main

import (
	"fmt"
	"main/testPackage"
	"main/testPackage2"
)

func main() {
	fmt.Println("Hello world")
	fmt.Println(testPackage.Greet())
	testPackage2.Mew()

	name := "String variable"

	fmt.Println(name)
}
