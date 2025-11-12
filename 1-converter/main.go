package main

import "fmt"

func main() {
	const USDtoEUR = 0.86
	const USDtoRUB = 81.29
	const EURtoRUB = USDtoRUB / USDtoEUR
	fmt.Println(EURtoRUB)
}
