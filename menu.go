package main

import (
	"fmt"
	"strings"
)

type Menu struct {
	itemNo    uint
	itemName  string
	itemPrice float64
}

var menu = []Menu{

	// itemNo , itemName , itemPrice

	{1, "Apple", 100},
	{2, "Mango", 120},
	{3, "Pineapple", 60},
	{4, "Banana", 40},
	{5, "Jackfruit", 65},
	{6, "Grapes", 70},
	{7, "Orange", 55.50},
	{8, "Watermelon", 75},
	{9, "Papaya", 20},
	{10, "Pomegranate", 80},
	{11, "Fig", 70},
}

func printMenu() {

	// To keep the menu in center 15s
	fmt.Printf("%20s\n", "Fruits List")

	// Repeat the dash by 35 times

	fmt.Printf("%s \n", strings.Repeat("-", 35))

	// Alignment the headers in left , middle and center

	fmt.Printf("%-7s  %6s  %12s\n", "S.No", "Price", "Item Name")

	// Repeat the dash by 35 times by using strings.Repeat
	fmt.Printf("%s \n", strings.Repeat("-", 35))

	// To list all the menu struct useing for loop
	for _, element := range menu {
		fmt.Printf("%-7d  %.2f      %-4s\n", element.itemNo, element.itemPrice, element.itemName)

	}

	// Repeat the dash by 35 times
	fmt.Printf("%s \n", strings.Repeat("-", 35))
	fmt.Println()

}
