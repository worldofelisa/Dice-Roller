package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	d4Result := RollaDie(4)
	fmt.Println(d4Result)

	d6Result := RollaDie(6)
	fmt.Println(d6Result)

	d8Result := RollaDie(8)
	fmt.Println(d8Result)

	d10Result := RollaDie(10)
	fmt.Println(d10Result)

	d12Result := RollaDie(12)
	fmt.Println(d12Result)

	d20Result := RollaDie(20)
	fmt.Println(d20Result)

	d100Result := RollaDie(100)
	fmt.Println(d100Result)
}

func RollaDie(numberofFaces int) (resultOfRoll int) {
	rand.Seed(time.Now().UnixNano())
	const min = 1
	return rand.Intn(numberofFaces-min) + min
}
