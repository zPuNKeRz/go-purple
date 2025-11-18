package main

import "fmt"

func main() {
	const USDtoEUR = 0.86
	const USDtoRUB = 81.29
	const EURtoRUB = USDtoRUB / USDtoEUR
	fmt.Println(EURtoRUB)
	result := getUserInfo()

	fmt.Println("Пользователь ввел: ", result)
}

func getUserInfo() float64 {
	var userData float64

	fmt.Println("Введите свои данные: ")
	fmt.Scan(&userData)

	return userData
}

func convert(count int, firstCurrency string, secondCurrency string) {}
