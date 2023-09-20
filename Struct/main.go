package main

import "fmt"

type CR7 struct {
	name string
	age  int
	foot string
}

type Messi struct {
	CR7
	curve int
}

func main() {
	m := Messi{
		curve: 100
	}
	fmt.Println(m)
}
