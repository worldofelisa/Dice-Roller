package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type Die struct {
	NumberOfFaces int
}

func (d *Die) Roll() (resultOfRoll int) {
	rand.Seed(time.Now().UnixNano())
	const min = 1
	return rand.Intn(d.NumberOfFaces-min) + min
}

func ParseArgs(inputArgs []string) []string {
	listOfArgs := []string{}
	for x := 1; x < len(inputArgs); x++ {
		//we start at one to ignore the program name
		listOfArgs = append(listOfArgs, inputArgs[x])
	}
	return listOfArgs
}

func SliceString(inputString string, start, end, defaultValueInputted int, errorMessage string) int {
	stringSegment, err := strconv.Atoi(inputString[start:end])
	if err != nil {
		fmt.Println(errorMessage)
		return defaultValueInputted
	}
	return stringSegment
}

func DieSetUp(inputArgs []string) []Die {
	parsedListOfArgs := ParseArgs(inputArgs)
	listOfDice := []Die{}
	for x := 0; x < len(parsedListOfArgs); x++ {
		currentDie := parsedListOfArgs[x]
		indexOfD := strings.Index(currentDie, "d")
		numberOfDice := SliceString(currentDie, 0, indexOfD, 1, "Must add a number of dice before d to state how many dice you want to roll. Assumed number of dice is 1.")
		numberOfFaces := SliceString(currentDie, indexOfD+1, len(currentDie), 6, "You need to add the type of dice you want to roll based on it face value. Assumed dice face value is 6.")

		for i := 0; i < numberOfDice; i++ {
			newDie := Die{
				NumberOfFaces: numberOfFaces,
			}
			listOfDice = append(listOfDice, newDie)
		}
	}
	return listOfDice
}

func main() {
	inputArgs := os.Args
	dice := DieSetUp(inputArgs)

	for _, die := range dice {
		result := die.Roll()
		fmt.Printf("I rolled a d%d and it was %d\n", die.NumberOfFaces, result)
	}
}

// TODO create an extension for the rollies game, or maybe a sub code
