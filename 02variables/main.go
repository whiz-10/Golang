package main

import "fmt"

const Random = "cbhd" //public

func main() {
	var username string = "harshit"
	fmt.Println(username)
	fmt.Printf("Variable is of type : %T \n", username)

	var isPresent bool = true
	fmt.Println(isPresent)
	fmt.Printf("Variable is of type : %T \n", isPresent)

	var newValue int // default values
	fmt.Println(newValue)
	fmt.Printf("Variable is of type : %T \n", newValue)

	var newVal = "google.com" // implicit type
	fmt.Println(newVal)
	fmt.Printf("Variable is of type : %T \n", newVal)

	temp := 1000 // no var style, only allowed inside a function not globally
	fmt.Println(temp)
	fmt.Printf("Variable is of type : %T \n", temp)

	fmt.Println(Random)
	fmt.Printf("Variable is of type : %T \n", Random)

}
