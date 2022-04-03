package castingPackage

import (
	"fmt"
	"reflect"
	"strconv"
)

func Cast() {
	var i int = 90
	var f float64 = float64(i)
	fmt.Printf("%.2f\n", f)

	// Losing Precision
	var fNumber float64 = 45.89
	var iNumber int = int(fNumber)
	fmt.Printf("%v\n", iNumber)

	// String Conversion

}

func StrCast() {
	var i int = 42
	var s string = strconv.Itoa(i)
	fmt.Printf("%q of type %v \n", s, reflect.TypeOf(s))

	var strValue string = "200"
	i, err := strconv.Atoi(strValue)

	fmt.Printf("%v, %T \n", i, i)
	fmt.Printf("%v, %T \n", err, err)

}
