package functionPackage

import (
	"fmt"
)

func Function() {
	fmt.Println(addNumbers(1, 2))
	fmt.Println(multiReturn())
	fmt.Println(multiReturnWithName())
	fmt.Println(Variadic("loop", 1, 2, 3, 4))
	fmt.Println(Variadic("Variadic", 1, 2))
	msg, _ := Variadic("Tame", 1, 2, 3)
	fmt.Println(msg)

	fmt.Println(Recursive(5))
	fmt.Println(AnonymousFunction())

	bb := func(l int, b int) int {
		return l * b
	}(1, 2)

	fmt.Printf("%T \n", bb)
	fmt.Println(bb)
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

func Variadic(looper string, numbers ...int) (string, []int) {
	return looper, numbers
}

func Recursive(n int) int {
	if n == 1 {
		return n
	}

	product := 1
	for i := 1; i <= n; i++ {
		product = product * i
	}
	return product
}

func AnonymousFunction() int {
	x := func(l int, b int) int {
		return l * b
	}

	fmt.Printf("%T \n", x)
	return x(20, 30)
}
