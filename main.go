package main

import (
	"fmt"
	"sync"
	"time"
)

type UserData struct {
	firstName   string
	lastName    string
	email       string
	userTickets uint
}

var conferenceName = "Go Conference"
var remainingTickets uint = ConferenceTickets
var bookings = make([]UserData, 0)

var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	//for {

	firstName, lastName, email, userTickets := captureUserInput()

	isValidName, isValidEmail, isValidTicketCount := ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

	if isValidName && isValidEmail && isValidTicketCount {

		bookTicket(userTickets, firstName, lastName, email)

		wg.Add(1)
		go sendTicket(userTickets, firstName, lastName, email)

		//print firstnames
		firstNames := getFirstNames()

		fmt.Printf("The firstnames of bookings are: %v\n", firstNames)

		if remainingTickets == 0 {
			fmt.Printf("We are sold out for this year. come back nexxt year\n")
			//break
		}
	} else {

		if !isValidEmail {
			fmt.Printf("Email you entered is not valid - %v\n", email)
		}

		if !isValidName {
			fmt.Printf("Firstname or lastname you entered is not valid - Firstname =  %v , LastName= %v\n", firstName, lastName)
		}

		if !isValidTicketCount {
			fmt.Printf("Ticket count you requested  %v is not valid. Remaining tickets we have is %v \n", userTickets, remainingTickets)
		}

		if !isValidEmail {
			fmt.Printf("Email you entered is not valid - %v\n", email)
		}

		fmt.Printf("your input data is invalid, please try again\n")
	}
	//	}
	wg.Wait()
}

func greetUsers() {

	fmt.Printf("Welcome to - %v booking application\n", conferenceName)

	fmt.Printf("We have total of %v tickets and %v are still available\n", ConferenceTickets, remainingTickets)
	fmt.Printf("Get your tickets here to attend\n")
}

func getFirstNames() []string {
	firstNames := []string{}

	for _, booking := range bookings {

		firstNames = append(firstNames, booking.firstName)
	}

	return firstNames
}

func captureUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	//ask user for their name
	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email:")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets you need:")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {

	//create a map for a user

	var userData = UserData{
		firstName:   firstName,
		lastName:    lastName,
		email:       email,
		userTickets: userTickets,
	}
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, userData)

	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets.  you will receive confirmation email at %v \n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v \n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {

	time.Sleep(10 * time.Second)

	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("********************************")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("********************************")

	wg.Done()
}
