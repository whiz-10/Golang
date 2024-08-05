package main

import "fmt"

func main() {
	fmt.Println("maps")

	languages := make(map[string]string)

	languages["js"] = "javascript"
	languages["rb"] = "ruby"
	languages["py"] = "python"

	fmt.Println("list of languages:", languages)
	fmt.Println("js shorts for:", languages["js"])

	delete(languages, "js")
	fmt.Println("list of languages:", languages)

	for key, value := range languages {
		fmt.Printf("for key %v, value is %v\n", key, value)
	}
}
