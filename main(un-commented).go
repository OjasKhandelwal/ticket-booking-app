package main
import (
	"fmt"
	"strings"
	"time"
	"sync"
)
var ConferenceName = "Go Conference" 
const NoOfTickets = 50 
var RemainingTickets = 50
var bookings = make([]UserData , 0)
type UserData struct{ 
	firstname string
	lastname string
	email string
	nooftickets int
}
var wg = sync.WaitGroup{} 
func main() {
	greetusers() 
	UserName, LastName, EMail, UserTicket := getuserinput()
	ValidName, ValidEmail, ValidTics := validateuserinput(UserName, LastName, EMail, int(UserTicket))
	if ValidName && ValidTics && ValidEmail { 
		booktickets(int(UserTicket), UserName, LastName, EMail)
		wg.Add(1)  
		go sendticket(UserTicket, UserName, LastName, EMail) 
		FirstNames := getfirstnames() 
		fmt.Printf("first names of bookings: %v \n", FirstNames)
		if RemainingTickets == 0 {
			fmt.Printf("!!!!!!!!!!!!booking ends!!!!!!!!!!!")
		}
	} else {
		if !ValidName {
			fmt.Println("name is too short!!!!!")
		}
		if !ValidEmail {
			fmt.Println("email should contain @")
		}
		if !ValidTics {
			fmt.Println("no. of tics book is not valid")
		}
	}
	wg.Wait() 
}
func greetusers() { 
	fmt.Printf("welcome to the %v booking application!!!!! \n", ConferenceName)
	fmt.Printf("We have total of %v  tickets and %v tickets are still available \n", NoOfTickets, RemainingTickets)
	fmt.Println("Get your tickets here to attend the event")
}
func getfirstnames() []string {
	FirstNames := []string{}
	for _, booking := range bookings { 
		FirstNames = append(FirstNames, booking.firstname)
	}
	return FirstNames
}
func getuserinput() (string, string, string, uint) {
	var UserName string 
	var UserTicket uint 
	var LastName string
	var EMail string
	fmt.Println("first name:")
	fmt.Scan(&UserName)
	fmt.Println("last name:")
	fmt.Scan(&LastName)
	fmt.Println("email:")
	fmt.Scan(&EMail)
	fmt.Println("how many tickets you want:")
	fmt.Scan(&UserTicket)
	return UserName, LastName, EMail, UserTicket
}
func validateuserinput(UserName string , LastName string , EMail string , UserTicket int ) (bool , bool , bool){
	ValidName := len(UserName)>2 && len(LastName)>2
	ValidEmail := strings.Contains(EMail , "@")
	ValidTics := UserTicket>0 && UserTicket<=RemainingTickets
	return ValidName , ValidEmail , ValidTics
}
func booktickets(UserTicket int, UserName string, LastName string, EMail string) {
	RemainingTickets = RemainingTickets - int(UserTicket)
	var userData = UserData{
		firstname : UserName,
		lastname:LastName,
		email: EMail, 
		nooftickets: UserTicket,
	}
	bookings = append(bookings, userData)
	fmt.Printf("list of bookings : %v \n", bookings)
	fmt.Printf(" Thankyou %v %v for booking %v tickets \n you will soon get conformation on your email address : %v \n", UserName, LastName, UserTicket, EMail)
	fmt.Printf("%v tickets remaining for %v \n", RemainingTickets, ConferenceName)
}
func sendticket(UserTicket uint , FirstName string , LastName string , EMail string) {
	time.Sleep(10* time.Second) 
	var tics = fmt.Sprintf("%v tickets for %v %v", UserTicket , FirstName , LastName)
	fmt.Println("*****************************")
	fmt.Printf("Sending ticket: \n %v \n to email address %v \n", tics , EMail )
	fmt.Println("*****************************")
	wg.Done() 
}




