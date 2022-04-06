package arrayPackage

import (
	"fmt"
)

func ArrayFunction() {
	var list [5]int = [5]int{10, 20, 30}

	fmt.Println(list)

	list1 := [5]int{10, 20, 30, 40, 50}
	fmt.Println(list1)

	list2 := [...]int{10, 20, 30}
	fmt.Println(list2)
	fmt.Println(len(list2))
	fmt.Println(list2[1])

	for i := 0; i < len(list1); i++ {
		fmt.Println("*", list1[i])
	}

	for index, element := range list1 {
		fmt.Println(index, "=>", element)
	}

	multArray := [3][2]int{{2, 4}, {4, 16}, {8, 64}}
	fmt.Println(multArray)
}

func Slice() {
	fmt.Println("Slice")
}

func Map() {
	fmt.Println("Map")
}
