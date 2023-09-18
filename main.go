package main

import (
	"fmt"
)

type WeaponTypes int

const(
	Axe WeaponTypes = iota
	Sword
	WoodenStick
	Knife
)

func getDamage(weapontype string)int{
	switch weapontype{
	case "Axe":
		return 100
	case "Sword":
		return 90
	case "WoodenStick":
		return 10
	case "Knife":
		return 40
	default:
		panic("Weapon Does Not Exist..")
	}
}

func main(){
	fmt.Println("Damage Of The Weapon:",getDamage("Knife"))
}