package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	inputArgs := os.Args
	listOfDice := []string{}

	for x := 1; x < len(inputArgs); x++ {
		//we start at one to ignore the program name
		listOfDice = append(listOfDice, inputArgs[x])
	}

	for x := 0; x < len(listOfDice); x++ {
		currentDie := listOfDice[x]
		indexOfD := strings.Index(currentDie, "d")
		numberOfDice, err := strconv.Atoi(currentDie[0:indexOfD])
		if err != nil {
			fmt.Println("You need to add a number before the number of faces to say how many dice you want to roll. Assumed dice number is 1.")
			fmt.Println("For example: roll 2 10")
			numberOfDice = 1
		}
		numberOfFaces, err := strconv.Atoi(currentDie[indexOfD+1 : len(currentDie)])
		if err != nil {
			fmt.Println("You need to add the number of faces on the die. Assumed dice face valye is 6.")
			numberOfFaces = 6
		}
		for i := 0; i < numberOfDice; i++ {
			result := RollaDie(numberOfFaces)
			fmt.Printf("I rolled a d%d and it was a %d\n", numberOfFaces, result)
		}
	}
}

func RollaDie(numberofFaces int) (resultOfRoll int) {
	rand.Seed(time.Now().UnixNano())
	const min = 1
	return rand.Intn(numberofFaces-min) + min
}

// TODO create an extension for the rollies game, or maybe a sub code
