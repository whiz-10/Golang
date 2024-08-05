package main

import "fmt"

func main() {

	if 9%2 == 0 {
		fmt.Println("even")
	} else {
		fmt.Println("odd")
	}

	if num := 3; num < 10 {
		fmt.Println("num is less than 10")
	} else {
		fmt.Println("num is not less than 10")
	}
}
