package main

import (
	"fmt"
	"strings"
)

var shooping = "Go shooping"
var productStock int = 100
var remainingProduct uint = 100
var buy = []string{}

func main() {
	//calling function
	greetUsers(shooping, productStock, remainingProduct)

	for {

		//Calling Function
		firstName, lastName, email, item := getUserInput()

		//calling function
		isValidEmail, isValidName, isValidItem := validateUserInput(firstName, lastName, email, item, remainingProduct)

		if isValidName && isValidEmail && isValidItem {

			//calling function
			order(item, firstName, lastName, email)

			//calling function
			firstNames := getFirstName()
			fmt.Printf("the first name of buy are : %v\n", firstNames)
			fmt.Println("===================================")

			noremainingProduct := remainingProduct == 0
			if noremainingProduct {
				//end program
				fmt.Println("The products is sold out!!!")
				fmt.Println("========================")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("Your first or last name to short!!!")
			}
			if !isValidEmail {
				fmt.Println("Email address is invalid!!!")
			}
			if !isValidItem {
				fmt.Println("Number of items you entered is invalid!!!")
			}
			fmt.Printf("=================================\n")
			continue
		}

	}

	// city := "London"

	// switch city {
	// case "New York":
	// 	//action here
	// case "London":
	// 	//action here
	// case "Jakarta":
	// 	//action here
	// case "Seoul", "Hongkong":
	// 	//action here
	// default:
	// 	fmt.Println("Invalid city selected")
	// }
}

//function greet user
func greetUsers(shop string, productStock int, remainingProduct uint) {
	fmt.Printf("%v now!\n", shop)
	fmt.Println("Welcome to our marketplace application")
	fmt.Printf("We have total of %v product. And %v product still available\n", productStock, remainingProduct)
	fmt.Printf("%v your product here\n", shop)
}

//function get first name
func getFirstName() []string {
	firstNames := []string{}
	for _, buy := range buy {
		var names = strings.Fields(buy)
		firstNames = append(firstNames, names[0]+", ")
	}
	return firstNames
}

//function validate input
func validateUserInput(firstName string, lastName string, email string, item uint, remainingProduct uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidItem := item > 0 && item <= remainingProduct
	return isValidEmail, isValidItem, isValidName
}

//function get user input
func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var item uint

	fmt.Println("====================================")

	// ask user
	fmt.Printf("Enter your first Name : ")
	fmt.Scan(&firstName)

	fmt.Printf("Enter your last Name : ")
	fmt.Scan(&lastName)

	fmt.Printf("Enter your email address : ")
	fmt.Scan(&email)

	fmt.Printf("item of product : ")
	fmt.Scan(&item)

	fmt.Println("===================================")
	return firstName, lastName, email, item
}

//function order
func order(item uint, lastName string, firstName string, email string) {
	remainingProduct = remainingProduct - item
	buy = append(buy, firstName+" "+lastName+", ")

	fmt.Printf("Thank you %v %v for buy %v product. You will receive a confimation email at %v\n", firstName, lastName, item, email)
	fmt.Printf("%v product remaining here attanded\n", remainingProduct)
}
