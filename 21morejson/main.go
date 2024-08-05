package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("welcome to json")
	//EncodeJson()
	DecodeJson()

}

func EncodeJson() {
	lcoCourses := []course{
		{"ReactJS ", 299, "reactOnline.in", "abc123", []string{"web-dev", "js"}},
		{"Go ", 399, "goOnline.in", "afvc123", []string{"backend", "go"}},
		{"angular ", 299, "angularOnline.in", "fgbc123", nil},
	}

	// package this data as JSON data

	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")
	// marshaling is converting a data structure or object into a JSON string

	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
        "coursename": "ReactJS ",
        "Price": 299,
        "website": "reactOnline.in",
        "tags": ["web-dev","js"]
    }
	`)

	var lcoCourse course

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("Json was valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("JSON WAS NOT VALID")
	}

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("key is %v and value is %v and Type is : %T\n", k, v, v)
	}
}
