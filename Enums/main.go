package main

import "fmt"

type WeaponType int

const(
	Axe WeaponType = iota
	Woodenstick 
	Knife
	Rolling_pin

)

func getDamage(weapontype WeaponType) int {
	switch weapontype {
	case Axe:
		return 100
	case Woodenstick:
		return 10
	case Knife:
		return 50
	case Rolling_pin:
		return 70
	default:
		panic("Weapon Does not exist..")
	}
}

func main() {
	fmt.Println("Damage Of Weapon ", getDamage(Rolling_pin))
}
