package deferPackage

import (
	"fmt"
)

func Defer() {
	fmt.Println("Before")
	defer fmt.Println("Deferred")
	fmt.Println("After")
}
