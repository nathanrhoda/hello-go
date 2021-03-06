package main

import (
	"fmt"
	"main/castingPackage"

	//"main/inputPackage"
	"main/arrayPackage"
	"main/deferPackage"
	"main/functionPackage"
	"main/interfacePackage"
	"main/methodSetPackage"
	"main/operatorPackage"
	"main/orderFunctionPackage"
	"main/pointerPackage"
	"main/structPackage"
	"main/studentMethodSetPackage"
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
	//inputPackage.InputWithErrors()

	typePackage.TypeFunction()

	castingPackage.Cast()
	castingPackage.StrCast()
	operatorPackage.Operator()
	arrayPackage.ArrayFunction()
	arrayPackage.Slice()
	arrayPackage.Map()
	functionPackage.Function()
	orderFunctionPackage.OrderFunctions()
	deferPackage.Defer()
	pointerPackage.Input()
	structPackage.Struct()
	methodSetPackage.MethodSet()
	studentMethodSetPackage.StudentMethods()
	interfacePackage.InterfaceMethod()
}
