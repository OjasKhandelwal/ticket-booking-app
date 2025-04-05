package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

const totalTickets = 50
var conferenceName = "Go Conference"
var remainingTickets = totalTickets
var bookings = make([]UserData, 0)
var wg sync.WaitGroup

type UserData struct {
	firstName   string
	lastName    string
	email       string
	numTickets  int
}

func main() {
	greetUsers()

	firstName, lastName, email, numTickets := getUserInput()
	isValidName, isValidEmail, isValidTickets := validateUserInput(firstName, lastName, email, int(numTickets))

	if isValidName && isValidEmail && isValidTickets {
		bookTickets(int(numTickets), firstName, lastName, email)

		wg.Add(1)
		go sendTicket(numTickets, firstName, lastName, email)

		firstNames := getFirstNames()
		fmt.Printf("First names of bookings: %v\n", firstNames)

		if remainingTickets == 0 {
			fmt.Println("All tickets are booked. See you at the event!")
		}
	} else {
		if !isValidName {
			fmt.Println("Name must be at least 2 characters long.")
		}
		if !isValidEmail {
			fmt.Println("Email address must contain '@'.")
		}
		if !isValidTickets {
			fmt.Println("Number of tickets must be between 1 and the available tickets.")
		}
	}

	wg.Wait()
}

func greetUsers() {
	fmt.Printf("Welcome to the %v booking application!\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", totalTickets, remainingTickets)
	fmt.Println("Get your tickets now to attend the event!")
}

func getUserInput() (string, string, string, uint) {
	var firstName, lastName, email string
	var numTickets uint

	fmt.Print("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Print("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Print("Enter your email: ")
	fmt.Scan(&email)

	fmt.Print("Enter number of tickets you want: ")
	fmt.Scan(&numTickets)

	return firstName, lastName, email, numTickets
}

func validateUserInput(firstName, lastName, email string, numTickets int) (bool, bool, bool) {
	isValidName := len(firstName) > 2 && len(lastName) > 2
	isValidEmail := strings.Contains(email, "@")
	isValidTickets := numTickets > 0 && numTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTickets
}

func bookTickets(numTickets int, firstName, lastName, email string) {
	remainingTickets -= numTickets

	user := UserData{
		firstName:  firstName,
		lastName:   lastName,
		email:      email,
		numTickets: numTickets,
	}

	bookings = append(bookings, user)

	fmt.Printf("Thank you %v %v for booking %v tickets.\n", firstName, lastName, numTickets)
	fmt.Printf("Confirmation will be sent to your email: %v\n", email)
	fmt.Printf("%v tickets remaining for %v\n\n", remainingTickets, conferenceName)
}

func getFirstNames() []string {
	var firstNames []string
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func sendTicket(numTickets uint, firstName, lastName, email string) {
	time.Sleep(10 * time.Second)
	ticket := fmt.Sprintf("%v ticket(s) for %v %v", numTickets, firstName, lastName)

	fmt.Println("------------------------------")
	fmt.Printf("Sending ticket:\n%v\nto email: %v\n", ticket, email)
	fmt.Println("------------------------------")

	wg.Done()
}
