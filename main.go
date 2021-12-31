package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	const faces = 6
	const min = 1
	result := rand.Intn(faces-min) + min
	fmt.Println(result)

}
