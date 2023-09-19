package main

import (
	"fmt"
	"goworld/Types"
	util "goworld/Util"
)

func main() {
	usr := Types.User{
		Username: "James Milner",
		Age:      util.GetAge(),
	}
	fmt.Println(usr)
}
