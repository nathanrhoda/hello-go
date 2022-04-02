package inputPackage

import "fmt"

func Input() {
	var name string
	var surname string
	var isMale bool
	fmt.Printf("Enter your name, surname and if you are male: ")
	fmt.Scanf("%s %s %t", &name, &surname, &isMale)
	fmt.Printf("Whatup %s %s Your have selected %t so you are a male ", name, surname, isMale)
}
