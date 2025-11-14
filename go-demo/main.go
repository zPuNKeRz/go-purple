package main

import (
	"fmt"
	"math"
)

func main() {
	const IMTPower = 2
	userHeight := 1.7
	userKg := 99.0
	IMT := userKg / math.Pow(userHeight, IMTPower)
	fmt.Print(IMT)
}
