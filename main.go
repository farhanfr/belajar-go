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

	// score := 85
	// if score >= 90 {
	// 	fmt.Println("A")
	// } else if score >= 80 {
	// 	fmt.Println("B")
	// } else if score >= 70 {
	// 	fmt.Println("C")
	// } else {
	// 	fmt.Println("D")
	// }

	transactions := [5]int{
		100000,
		200000,
		150000,
		300000,
		50000,
	}

	total := 0

	// for i := 0; i < len(transactions); i++ {
	// 	total += transactions[i]
	// 	fmt.Println("Transaction:", transactions[i])
	// }

	for _, transaction := range transactions {
		total += transaction
		fmt.Println("Transaction:", transaction)
	}

	fmt.Println("Total Transaction: ", total)

}
