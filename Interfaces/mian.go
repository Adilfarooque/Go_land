package main

import (
	"fmt"
	"math/rand"
)

type Player interface{
	KickBall()
}

type FootballPlayer struct {
	stamina int
	power   int
}

type CR7 struct {
	stamina int
	power   int
	sui     int
}


func (c CR7) KickBall() {
	shot := c.stamina + c.power*c.sui
	fmt.Println("Cr7 fentasitc acrobatic finish...", shot)
}


type Messi struct {
	stamina int
	power   int
	curve   int
}

func (m Messi) KickBall() {
	shot := m.stamina + m.power*m.curve
	fmt.Println("Messi Curling to the Corner.....have a fabiolous freekick.", shot)
}


func (f FootballPlayer) KickBall() {
	shot := f.stamina + f.power
	fmt.Println("I'm kicking the ball", shot)
}


func main() {
	team := make([]Player, 11)
	for i := 0; i < len(team); i++ {
		team[i] = FootballPlayer{
			stamina: rand.Intn(10),
			power:   rand.Intn(10),
		}
	}
	team[len(team)-2]=Messi{
		stamina: 8,
		power: 9,
		curve: 10,
	}
	team[len(team)-1]=CR7{
		stamina: 10,
		power : 10,
		sui :10,
	}
	for i := 0; i < len(team); i++ {
		team[i].KickBall()
	}
}
