package main

import "fmt"

func main() {
	fmt.Println("welcome")

	var arr [4]string
	arr[0] = "fnb"
	arr[2] = "anj"
	arr[3] = "qwk"

	fmt.Println("string is", arr)
	fmt.Println("string is", len(arr))

	var veglist = [4]string{"eruif", "wiurfd", "mzaed"}
	fmt.Println("string is", veglist)
	fmt.Println("string is", len(veglist))

}
