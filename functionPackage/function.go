package functionPackage

import (
	"fmt"
)

func Function() {
	fmt.Println(addNumbers(1, 2))
	fmt.Println(multiReturn())
}

func addNumbers(a int, b int) int {
	sum := a + b
	return sum
}

func multiReturn() (int, string) {
	return 1, "woof"
}
