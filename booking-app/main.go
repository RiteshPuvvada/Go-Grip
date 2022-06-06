package main

import "fmt"

func main(){

	var eventName = "Nothing"
	const tickets int = 150
	var remainTickets uint = 50
	fmt.Println("--------------------------------------------")	
	fmt.Printf("Welcome to %v Ease Booking app\n", eventName)
	fmt.Println("--------------------------------------------")	
	fmt.Printf("We have total of %v tickets and %v are still available for you.\n", tickets, remainTickets)
	fmt.Println("Get your tickets RIGHT NOW!! to attend")

	// User details input

	var userFirstName string
	var userLasttName string
	var userEmail string
	var userTickets uint

	fmt.Print("Please enter your first name : ")
	fmt.Scan(&userFirstName)
	fmt.Println()
	
	fmt.Print("Please enter your last name : ")
	fmt.Scan(&userLasttName)
	fmt.Println()
	
	fmt.Print("Please enter your E-mail address : ")
	fmt.Scan(&userEmail)
	fmt.Println()
	
	fmt.Print("Enter number of tickets need to booked : ")
	fmt.Scan(&userTickets)
	fmt.Println()
	
	fmt.Printf("Thank you %v %v for booking %v tickets. You will recive a confirmation mail at %v\n ", userFirstName, userLasttName, userTickets, userEmail)

}

