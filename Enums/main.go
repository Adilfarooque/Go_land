package main

import "fmt"

type WeaponType int

func getDamage(weapontype string) int {
	switch weapontype {
	case "Axe":
		return 100
	case "Woodenstick":
		return 10
	case "Knife":
		return 50
	case "Rolling pin":
		return 70
	default:
		panic("Weapon Does not exist..")
	}
}

func main() {
	fmt.Println("Damage Of Weapon ", getDamage("Knife"))
}
