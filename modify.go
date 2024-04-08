package main

import (
	"fmt"
)

func modifyOrder() {

	for {

		var isOrderOk string

		fmt.Println("\nDo you want to change the order?[Y/N]")
		fmt.Scan(&isOrderOk)

		if isOrderOk != "y" {
			return

		}

		var modifyType uint
		var serialNumber uint

		fmt.Println("Please enter the respective number to proceed further")
		fmt.Println("Press '1' to update the item quantity")
		fmt.Println("Press  '2' to delete the item from the list")
		fmt.Println("Press '3' to to add items in the order list")
		fmt.Scan(&modifyType)

		switch modifyType {
		case 1:
			printMenu()

			fmt.Println("Kinldy type the serial number to update the item")
			fmt.Scan(&serialNumber)
			updateQuantity(serialNumber)

		case 2:

			printMenu()
			fmt.Println("Please enter the S.No. of the item to be deleted: ")
			fmt.Scan(&serialNumber)
			delFromOrder(serialNumber)

		case 3:
			insertIntoOrder()

		default:
			return

		}

		displayGeneratingBill()
		generateBill()

	}

}

func updateQuantity(serialNumber uint) {

	var newQuantity uint

	for _, targetItem := range menu {

		if serialNumber == targetItem.itemNo {

			itemName := targetItem.itemName
			oldQuantity := customerOrder[itemName]

			fmt.Printf("Current qunatity of %v is %v \n", itemName, oldQuantity)
			fmt.Printf("Now, enter the updated quantity of %v to be ordered: \n", itemName)
			fmt.Scan(&newQuantity)

			if newQuantity == 0 {
				delFromOrder(serialNumber)
				return
			}

			fmt.Printf("")

			customerOrder[targetItem.itemName] = newQuantity
			fmt.Printf("Updated the quantity of %v from %v to %v.\n", itemName, oldQuantity, newQuantity)

			subTotalBill -= float64(oldQuantity) * float64(targetItem.itemPrice)
			subTotalBill += float64(oldQuantity) * float64(targetItem.itemPrice)
			break

		}

	}

}

func delFromOrder(serialNumber uint) {

	for _, tergetItem := range menu {

		if serialNumber == tergetItem.itemNo {
			itemName := tergetItem.itemName

			subTotalBill -= float64(customerOrder[itemName]) * float64(tergetItem.itemPrice)

			delete(customerOrder, itemName)
			fmt.Printf("%v removed from the order list.\n", itemName)
			break

		}

	}

}

func insertIntoOrder() {
	orderItem()

}
