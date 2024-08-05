package main

import "fmt"

func main() {

	// days := []string{"sun", "mon", "tue", "wed"}

	// fmt.Println(days)

	// for i := 0; i < len(days); i++ {
	// 	fmt.Println(days[i])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// for index, day := range days {
	// 	fmt.Printf("index is %v and day is %v\n", index, day)
	// }

	val := 1

	for val < 10 {

		if val == 5 {
			fmt.Println("there was 5 once")
			val++
			continue
		}

		if val == 9 {
			goto nix
		}

		fmt.Println("value is:", val)
		val++
	}

nix:
	fmt.Println("now you jump to nix label ")
}
