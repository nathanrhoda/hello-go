package functionPackage

import (
	"fmt"
)

func Function() {
	fmt.Println(addNumbers(1, 2))
}

func addNumbers(a int, b int) int {
	sum := a + b
	return sum
}