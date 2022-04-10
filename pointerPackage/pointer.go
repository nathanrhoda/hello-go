package pointerPackage

import "fmt"

func Input() {
	i := 10
	fmt.Printf("%T %v \n", &i, &i)
	fmt.Printf("%T %v \n", *(&i), *(&i))

	var ptr_i *int
	fmt.Println(ptr_i)

	var ptr_ivar *int = &i
	fmt.Println(ptr_ivar)
	fmt.Println(*ptr_ivar)

	pointerVar := &i
	fmt.Println(pointerVar)
	fmt.Println(*pointerVar)
}
