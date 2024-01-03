package main

import "fmt"

func main() {
	// Write your code here
	fmt.Println("Earned amount:")
	fmt.Println("Bubblegum: $202")
	fmt.Println("Toffee: $118")
	fmt.Println("Ice cream: $2250")
	fmt.Println("Milk chocolate: $1680")
	fmt.Println("Doughnut: $1075")
	fmt.Println("Pancake: $80")
	fmt.Println("\nIncome: $5405.0")

	income := 5405
	var staffExpenses int
	var otherExpenses int

	fmt.Println("Staff expenses:")
	fmt.Scan(&staffExpenses)

	fmt.Println("Other expenses:")
	fmt.Scan(&otherExpenses)

	var netIncome = income - (staffExpenses + otherExpenses)
	fmt.Print("\nNet income: $")
	fmt.Print(netIncome)
}
