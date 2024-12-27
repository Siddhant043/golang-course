package main

import "fmt"

func main() {
	age := 32

	var agePointer *int
	agePointer = &age
	fmt.Println("Age is:", *agePointer)

	//adultYears := calcAdultYears(age)
	adultYears := calcAdultYears(agePointer)
	fmt.Println("Adult Years:", adultYears)
}

//	func calcAdultYears(age int) int {
//		return age - 18
//	}
func calcAdultYears(age *int) int {
	return *age - 18
}
