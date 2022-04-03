package inputPackage

import "fmt"

func Input() {
	var name string
	var surname string
	var isMale bool
	fmt.Printf("Enter your name, surname and if you are male: ")
	fmt.Scanf("%s %s %t", &name, &surname, &isMale)
	fmt.Printf("Whatup %s %s Your have selected %t so you are a male \n", name, surname, isMale)
}

func InputWithErrors() {
	var a string
	var b int

	fmt.Print("Enter a string and a number: ")

	count, err := fmt.Scanf("%s %d", &a, &b)

	fmt.Println("count : ", count)
	fmt.Println("Error : ", err)
	fmt.Println("a: : ", a)
	fmt.Println("b: : ", b)
}
