package main

import "fmt"

const Logintoken string = "sdgfdgdsg" // Public Variable

func main() {
	var username string = "Mann"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)

	var smallVal uint16 = 65535
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var val uint16
	fmt.Println(val)
	fmt.Printf("Variable is of type: %T \n", val)

	var newUsername string
	fmt.Println(newUsername)
	fmt.Printf("Variable is of type: %T \n", newUsername)

	//implicit type
	var website = "www.google.com"
	fmt.Println(website)

	//No var used // Walrus operator // Can only be used inside function and not globally
	newNum := 64233
	fmt.Println(newNum)

	fmt.Println(Logintoken)
}
