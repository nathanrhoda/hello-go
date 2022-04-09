package functionPackage

import (
	"fmt"
)

func Function() {
	fmt.Println(addNumbers(1, 2))
	fmt.Println(multiReturn())
	fmt.Println(multiReturnWithName())
}

func addNumbers(a int, b int) int {
	sum := a + b
	return sum
}

func multiReturn() (int, string) {
	return 1, "woof"
}

func multiReturnWithName() (num int, bark string) {
	bark = "meow"
	num = 2
	return num, bark
}
