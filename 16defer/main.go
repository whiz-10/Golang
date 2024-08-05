package main

import "fmt"

func main() {
	defer fmt.Println("world2")
	defer fmt.Println("world1")
	fmt.Println("hello")
	mydefer()

}
func mydefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}
