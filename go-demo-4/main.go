package main

import (
	"fmt"
	"go-demo-4/account"
	"go-demo-4/file"
)

func main() {
	fmt.Println("__Менеджер паролей__")
	for {
		variant := getMenu()

		switch variant {
		case 1:
			createAccount()
		case 2:
			findAccount()
		case 3:
			deleteAccount()
		case 4:
			fmt.Println("Выход")
			return
		default:
			fmt.Println("Неверный выбор")
		}
	}

}

func getMenu() int {
	var variant int
	fmt.Println("Выберите вариант:")

	fmt.Println("1. Создать аккаунт")
	fmt.Println("2. Найти аккаунт")
	fmt.Println("3. Удалить аккаунт")
	fmt.Println("4. Выход")
	fmt.Scan(&variant)

	return variant
}

func createAccount() {
	login := promptData("Введите логин:")
	password := promptData("Введите пароль:")
	url := promptData("Введите URL:")

	myAccount, error := account.NewAccount(login, password, url)
	if error != nil {
		fmt.Println("Неверный формат URL или логин")
		return
	}

	vault := account.NewVault()
	vault.AddAccount(*myAccount)

	data, err := vault.ToBytes()
	if err != nil {
		fmt.Println("Не удалось преобразовать в JSON")

		return
	}

	file.WriteFile(data, "data.json")
	fmt.Println("Данные сохранены.")
}

func findAccount() {}

func deleteAccount() {}

func promptData(prompt string) string {
	fmt.Print(prompt + " ")
	var result string
	fmt.Scanln(&result)

	return result
}
