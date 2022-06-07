package main

import (
	"fmt"
	"strings"
)

func main() {

	var eventName = "Nothing"
	const tickets int = 150
	var remainTickets uint = 50
	var ticketsBooking []string

	fmt.Println("--------------------------------------------")
	fmt.Printf("Welcome to %v Ease Booking app\n", eventName)
	fmt.Println("--------------------------------------------")
	fmt.Printf("We have total of %v tickets and %v are still available for you.\n", tickets, remainTickets)
	fmt.Println("Get your tickets RIGHT NOW!! to attend")

	// User details input

	var userFirstName string
	var userLastName string
	var userEmail string
	var userTickets uint

	for {
		fmt.Print("Please enter your first name : ")
		fmt.Scan(&userFirstName)
		fmt.Println()

		fmt.Print("Please enter your last name : ")
		fmt.Scan(&userLastName)
		fmt.Println()

		fmt.Print("Please enter your E-mail address : ")
		fmt.Scan(&userEmail)
		fmt.Println()

		fmt.Print("Enter number of tickets need to booked : ")
		fmt.Scan(&userTickets)
		fmt.Println()

		remainTickets = remainTickets - userTickets
		ticketsBooking = append(ticketsBooking, userFirstName+" "+userLastName)

		fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation mail at %v\n ", userFirstName, userLastName, userTickets, userEmail)
		fmt.Println()

		fmt.Printf("%v tickets remaining for %v", remainTickets, eventName)
		fmt.Println()

		userFirstNames := []string{}
		for _, booking := range ticketsBooking{
			var names = strings.Fields(booking)
			userFirstNames = append(userFirstNames,names[0])
		}

		fmt.Printf("The first names of bookings are : %v", userFirstNames)
		fmt.Println()
	}
}
