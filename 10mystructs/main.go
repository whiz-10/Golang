package main

import "fmt"

func main() {
	fmt.Println("struct in golang")

	// no inheritance in go an no super or parent
	harshit := User{"Harshit", 21, true, 1000.4056}
	fmt.Println(harshit)

	fmt.Printf("harshit details are %v\n", harshit)
	fmt.Printf("harshit's age is %v\n", harshit.Age)

}

type User struct {
	Name   string
	Age    int
	status bool
	salary float64
}
