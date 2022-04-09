package functionPackage

import (
	"fmt"
)

func Function() {
	fmt.Println(addNumbers(1, 2))
	fmt.Println(multiReturn())
	fmt.Println(multiReturnWithName())
	fmt.Println(Veriadic("loop", 1, 2, 3, 4))
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

func Veriadic(looper string, numbers ...int) (string, []int) {
	return looper, numbers
}
