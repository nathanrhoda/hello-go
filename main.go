package main

import (
	"fmt"
	"main/inputPackage"
	"main/testPackage"
	"main/testPackage2"
	"main/typePackage"
	"main/varPackage"
)

func main() {
	fmt.Println("Hello world")
	fmt.Println(testPackage.Greet())
	testPackage2.Mew()
	varPackage.Woof()
	//inputPackage.Input()
	inputPackage.InputWithErrors()

	typePackage.TypeFunction()
}
