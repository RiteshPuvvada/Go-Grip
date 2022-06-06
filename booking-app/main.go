package main

import "fmt"

func main(){

	var eventName = "Nothing"
	const tickets int = 150
	var remainTickets int = 50
	fmt.Println("--------------------------------------------")	
	fmt.Printf("Welcome to %v Ease Booking app\n", eventName)
	fmt.Println("--------------------------------------------")	
	fmt.Printf("We have total of %v tickets and %v are still available for you.\n", tickets, remainTickets)
	fmt.Println("Get your tickets RIGHT NOW!! to attend")

	// User name input

	var userName string = "XYZ"
	var userTickets int = 10

	fmt.Printf("User %v booked %v tickets\n", userName, userTickets)
	
}

