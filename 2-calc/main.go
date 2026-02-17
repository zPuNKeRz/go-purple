package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var action string
	var numbers string

	fmt.Println("Введите дейтсиве AVG, SUM, MED:")
	fmt.Scan(&action)
	fmt.Println("Введите числа через запятую:")
	fmt.Scan(&numbers)

	switch action {
	case "AVG":
		fmt.Println("Среднее: ", calculateAvg(stringToSlice(numbers)))
	case "SUM":
		fmt.Println("Сумма: ", calculateSum(stringToSlice(numbers)))
	case "MED":
		fmt.Println("Медиана: ", calculateMed(stringToSlice(numbers)))
	default:
		fmt.Println("Неверное действие")
	}
}

func stringToSlice(input string) []float64 {
	var numbers []float64

	for s := range strings.SplitSeq(input, ",") {
		if f, err := strconv.ParseFloat(strings.TrimSpace(s), 64); err == nil {
			numbers = append(numbers, f)
		}
	}

	return numbers
}

func calculateAvg(numbers []float64) (result float64) {
	for _, value := range numbers {
		result += value
	}

	result = result / float64(len(numbers))

	return
}

func calculateSum(numbers []float64) (result float64) {
	for _, value := range numbers {
		result += value
	}

	return
}

func calculateMed(numbers []float64) float64 {
	temp := make([]float64, len(numbers))
	copy(temp, numbers)

	sort.Float64s(temp)

	n := len(temp)
	middle := n / 2

	if n%2 != 0 {
		return float64(temp[middle])
	} else {
		return float64(temp[middle-1]+temp[middle]) / 2.0
	}
}
