package main

import "fmt"

type CR7 struct {
	stamina int
	power   int
	sui     int
}

func (c CR7) KickBall()int {
	return c.stamina + c.power * c.sui
}

type Messi struct{
	stamina int 
	power int
	curve int
}


func main() {

}
