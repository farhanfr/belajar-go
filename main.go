package main

import "fmt"

type Transaction struct {
	ID     int
	Title  string
	Amount int
}

func calculateTotal(transactions []Transaction) int {
	total := 0
	for _, transaction := range transactions {
		total += transaction.Amount
	}
	return total
}

func findTransaction(transactions []Transaction, id int) (Transaction, bool) {
	for _, transaction := range transactions {
		if transaction.ID == id {
			return transaction, true
		}
	}
	return Transaction{}, false
}

func main() {

	transactions := []Transaction{
		{ID: 1, Title: "Beli Baju", Amount: 100000},
		{ID: 2, Title: "Beli Celana", Amount: 200000},
		{ID: 3, Title: "Beli Sepatu", Amount: 150000},
		{ID: 4, Title: "Beli Jaket", Amount: 300000},
		{ID: 5, Title: "Beli Topi", Amount: 50000},
	}

	total := calculateTotal(transactions)
	fmt.Println("Total Transaction: ", total)

	transaction, found := findTransaction(transactions, 3)
	if found {
		fmt.Println("ID: ", transaction.ID, "Title: ", transaction.Title, "Amount: ", transaction.Amount)
	} else {
		fmt.Println("Transaction not found")
	}

}

// package main

// import "fmt"

// func main() {
// 	// name := "Farhan"
// 	// age := 25
// 	// language := "PHP, JavaScript"
// 	// isLearningGo := true

// 	// fmt.Println(name)
// 	// fmt.Println(age)
// 	// fmt.Println(language)
// 	// fmt.Println(isLearningGo)

// 	// score := 85
// 	// if score >= 90 {
// 	// 	fmt.Println("A")
// 	// } else if score >= 80 {
// 	// 	fmt.Println("B")
// 	// } else if score >= 70 {
// 	// 	fmt.Println("C")
// 	// } else {
// 	// 	fmt.Println("D")
// 	// }

// 	// transactions := [5]int{
// 	// 	100000,
// 	// 	200000,
// 	// 	150000,
// 	// 	300000,
// 	// 	50000,
// 	// }

// 	// total := 0

// 	// for i := 0; i < len(transactions); i++ {
// 	// 	total += transactions[i]
// 	// 	fmt.Println("Transaction:", transactions[i])
// 	// }

// 	// for _, transaction := range transactions {
// 	// 	total += transaction
// 	// 	fmt.Println("Transaction:", transaction)
// 	// }

// 	// fmt.Println("Total Transaction: ", total)

// 	type Transaction struct {
// 		ID     int
// 		Title  string
// 		Amount int
// 	}

// 	total := 0

// 	transactions := []Transaction{
// 		{ID: 1, Title: "Beli Baju", Amount: 100000},
// 		{ID: 2, Title: "Beli Celana", Amount: 200000},
// 		{ID: 3, Title: "Beli Sepatu", Amount: 150000},
// 		{ID: 4, Title: "Beli Jaket", Amount: 300000},
// 		{ID: 5, Title: "Beli Topi", Amount: 50000},
// 	}

// 	for _, transaction := range transactions {
// 		total += transaction.Amount
// 		fmt.Println("ID: ", transaction.ID, "Title: ", transaction.Title, "Amount: ", transaction.Amount)
// 	}

// 	fmt.Println("Total Transaction: ", total)

// }
