package main

import "fmt"

func userInput() {
	var fName string
	fmt.Printf("enter your first name")
	fmt.Scan(&fName)
	var lName string
	fmt.Printf("enter your last name")
	fmt.Scan(&lName)
	fmt.Print("your full name" + fName + " " + lName)
}
