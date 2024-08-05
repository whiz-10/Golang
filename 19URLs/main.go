package main

import (
	"fmt"
	"net/url"
)

const myurl = "https://courses.learncodeonline.in/learn/Go-with-Golang"

func main() {

	fmt.Println("handling URLs")

	fmt.Println(myurl)

	//	parsing
	result, _ := url.Parse(myurl)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery) // all queries are grabbed (nothing for this one)

	qparams := result.Query()
	fmt.Printf("type of query params are : %T\n", qparams)

	fmt.Println(qparams)

	partOfUrl := &url.URL{ //	always pass by reference
		Scheme: "https",
		Host:   "hackerrank.com",
		Path:   "profile/whiz_10",
	}

	anotherUrl := partOfUrl.String()
	fmt.Println(anotherUrl)

}
