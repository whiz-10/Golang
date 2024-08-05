package main

import "fmt"

func main() {
	fmt.Println("struct in golang")

	// no inheritance in go an no super or parent
	harshit := User{"Harshit", 21, true, 1000.4056}
	fmt.Println(harshit)

	fmt.Printf("harshit details are %v\n", harshit)
	fmt.Printf("harshit's age is %v\n", harshit.Age)
	harshit.GetStatus()
	harshit.NewSalary()

}

type User struct {
	Name   string
	Age    int
	Status bool
	Salary float64
}

func (u User) GetStatus() {
	fmt.Println("is user active:", u.Status)
}
func (u User) NewSalary() { // only copy is passed in this case
	u.Salary = 59629.596
	fmt.Println("new salary is :", u.Salary)
}
