package main

import (
	"fmt"
	"math"
)

func main() {

	IMT := calculateIMT(getUserInput())

	outputResult(IMT)
}

func outputResult(IMT float64) {
	fmt.Printf("Ваш индекс массы тела: %.2f \n", IMT)
}

func calculateIMT(userKg float64, userHeight float64) float64 {
	const IMTPower = 2
	return userKg / math.Pow(userHeight/100, IMTPower)
}

func getUserInput() (userKg float64, userHeight float64) {
	fmt.Print("__ Калькулятор индекса массы тела __ \n")
	fmt.Print("Введите свой рост в сантиметрах: ")
	fmt.Scan(&userHeight)
	fmt.Print("Введите свой вес: ")
	fmt.Scan(&userKg)

	return
}
