package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	var x int = rand.Intn(7)
	var dirX [2]string
	dirX[0] = "North"
	dirX[1] = "South"
	var cardX = rand.Intn(2)

	if x == 0 {
		fmt.Println("Stay Still")
	} else {
		fmt.Println("Go", (x * 5), "feet", dirX[cardX])
	}
	var y int = rand.Intn(7)
	var dirY [2]string
	dirY[0] = "East"
	dirY[1] = "West"
	var cardY = rand.Intn(2)
	if y == 0 {
		fmt.Println("Stay Still")
	} else {

		fmt.Println("Then go", (y * 5), "feet", dirY[cardY])
	}
	var saveX int = rand.Intn(6)
	var saveType [6]string
	saveType[0] = "strength"
	saveType[1] = "dexterity"
	saveType[2] = "constitution"
	saveType[3] = "intelligence"
	saveType[4] = "wisdom"
	saveType[5] = "charisma"

	var dmgX int = rand.Intn(13)
	var dmgType [13]string
	dmgType[0] = "acid damage"
	dmgType[1] = "bludgeoning damage"
	dmgType[2] = "cold damage"
	dmgType[3] = "fire damage"
	dmgType[4] = "force damage"
	dmgType[5] = "lightning damage"
	dmgType[6] = "necrotic damage"
	dmgType[7] = "piercing damage"
	dmgType[8] = "poison damage"
	dmgType[9] = "psychic damage"
	dmgType[10] = "radiant damage"
	dmgType[11] = "slashing damage"
	dmgType[12] = "thunder damage"
	fmt.Println("Or make a", saveType[saveX], "saving throw against", dmgType[dmgX])
}
