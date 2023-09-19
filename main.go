package main

import (
	"fmt"
	"goworld/Types"
)

func main() {
	usr := Types.User{
		Username: getName(),
		Age:      getAge(),
		Address:  getAddress(),
	}
	fmt.Println(usr)
}
