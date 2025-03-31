// start by declaring the main package
package main

// import "fmt" package

import "fmt"

// declaring the main function
func main() {
	// var keyword is reserved in go to decalre variables
	//syntax: var variable_name data_type = value

	var mango string = "Mango" 
	
	fmt.Println(mango)

	//shorthand declaration

	age := 54
	city := "Nairobi"

	fmt.Println("My age is", age)
	fmt.Println("My city is", city)

	//declaring multiple variables in a single line

	 apples, oranges  := 10, 20

	fmt.Println("I have ", apples, "apples and", oranges, "oranges.")
}