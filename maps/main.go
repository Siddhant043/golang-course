package main

import "fmt"

func main() {
	websites := map[string]string{
		"Google": "www.google.com",
		"Amazon": "www.amazon.in",
	}
	fmt.Println(websites)

	fmt.Println(websites["Google"])

	websites["LinkedIn"] = "www.linkedin.com"
	fmt.Println(websites)

	delete(websites, "Google")
	fmt.Println(websites)

	usingMake()
}

func usingMake() {
	foodRatings := make(map[string]float64, 3)

	foodRatings["Samosa"] = 4.3
	foodRatings["Khasta"] = 4.6
	foodRatings["Kulcha"] = 4.4

	fmt.Println(foodRatings)
}
