package main

import "fmt"

func greet(customerName string) {

	fmt.Printf("👋 Hi dear  %s\n", customerName)
	fmt.Println("Welcome to fruit center")
	fmt.Println()

}

func sayTata(customerName string) {
	fmt.Println()
	fmt.Printf("%55s %s\n", "Thank you for visiting the fruit center", customerName)
	fmt.Printf("%51s\n", "We hope to see you again!")
	fmt.Printf("%51s\n", "Have a nice day!😀 ")

}
