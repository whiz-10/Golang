package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the message :")

	//	comma ok syntax || comma error syntax
	input, _ := reader.ReadString('\n')

	fmt.Printf("message is : %T", input)

}
