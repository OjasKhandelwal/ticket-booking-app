package main

// int : whole numbers
// uint : positive , whole numbers

import (
	"fmt"
	"strings"
	//"strconv"
	"time"
	"sync"
)

// the variables declared out all the funcs are called package level variables //all func have access to them
// package level variables are created , so that we dont have to define same variable as parameter again n again
var ConferenceName = "Go Conference" //var = variable
//here we don't need to define the datatype as go can detect itself

const NoOfTickets = 50 //const = constant (value cannot be changed)

var RemainingTickets = 50

//RemainingTickets := 50 , this is another syntax of creating a variable in Go
//this syntax can only be used inside a func

//var bookings [50]string  //arrays are list used to store values //in [] we define the no. of values we want in the list

//var bookings = make([]map[string]string , 0)  //[] before map is used to create a list of maps
//here 0 is used to define the initial size of the list

var bookings = make([]UserData , 0)


//var bookings = []string{}//slice is abstraction of array , but it is more flexible & efficient as we dont need to define size of the list


type UserData struct{ //type keyword is used to create a new type , with the name you give to it
	//struct stands for structure , it can hold mixed datatype(like in map we've same datatype throughout , but here we can have mixed datatype)
	firstname string
	lastname string
	email string
	nooftickets int
	
	
}

var wg = sync.WaitGroup{} //waitgroup is used to wait for launched goroutines to finish
//package sync provides basic syncronization functionality 


func main() {
	//bookings := []string{}     alternative syntax to create slice
	greetusers() //calling an explicit function

	//fmt.Printf("conferencename is %T , nooftickets is %T , remainingtickets is %T \n" , ConferenceName , NoOfTickets , RemainingTickets)
	// %T tells the type of datatype

	//for { //go only has for loop

	UserName, LastName, EMail, UserTicket := getuserinput()

	ValidName, ValidEmail, ValidTics := validateuserinput(UserName, LastName, EMail, int(UserTicket))

	if ValidName && ValidTics && ValidEmail { //if , elseif , else

		booktickets(int(UserTicket), UserName, LastName, EMail)
		wg.Add(1) //add - sets the no. of goroutines to wait for 
		go sendticket(UserTicket, UserName, LastName, EMail) //go keyword is used to call a func in a separate goroutine (make it concurrent)

		//as the sendticket func is taking 10secs to execute , it blocks the next booking for 10secs  ,but by making it concurrent it will run in the background and will not block the main func
		

		FirstNames := getfirstnames() //whenever we return a func , we have to store called func in a variable of same name of what we return
		fmt.Printf("first names of bookings: %v \n", FirstNames)

		if RemainingTickets == 0 {
			fmt.Printf("!!!!!!!!!!!!booking ends!!!!!!!!!!!")
			//break
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
 //}
	wg.Wait() //wait - waits for all goroutines that were added to be done , before exit

}

func greetusers() { //func nameoffunc(parameter/argument datatype)
	fmt.Printf("welcome to the %v booking application!!!!! \n", ConferenceName)
	fmt.Printf("We have total of %v  tickets and %v tickets are still available \n", NoOfTickets, RemainingTickets)
	fmt.Println("Get your tickets here to attend the event")

}

func getfirstnames() []string {
	FirstNames := []string{}
	for _, booking := range bookings { //_ is used to ignore a variable that you dont wanna use
		//for index , whatever you wanna call the element of the list := range list
		//this is the index to iterate(loop) through a list , element by element
		//range iterates over elements & provides index n value for each element

		//var names = strings.Fields(booking) //splits the string with space as sperator & return slices with split elements
	//	FirstNames = append(FirstNames, booking["firstName"])
		FirstNames = append(FirstNames, booking.firstname)

	}
	return FirstNames
}

func getuserinput() (string, string, string, uint) {
	var UserName string //string & int are data types
	var UserTicket uint // data types must be defined to prevent assigning the wrong value to wrong var
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
	//Scan is used to get the user input
	// '&' is known as pointer , so basically when we assign a value to a var , the value gets store in memory
	// so the pointer is another var that points to memory address of another variable

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

	//creating an map
	//var userData = make(map[string]string)  //make is used to create an empty map... [datatype of key]datatype of value
	//key n value will have same datatype	
	//userData["firstName"] = UserName
	//userData["lastName"] = LastName
	//userData["email"] = EMail
	//userData["No.oftickets"] = strconv.FormatInt(int64(UserTicket) ,10)

	var userData = UserData{
		firstname : UserName,
		lastname:LastName,
	  email: EMail, 
    nooftickets: UserTicket,
	}
	
	

	
	bookings = append(bookings, userData)
	//we use append to add elements in slices , also we dont need to keep track of index number

	fmt.Printf("list of bookings : %v \n", bookings)

	fmt.Printf(" Thankyou %v %v for booking %v tickets \n you will soon get conformation on your email address : %v \n", UserName, LastName, UserTicket, EMail)
	fmt.Printf("%v tickets remaining for %v \n", RemainingTickets, ConferenceName)

}




func sendticket(UserTicket uint , FirstName string , LastName string , EMail string) {
	time.Sleep(10* time.Second) //sleep stops the current thread execution for a defined duration
	var tics = fmt.Sprintf("%v tickets for %v %v", UserTicket , FirstName , LastName)
	//sprintf helps to put together a string , but instead of printing it out , we can store it in a variable
	fmt.Println("*****************************")
	fmt.Printf("Sending ticket: \n %v \n to email address %v \n", tics , EMail )
	fmt.Println("*****************************")

	wg.Done() //decrements the waitgroup counter by 1 
}

//Local var - defined within a func or within a  block , accessed within that func or block onl
//Package var - defined outside all the funcs , can be used everywhere in the same package
//Global var - defined outside all the funcs & first letter is upper case , can be used everywhere across all the packages


//bookings[0] = UserName + " " + LastName //arrays are index-based (we can add & access elements using index number )
//bookings[5] = "ojas"

//fmt.Printf("the whole slice : %v\n", bookings)
//fmt.Printf("the first value : %v\n", bookings[0])
//fmt.Printf("the type of slice : %T\n", bookings)
//fmt.Printf("the length slice : %v\n", len(bookings))


/*city := "london"
switch city{
case "mumbai":
case "singapore" , "cali":  //we can write like this if 2 cases have same code
case "london":
case "tokyo":
case "karachi":
default:
	fmt.Printf("invalid city")
}*/


