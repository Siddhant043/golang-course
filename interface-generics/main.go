package main

import "fmt"

func main() {
	a := "fruit: "
	b := "mango"

	add(a, b)
}

func add(a, b interface{}) {
	aInt, aIsInt := a.(int)
	bInt, bIsInt := b.(int)

	if aIsInt && bIsInt {
		fmt.Println(aInt + bInt)
	}

	aFloat, aIsFloat := a.(float64)
	bFloat, bIsFloat := b.(float64)

	if aIsFloat && bIsFloat {
		fmt.Println(aFloat + bFloat)
	}

	aString, aIsString := a.(string)
	bString, bIsString := b.(string)

	if aIsString && bIsString {
		fmt.Println(aString + bString)
	}

}
