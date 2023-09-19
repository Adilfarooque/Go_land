package main

import "fmt"

func main() {

	usr := User{
		username: getName(),
		age:      getAge(),
		Address:  getAddress(),
	}
	fmt.Println(usr)
}
