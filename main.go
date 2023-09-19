package main

import "fmt"

func main() {
	str := getString()
	fmt.Println("", str)
	num := getNumber()
	fmt.Println("", num)
	usr := User{
		username: "muhammed anas",
		age:      23,
		Address:  "Poonghadan house",
	}
	fmt.Println(usr)
}
