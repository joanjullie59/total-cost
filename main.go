package main

import "fmt"

func main() {
	/*cost1=price of a drink
	cost2=price of a snack
	totalCost=the price of both a drink and a snack
	*/
	var cost1 int           //value in shillings
	var cost2 int           //value in shillings
	var availableAmount int //value in shillings
	totalCost := cost1 + cost2
	//input values
	fmt.Print("Enter value of cost1")
	fmt.Scan(&cost1)
	fmt.Print("Enter value of cost2")
	fmt.Scan(&cost2)
	fmt.Print("Enter value of available amount")
	fmt.Scan(&availableAmount)
	//calculating values
	if totalCost <= availableAmount {
		fmt.Printf("You can afford both the drink and the snack")
	} else {
		fmt.Printf("You cannot afford both the drink and the snack")
	}
}
