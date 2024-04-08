package main

import (
	"fmt"
	"strings"
)

func orderItem() {

	printMenu()

	var itemNumber uint
	var noOfKgs uint

	for {

		fmt.Println()
		fmt.Println("Enter 0 to exit.")
		fmt.Print("Enter the item no to Order: ")

		fmt.Scan(&itemNumber)
		if itemNumber == 0 {
			break
		}

		var choiceName string
		var itemPrice float64

		for index, item := range menu {

			if index+1 == int(itemNumber) {
				choiceName = item.itemName
				itemPrice = item.itemPrice
				break
			}

		}

		fmt.Printf("How many kgs of  %s do you want: ", choiceName)
		fmt.Scan(&noOfKgs)
		if noOfKgs == 0 {
			continue
		} else {

			for index /*, item */ := range menu {

				if index+1 == int(itemNumber) {
					customerOrder[choiceName] += noOfKgs

					subTotalBill += itemPrice * float64(noOfKgs)
					break

				}

			}
			fmt.Printf("You just ordered %vkgs %v which amounts for ₹%v", noOfKgs, choiceName, itemPrice)

			orderTillNow()

		}

	}

}

func orderTillNow() {

	fmt.Println("\nYou ordered till now")
	fmt.Printf("%s\n", strings.Repeat("-", 35))
	fmt.Printf("%-12s %s\n", "Kgs", "Item")
	fmt.Printf("%s\n", strings.Repeat("-", 35))

	for i := range customerOrder {

		fmt.Printf("%-12v %v\n", customerOrder[i], i)
	}

	fmt.Printf("%s\n", strings.Repeat("-", 32))

}
