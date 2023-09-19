package main

import (
	"fmt"
	"goworld/Types"
	util "goworld/Util"
)

func main() {
	usr := Types.User{
		Username: util.GetName(),
		Age:      util.GetAge(),
	}
	fmt.Println(usr)
}
