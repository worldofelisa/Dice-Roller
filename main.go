package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	inputArguments := os.Args
	numberofDice, err := strconv.Atoi(inputArguments[2])
	if err != nil {
		fmt.Println("You need to add a number before the number of faces to say how many dice you want to roll.")
		fmt.Println("For example: roll 2 10")
	}

	numberofFaces, err := strconv.Atoi(inputArguments[3])
	if err != nil {
		fmt.Println("You need to add the number of faces on the die.")
		log.Fatalln(err)
	}

	for x := 0; x < numberofDice; x++ {
		result := RollaDie(numberofFaces)
		fmt.Println(result)
	}
}

func RollaDie(numberofFaces int) (resultOfRoll int) {
	rand.Seed(time.Now().UnixNano())
	const min = 1
	return rand.Intn(numberofFaces-min) + min
}

// TODO create an extension for the rollies game, or maybe a sub code
