package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("welcome")

	var fruitList = []string{"Apple", "Mango", "Banana"}
	fmt.Println(fruitList)

	fruitList = append(fruitList, "Orange", "Grapes")
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:3])
	fmt.Println(fruitList)

	highScores := make([]int, 4)

	highScores[0] = 675
	highScores[1] = 986
	highScores[2] = 649
	highScores[3] = 754

	fmt.Println(highScores)

	highScores = append(highScores, 325, 458, 345)
	fmt.Println(highScores)

	fmt.Println(sort.IntsAreSorted(highScores))
	sort.Ints(highScores)
	fmt.Println(highScores)

	// remove a value from slices based on index

	var courses = []string{"react", "java", "swift", "php", "ruby"}
	fmt.Println(courses)

	index := 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
