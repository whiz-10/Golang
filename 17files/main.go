package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("files")
	content := "this need to go in a file"

	file, err := os.Create("./GOLANG.txt")

	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}

	fmt.Println("length is :", length)
	defer file.Close()
	readFile("./GOLANG.txt")
}

func readFile(filename string) {
	databyte, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fmt.Println("text data inside the file is :", string(databyte))
}
