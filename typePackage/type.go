package typePackage

import (
	"fmt"
	"reflect"
)

func TypeFunction() {
	var grades int = 42
	var message string = "hello world"

	fmt.Printf("variable grades = %v is of type %T \n", grades, grades)
	fmt.Printf("variable message = %v is of type %T \n", message, message)

	fmt.Printf("Type: %v \n", reflect.TypeOf(grades))
	fmt.Printf("Type: %v \n", reflect.TypeOf(message))
}
