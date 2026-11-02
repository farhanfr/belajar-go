package main

import "fmt"

func main() {
	// name := "Farhan"
	// age := 25
	// language := "PHP, JavaScript"
	// isLearningGo := true

	// fmt.Println(name)
	// fmt.Println(age)
	// fmt.Println(language)
	// fmt.Println(isLearningGo)

	score := 85
	if score >= 90 {
		fmt.Println("A")
	} else if score >= 80 {
		fmt.Println("B")
	} else if score >= 70 {
		fmt.Println("C")
	} else {
		fmt.Println("D")
	}
}
