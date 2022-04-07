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
	arraySlice := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sliceName := arraySlice[1:8]
	fmt.Println(sliceName)

	slicedSlice := sliceName[0:2]
	fmt.Println(slicedSlice)

	slicery := make([]int, 5, 10)
	fmt.Println(slicery)
	fmt.Println("Length: ", len(slicery))
	fmt.Println("Capacity: ", cap(slicery))
	//copy / make
}

func Map() {
	codes := map[string]string{"en": "English", "fr": "French"}
	fmt.Println(codes)

	codes["en"] = "Engels"
	fmt.Println(codes)
	fmt.Println("Length: ", len(codes))
	fmt.Println(codes["en"])
	fmt.Println(codes["fr"])
	value, found := codes["fr"]
	fmt.Println(value, found)
	value1, found1 := codes["hh"]
	fmt.Println(value1, found1)
	delete(codes, "en")
	fmt.Println(codes)
}
