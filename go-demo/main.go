package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	fmt.Print("__ Калькулятор индекса массы тела __ \n")
	for {
		IMT, err := calculateIMT(getUserInput())
		if err != nil {
			panic(err)

			// fmt.Println(err)
			// continue
		}

		outputResult(IMT)

		isRepeateCalculation := checkRepeatCalculation()
		if !isRepeateCalculation {
			break
		}
	}
}

func outputResult(IMT float64) {
	fmt.Printf("Ваш индекс массы тела: %.2f \n", IMT)

	switch {
	case IMT < 16:
		fmt.Println("У вас сильный дефицит массы тела")
	case IMT < 18.5:
		fmt.Println("У вас дефицит массы тела")
	case IMT < 25:
		fmt.Println("У вас норма")
	case IMT < 30:
		fmt.Println("У вас избыточная масса")
	case IMT < 35:
		fmt.Println("У вас 1-я степерь ожирения")
	case IMT < 40:
		fmt.Println("У вас 2-я степерь ожирения")
	case IMT >= 40:
		fmt.Println("У вас 3-я степерь ожирения")
	default:
		fmt.Println("Значение по умолчанию")
	}
}

func calculateIMT(userKg float64, userHeight float64) (float64, error) {
	if userKg <= 0 || userHeight <= 0 {
		return 0, errors.New("Не указан рост или вес")
	}

	const IMTPower = 2
	return userKg / math.Pow(userHeight/100, IMTPower), nil
}

func getUserInput() (userKg float64, userHeight float64) {
	fmt.Print("Введите свой рост в сантиметрах: ")
	fmt.Scan(&userHeight)
	fmt.Print("Введите свой вес: ")
	fmt.Scan(&userKg)

	return
}

func checkRepeatCalculation() bool {
	var userChoice string
	fmt.Println("Хотите повторить вычисления? (y/n)")
	fmt.Scan(&userChoice)

	if userChoice == "y" || userChoice == "Y" {
		return true
	}

	return false
}
