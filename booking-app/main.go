package main

import (
	"fmt"
	"time"
	"sync"
)

const tickets int = 150

var eventName = "GOLANG"
var remainTickets uint = 50
var ticketsBooking = make([]UserData, 0)

type UserData struct {
	firstname       string
	lastname        string
	email           string
	numberOfTickets uint
}
var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	userFirstName, userLastName, userEmail, userTickets := getUserInput()

	isValidName, isValidEmail, isValidTicketNumber := ValidateUserInput(userFirstName, userLastName, userEmail, userTickets)

	if isValidName && isValidEmail && isValidTicketNumber {

		bookTicket(userTickets, userFirstName, userLastName, userEmail)

		wg.Add(1)
		go sendTicket(userTickets, userFirstName, userLastName, userEmail)

		userFirstNames := getFirstNames()
		fmt.Printf("The first names of bookings are : %v", userFirstNames)
		fmt.Println()

		if remainTickets == 0 {
			fmt.Println("Our conference is booked out!")
			// break
		}

	} else {
		if !isValidName {
			fmt.Println("First name or Last name you entered is too short")
		}
		if !isValidEmail {
			fmt.Println("Email address you entered doesn't contain @ sign")
		}
		if !isValidTicketNumber {
			fmt.Println("Number of tickets you entered is invalid")
		}
	}

	wg.Wait()
}

func greetUsers() {
	fmt.Println("--------------------------------------------")
	fmt.Printf("Welcome to %v Ease Booking app\n", eventName)
	fmt.Println("--------------------------------------------")
	fmt.Printf("We have total of %v tickets and %v are still available for you.\n", tickets, remainTickets)
	fmt.Println("Get your tickets RIGHT NOW!! to attend")

}

func getFirstNames() []string {
	userFirstNames := []string{}
	for _, booking := range ticketsBooking {
		userFirstNames = append(userFirstNames, booking.firstname)
	}
	return userFirstNames
}

func getUserInput() (string, string, string, uint) {
	// User details input

	var userFirstName string
	var userLastName string
	var userEmail string
	var userTickets uint

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

	return userFirstName, userLastName, userEmail, userTickets
}

func bookTicket(userTickets uint, userFirstName string, userLastName string, userEmail string) {
	remainTickets = remainTickets - userTickets

	// map for user
	var userData = UserData{
		firstname:       userFirstName,
		lastname:        userLastName,
		email:           userEmail,
		numberOfTickets: userTickets,
	}

	ticketsBooking = append(ticketsBooking, userData)
	fmt.Printf("List of bookings %v\n", ticketsBooking)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation mail at %v\n ", userFirstName, userLastName, userTickets, userEmail)
	fmt.Println()

	fmt.Printf("%v tickets remaining for %v", remainTickets, eventName)
	fmt.Println()
}

func sendTicket(userTickets uint, userFirstName string, userLastName string, userEmail string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, userFirstName, userLastName)
	fmt.Println("--------------------------------------------")
	fmt.Printf("Sending ticket : \n %v \n to email address %v\n", ticket, userEmail)
	fmt.Println("--------------------------------------------")
	wg.Done()
}
