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
}
