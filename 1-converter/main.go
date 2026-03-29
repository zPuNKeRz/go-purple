package main

import (
	"fmt"
	"strings"
)

func main() {
	sourceCurrency := strings.ToUpper(getSourceCurrency())
	coinsCount := getCoinsCount()
	targetCurrency := getTargetCurrency(sourceCurrency)
	fmt.Println("--- MAIN ---")
	fmt.Println(sourceCurrency)
	fmt.Println(coinsCount)
	fmt.Println(targetCurrency)

	result := convert(coinsCount, sourceCurrency, targetCurrency)

	fmt.Println("Результат: ", result)
}

func getSourceCurrency() (sourceCurrency string) {

	for {

		fmt.Println("Введите исходную валюту. ")
		fmt.Println("Варианты: ")
		fmt.Println("- USD")
		fmt.Println("- EUR")
		fmt.Println("- RUB")
		fmt.Println("---------")

		fmt.Print("Ввод: ")
		fmt.Scan(&sourceCurrency)

		switch {
		case strings.ToUpper(sourceCurrency) == "USD":
			fmt.Println("Исходная валюта USD")
			fmt.Println("----")
			return

		case strings.ToUpper(sourceCurrency) == "EUR":
			fmt.Println("Исходная валюта EUR")
			fmt.Println("----")
			return

		case strings.ToUpper(sourceCurrency) == "RUB":
			fmt.Println("Исходная валюта RUB")
			fmt.Println("----")
			return
		default:
			fmt.Println("Такой валюты нет. Выберите из списка.")
			continue

		}

	}
}

func getCoinsCount() (coinsCount float64) {

	for {

		fmt.Println("Введите кол-во монет. ")
		fmt.Println("---------")

		fmt.Print("Ввод: ")
		fmt.Scan(&coinsCount)

		if coinsCount > 0 {
			fmt.Println(coinsCount)
			return
		} else {
			fmt.Println("Введите правильное кол-во")
			continue
		}

		// if strconv.ParseFloat()

		// switch {
		// case strings.ToUpper(sourceCurrency) == "USD":
		// 	fmt.Println("Исходная валюта USD")
		// 	fmt.Println("----")
		// 	return

		// case strings.ToUpper(sourceCurrency) == "EUR":
		// 	fmt.Println("Исходная валюта EUR")
		// 	fmt.Println("----")
		// 	return

		// case strings.ToUpper(sourceCurrency) == "RUB":
		// 	fmt.Println("Исходная валюта RUB")
		// 	fmt.Println("----")
		// 	return
		// default:
		// 	fmt.Println("Такой валюты нет. Выберите из списка.")
		// 	continue

		// }

	}

}

func getTargetCurrency(sourceCurrency string) (targetCurrency string) {

	for {

		fmt.Println("Введите целевую валюту. ")
		fmt.Println("Варианты: ")
		if sourceCurrency != "USD" {
			fmt.Println("- USD")
		}

		if sourceCurrency != "EUR" {
			fmt.Println("- EUR")
		}

		if sourceCurrency != "RUB" {
			fmt.Println("- RUB")
		}

		fmt.Println("---------")

		fmt.Print("Ввод: ")
		fmt.Scan(&targetCurrency)

		switch {
		case strings.ToUpper(targetCurrency) == "USD":
			fmt.Println("Целевая валюта USD")
			fmt.Println("----")
			return

		case strings.ToUpper(targetCurrency) == "EUR":
			fmt.Println("Целевая валюта EUR")
			fmt.Println("----")
			return

		case strings.ToUpper(targetCurrency) == "RUB":
			fmt.Println("Целевая валюта RUB")
			fmt.Println("----")
			return
		default:
			fmt.Println("Такой целевой валюты нет. Выберите из списка.")
			continue
		}

	}

}

func convert(count float64, sourceCurrency string, targetCurrency string) (result float64) {
	const USDtoEUR = 0.86
	const USDtoRUB = 81.29

	// Вспомогательная переменная для суммы в USD
	var usdAmount float64

	// --- ШАГ 1: Конвертация исходной валюты в USD ---
	switch sourceCurrency {
	case "USD":
		// Если исходная валюта USD, сумма остается прежней
		usdAmount = count
	case "EUR":
		// Чтобы получить USD из EUR, делим на курс USDtoEUR
		usdAmount = count / USDtoEUR
	case "RUB":
		// Чтобы получить USD из RUB, делим на курс USDtoRUB
		usdAmount = count / USDtoRUB
	default:
		// Обработка невалидной исходной валюты
		fmt.Printf("Ошибка: Неизвестная исходная валюта %s. Возвращаем 0.0.\n", sourceCurrency)
		return 0.0
	}

	// --- ШАГ 2: Конвертация из USD в целевую валюту ---
	switch targetCurrency {
	case "USD":
		// Если целевая валюта USD, результат равен сумме в USD
		result = usdAmount
	case "EUR":
		// Чтобы получить EUR из USD, умножаем на курс USDtoEUR
		result = usdAmount * USDtoEUR
	case "RUB":
		// Чтобы получить RUB из USD, умножаем на курс USDtoRUB
		result = usdAmount * USDtoRUB
	default:
		// Обработка невалидной целевой валюты
		fmt.Printf("Ошибка: Неизвестная целевая валюта %s. Возвращаем 0.0.\n", targetCurrency)
		return 0.0
	}

	return result
}
