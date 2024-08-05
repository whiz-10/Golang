package main

import "fmt"

func main() {
	greeter()

	result := adder(984, 95)
	fmt.Println("result is:", result)

	proRes, msg := proadder(4, 96, 1564, 56)
	fmt.Printf("resul is %v and msg is %v:", proRes, msg)

}
func adder(val1 int, val2 int) int {
	return val1 + val2
}
func proadder(values ...int) (int, string) {
	total := 0
	for _, val := range values {
		total += val
	}
	return total, "hi what's up"
}
func greeter() {
	fmt.Println("Namastey!!")
}
